package httpapi

import (
	"net/http"
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/commands"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/pkg/response"
)

// RouterConfig конфигурация роутера
type RouterConfig struct {
	CommandRouter command.Router
	QueryRouter   query.Router
	UowFactory    repository.Factory
	JWTService    *jwt.Service
}

// NewRouter создает новый HTTP роутер
func NewRouter(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	// ========== AUTH ЭНДПОИНТЫ (без CQRS) ==========
	authHandler := commands.NewAuthHandler(cfg.UowFactory, cfg.JWTService)

	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	//mux.HandleFunc("GET /auth/me", authHandler.Me)

	authMiddleware := NewAuthMiddleware(cfg.JWTService)
	//// Защищенные эндпоинты (требуют аутентификации)
	//mux.Handle("POST /auth/logout", authMiddleware.Authenticate(
	//	http.HandlerFunc(authHandler.Logout),
	//))
	//mux.Handle("GET /auth/me", authMiddleware.Authenticate(
	//	http.HandlerFunc(authHandler.Me),
	//))

	// ========== CQRS ЭНДПОИНТЫ ==========
	// Команды и запросы идут через единый endpoint с аутентификацией

	protectedMux := http.NewServeMux()
	protectedMux.Handle("/commands/", CommandEndpoint(cfg.CommandRouter))
	protectedMux.Handle("/queries/", QueryEndpoint(cfg.QueryRouter))

	// Применяем middleware для защиты
	protectedHandler := authMiddleware.Authenticate(protectedMux)

	// Монтируем защищенные эндпоинты
	mux.Handle("/api/", http.StripPrefix("/api", protectedHandler))

	// Опционально: эндпоинт для health check (без аутентификации)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(map[string]string{"status": "ok"}).WriteJSON(w, http.StatusOK)
	})

	// Применяем глобальные middleware (логирование, CORS, recovery)
	return withGlobalMiddleware(mux)
}

// withGlobalMiddleware применяет глобальные middleware ко всем запросам
func withGlobalMiddleware(next http.Handler) http.Handler {
	// Цепочка middleware (порядок важен!)
	handler := loggingMiddleware(next)
	handler = corsMiddleware(handler)
	handler = recoveryMiddleware(handler)
	return handler
}

// corsMiddleware добавляет CORS заголовки
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Organization-ID")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware восстанавливается после паники
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response.WriteInternalError(w, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
