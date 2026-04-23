package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/modules/auth/domain"
	"time"

	"github.com/redis/go-redis/v9"
)

// SessionRepository репозиторий для хранения сессий в Redis
type SessionRepository struct {
	client *redis.Client
}

// NewSessionRepository создает новый Redis репозиторий
func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{client: client}
}

// SessionData данные сессии
type SessionData struct {
	UserID    string      `json:"user_id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Role      domain.Role `json:"role"`
	Token     string      `json:"token"`
	CreatedAt time.Time   `json:"created_at"`
	ExpiresAt time.Time   `json:"expires_at"`
}

// SaveToken сохраняет токен в Redis
func (r *SessionRepository) SaveToken(ctx context.Context, token string, userID, username, email string, role domain.Role, ttl time.Duration) error {
	session := SessionData{
		UserID:    userID,
		Username:  username,
		Email:     email,
		Role:      role,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(ttl),
	}

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	// Сохраняем по ключу token
	err = r.client.Set(ctx, "session:"+token, data, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	// Сохраняем индекс по user_id
	err = r.client.SAdd(ctx, "user_sessions:"+userID, token).Err()
	if err != nil {
		return fmt.Errorf("failed to save user session index: %w", err)
	}

	return nil
}

// GetToken извлекает токен из Redis
func (r *SessionRepository) GetToken(ctx context.Context, token string) (*SessionData, error) {
	data, err := r.client.Get(ctx, "session:"+token).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, domain.ErrTokenExpired
		}
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	var session SessionData
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	// Проверяем, не истек ли срок
	if session.ExpiresAt.Before(time.Now()) {
		r.DeleteToken(ctx, token)
		return nil, domain.ErrTokenExpired
	}

	return &session, nil
}

// DeleteToken удаляет токен из Redis
func (r *SessionRepository) DeleteToken(ctx context.Context, token string) error {
	// Получаем сессию, чтобы знать user_id для удаления индекса
	session, err := r.GetToken(ctx, token)
	if err != nil {
		// Если сессия не найдена, просто удаляем ключ
		return r.client.Del(ctx, "session:"+token).Err()
	}

	// Удаляем индекс
	if err := r.client.SRem(ctx, "user_sessions:"+session.UserID, token).Err(); err != nil {
		return fmt.Errorf("failed to remove user session index: %w", err)
	}

	// Удаляем саму сессию
	return r.client.Del(ctx, "session:"+token).Err()
}

// DeleteAllUserTokens удаляет все токены пользователя
func (r *SessionRepository) DeleteAllUserTokens(ctx context.Context, userID string) error {
	// Получаем все токены пользователя
	tokens, err := r.client.SMembers(ctx, "user_sessions:"+userID).Result()
	if err != nil {
		return fmt.Errorf("failed to get user sessions: %w", err)
	}

	// Удаляем каждый токен
	for _, token := range tokens {
		if err := r.client.Del(ctx, "session:"+token).Err(); err != nil {
			// Логируем ошибку, но продолжаем
			continue
		}
	}

	// Удаляем индекс
	return r.client.Del(ctx, "user_sessions:"+userID).Err()
}

// ExtendToken продлевает срок действия токена
func (r *SessionRepository) ExtendToken(ctx context.Context, token string, ttl time.Duration) error {
	return r.client.Expire(ctx, "session:"+token, ttl).Err()
}

// IsTokenValid проверяет, существует ли токен и не истек ли он
func (r *SessionRepository) IsTokenValid(ctx context.Context, token string) bool {
	exists, err := r.client.Exists(ctx, "session:"+token).Result()
	if err != nil {
		return false
	}
	return exists > 0
}

// GetActiveTokens возвращает все активные токены пользователя
func (r *SessionRepository) GetActiveTokens(ctx context.Context, userID string) ([]string, error) {
	return r.client.SMembers(ctx, "user_sessions:"+userID).Result()
}
