package middleware

import (
	"context"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"samurenkoroma/services/configs"
	"samurenkoroma/services/pkg/jwt"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func JWTProtected(cfg configs.AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(cfg.AccessSecret)},
			SuccessHandler: func(c *fiber.Ctx) error {
				authHeader := c.Get("Authorization")
				token := strings.TrimPrefix(authHeader, "Bearer ")
				_, data := jwt.NewJWT(cfg).ParseAccess(token)

				c.Locals("email", data.Email)
				return c.Next()
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				// Return status 401 and failed authentication error.
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				})
			},
		})(c)
	}
}
func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			log.Println("not auth header")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth).ParseAccess(token)

		if !isValid {
			log.Println("invalid auth token")
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
