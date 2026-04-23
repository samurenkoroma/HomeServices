package memory

import (
	"context"
	"samurenkoroma/services/internal/modules/auth/domain"
	"sync"
	"time"
)

// UserRepository in-memory реализация репозитория
type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewUserRepository создает новый in-memory репозиторий
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
	}
}

// Save сохраняет пользователя
func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return domain.ErrUserAlreadyExists
	}

	// Сохраняем копию
	userCopy := *user
	r.users[user.ID] = &userCopy
	return nil
}

// Update обновляет пользователя
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}

	userCopy := *user
	r.users[user.ID] = &userCopy
	return nil
}

// FindByID находит пользователя по ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}

	userCopy := *user
	return &userCopy, nil
}

// FindByEmail находит пользователя по email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			userCopy := *user
			return &userCopy, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

// FindByUsername находит пользователя по username
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			userCopy := *user
			return &userCopy, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

// Delete удаляет пользователя
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return domain.ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}

// List возвращает список пользователей с фильтрацией
func (r *UserRepository) List(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []*domain.User

	for _, user := range r.users {
		// Фильтр по роли
		if filter.Role != "" && user.Role != filter.Role {
			continue
		}

		// Фильтр по статусу
		if filter.Status != "" && user.Status != filter.Status {
			continue
		}

		// Поиск по тексту
		if filter.Search != "" {
			if user.Username != filter.Search &&
				user.Email != filter.Search &&
				user.FirstName != filter.Search &&
				user.LastName != filter.Search {
				continue
			}
		}

		userCopy := *user
		filtered = append(filtered, &userCopy)
	}

	total := len(filtered)

	// Пагинация
	start := filter.Offset
	end := filter.Offset + filter.Limit

	if start > total {
		start = total
	}
	if end > total || filter.Limit == 0 {
		end = total
	}

	return filtered[start:end], total, nil
}

// FindByRole находит пользователей по роли
func (r *UserRepository) FindByRole(ctx context.Context, role domain.Role) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []*domain.User
	for _, user := range r.users {
		if user.Role == role {
			userCopy := *user
			users = append(users, &userCopy)
		}
	}
	return users, nil
}

// UpdateLastLogin обновляет время последнего входа
func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[userID]
	if !exists {
		return domain.ErrUserNotFound
	}

	now := time.Now()
	user.LastLogin = &now
	user.UpdatedAt = now

	return nil
}

// Clear очищает репозиторий (для тестов)
func (r *UserRepository) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users = make(map[string]*domain.User)
}

// Count возвращает количество пользователей (для тестов)
func (r *UserRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.users)
}
