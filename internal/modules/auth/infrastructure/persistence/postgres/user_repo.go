package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type UserRepository struct {
	tx *sql.Tx
}

func NewUserRepository(tx *sql.Tx) *UserRepository {
	return &UserRepository{tx: tx}
}

// Save сохраняет нового пользователя
func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	query := `
        INSERT INTO auth_users (
            id, email, username, password, first_name, last_name, phone,
            role, status, current_organization_id, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `

	_, err := r.tx.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Phone,
		string(user.Role),
		string(user.Status),
		user.CurrentOrganizationID,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

// Update обновляет пользователя
func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
        UPDATE auth_users SET
            email = $2,
            username = $3,
            password = $4,
            first_name = $5,
            last_name = $6,
            phone = $7,
            role = $8,
            status = $9,
            current_organization_id = $10,
            last_login = $11,
            updated_at = $12
        WHERE id = $1
    `

	result, err := r.tx.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Phone,
		string(user.Role),
		string(user.Status),
		user.CurrentOrganizationID,
		user.LastLogin,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// FindByID находит пользователя по ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := `
        SELECT id, email, username, password, first_name, last_name, phone,
               role, status, current_organization_id, last_login,
               created_at, updated_at
        FROM auth_users
        WHERE id = $1
    `

	var user domain.User
	var currentOrgID sql.NullString
	var lastLogin sql.NullTime

	err := r.tx.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Role,
		&user.Status,
		&currentOrgID,
		&lastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if currentOrgID.Valid {
		user.CurrentOrganizationID = &currentOrgID.String
	}
	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

// FindByEmail находит пользователя по email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
        SELECT id, email, username, password, first_name, last_name, 
               role, status, current_organization_id, last_login,
               created_at, updated_at
        FROM auth_users
        WHERE email = $1
    `

	var user domain.User
	var currentOrgID sql.NullString
	var lastLogin sql.NullTime

	err := r.tx.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.Status,
		&currentOrgID,
		&lastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	if currentOrgID.Valid {
		user.CurrentOrganizationID = &currentOrgID.String
	}
	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

// FindByUsername находит пользователя по username
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `
        SELECT id, email, username, password, first_name, last_name, phone,
               role, status, current_organization_id, last_login,
               created_at, updated_at
        FROM auth_users
        WHERE username = $1
    `

	var user domain.User
	var currentOrgID sql.NullString
	var lastLogin sql.NullTime

	err := r.tx.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Role,
		&user.Status,
		&currentOrgID,
		&lastLogin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	if currentOrgID.Valid {
		user.CurrentOrganizationID = &currentOrgID.String
	}
	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

// Delete удаляет пользователя
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM auth_users WHERE id = $1`

	result, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// List возвращает список пользователей с фильтрацией
func (r *UserRepository) List(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int, error) {
	// Базовый запрос
	baseQuery := `
        FROM auth_users u
        WHERE 1=1
    `
	var args []interface{}
	argIndex := 1

	// Фильтр по поиску
	if filter.Search != "" {
		baseQuery += fmt.Sprintf(" AND (u.email ILIKE $%d OR u.username ILIKE $%d OR u.first_name ILIKE $%d OR u.last_name ILIKE $%d)",
			argIndex, argIndex, argIndex, argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	// Фильтр по статусу
	if filter.Status != "" {
		baseQuery += fmt.Sprintf(" AND u.status = $%d", argIndex)
		args = append(args, string(filter.Status))
		argIndex++
	}

	// Фильтр по глобальной роли
	if filter.Role != "" {
		baseQuery += fmt.Sprintf(" AND u.role = $%d", argIndex)
		args = append(args, string(filter.Role))
		argIndex++
	}

	// Фильтр по организации (через membership)
	if filter.OrgID != "" {
		baseQuery += fmt.Sprintf(" AND EXISTS (SELECT 1 FROM memberships m WHERE m.user_id = u.id AND m.organization_id = $%d)", argIndex)
		args = append(args, filter.OrgID)
		argIndex++
	}

	// Считаем общее количество
	countQuery := `SELECT COUNT(*) ` + baseQuery
	var total int
	err := r.tx.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	// Запрос данных
	dataQuery := `
        SELECT u.id, u.email, u.username, u.password, u.first_name, u.last_name, u.phone,
               u.role, u.status, u.current_organization_id, u.last_login,
               u.created_at, u.updated_at
    ` + baseQuery + `
        ORDER BY u.created_at DESC
    `

	if filter.Limit > 0 {
		dataQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, filter.Limit, filter.Offset)
	}

	rows, err := r.tx.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		var currentOrgID sql.NullString
		var lastLogin sql.NullTime

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Role,
			&user.Status,
			&currentOrgID,
			&lastLogin,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan user: %w", err)
		}

		if currentOrgID.Valid {
			user.CurrentOrganizationID = &currentOrgID.String
		}
		if lastLogin.Valid {
			user.LastLogin = &lastLogin.Time
		}

		users = append(users, &user)
	}

	return users, total, nil
}

// UpdateLastLogin обновляет время последнего входа
func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	query := `UPDATE auth_users SET last_login = NOW(), updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to update last login: %w", err)
	}

	return nil
}

// UpdateCurrentOrganization обновляет текущую организацию пользователя
func (r *UserRepository) UpdateCurrentOrganization(ctx context.Context, userID, organizationID string) error {
	var orgID interface{}
	if organizationID == "" {
		orgID = nil
	} else {
		orgID = organizationID
	}

	query := `UPDATE auth_users SET current_organization_id = $2, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, userID, orgID)
	if err != nil {
		return fmt.Errorf("failed to update current organization: %w", err)
	}

	return nil
}
