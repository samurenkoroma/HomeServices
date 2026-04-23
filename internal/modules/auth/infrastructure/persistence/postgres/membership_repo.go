package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type MembershipRepository struct {
	tx *sql.Tx
}

func NewMembershipRepository(tx *sql.Tx) *MembershipRepository {
	return &MembershipRepository{tx: tx}
}

// Save сохраняет членство
func (r *MembershipRepository) Save(ctx context.Context, m *domain.Membership) error {
	query := `
        INSERT INTO auth_memberships (
            id, user_id, organization_id, role, is_active, joined_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err := r.tx.ExecContext(ctx, query,
		m.ID,
		m.UserID,
		m.OrganizationID,
		string(m.Role),
		m.IsActive,
		m.JoinedAt,
		m.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save membership: %w", err)
	}

	return nil
}

// Update обновляет членство
func (r *MembershipRepository) Update(ctx context.Context, m *domain.Membership) error {
	query := `
        UPDATE auth_memberships SET
            role = $2,
            is_active = $3,
            updated_at = $4
        WHERE id = $1
    `

	result, err := r.tx.ExecContext(ctx, query,
		m.ID,
		string(m.Role),
		m.IsActive,
		m.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update membership: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrMembershipNotFound
	}

	return nil
}

// FindByID находит членство по ID
func (r *MembershipRepository) FindByID(ctx context.Context, id string) (*domain.Membership, error) {
	query := `
        SELECT id, user_id, organization_id, role, is_active, joined_at, updated_at
        FROM auth_memberships
        WHERE id = $1
    `

	var m domain.Membership
	err := r.tx.QueryRowContext(ctx, query, id).Scan(
		&m.ID,
		&m.UserID,
		&m.OrganizationID,
		&m.Role,
		&m.IsActive,
		&m.JoinedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrMembershipNotFound
		}
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}

	return &m, nil
}

// FindByUser возвращает все членства пользователя
func (r *MembershipRepository) FindByUser(ctx context.Context, userID string) ([]*domain.Membership, error) {
	query := `
        SELECT id, user_id, organization_id, role, is_active, joined_at, updated_at
        FROM auth_memberships
        WHERE user_id = $1
        ORDER BY joined_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find memberships by user: %w", err)
	}
	defer rows.Close()

	var memberships []*domain.Membership
	for rows.Next() {
		var m domain.Membership
		err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.OrganizationID,
			&m.Role,
			&m.IsActive,
			&m.JoinedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan membership: %w", err)
		}
		memberships = append(memberships, &m)
	}

	return memberships, nil
}

// FindByOrganization возвращает все членства в организации
func (r *MembershipRepository) FindByOrganization(ctx context.Context, orgID string) ([]*domain.Membership, error) {
	query := `
        SELECT id, user_id, organization_id, role, is_active, joined_at, updated_at
        FROM auth_memberships
        WHERE organization_id = $1
        ORDER BY joined_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to find memberships by organization: %w", err)
	}
	defer rows.Close()

	var memberships []*domain.Membership
	for rows.Next() {
		var m domain.Membership
		err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.OrganizationID,
			&m.Role,
			&m.IsActive,
			&m.JoinedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan membership: %w", err)
		}
		memberships = append(memberships, &m)
	}

	return memberships, nil
}

// FindByUserAndOrganization находит членство пользователя в организации
func (r *MembershipRepository) FindByUserAndOrganization(ctx context.Context, userID, orgID string) (*domain.Membership, error) {
	query := `
        SELECT id, user_id, organization_id, role, is_active, joined_at, updated_at
        FROM auth_memberships
        WHERE user_id = $1 AND organization_id = $2
    `

	var m domain.Membership
	err := r.tx.QueryRowContext(ctx, query, userID, orgID).Scan(
		&m.ID,
		&m.UserID,
		&m.OrganizationID,
		&m.Role,
		&m.IsActive,
		&m.JoinedAt,
		&m.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrMembershipNotFound
		}
		return nil, fmt.Errorf("failed to find membership: %w", err)
	}

	return &m, nil
}

// Exists проверяет существование членства
func (r *MembershipRepository) Exists(ctx context.Context, userID, orgID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM auth_memberships WHERE user_id = $1 AND organization_id = $2)`

	var exists bool
	err := r.tx.QueryRowContext(ctx, query, userID, orgID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check membership existence: %w", err)
	}

	return exists, nil
}

// List возвращает список членств с фильтрацией
func (r *MembershipRepository) List(ctx context.Context, filter domain.MembershipFilter) ([]*domain.Membership, int, error) {
	baseQuery := `
        FROM auth_memberships m
        WHERE 1=1
    `
	var args []interface{}
	argIndex := 1

	if filter.UserID != "" {
		baseQuery += fmt.Sprintf(" AND m.user_id = $%d", argIndex)
		args = append(args, filter.UserID)
		argIndex++
	}

	if filter.OrganizationID != "" {
		baseQuery += fmt.Sprintf(" AND m.organization_id = $%d", argIndex)
		args = append(args, filter.OrganizationID)
		argIndex++
	}

	if filter.Role != "" {
		baseQuery += fmt.Sprintf(" AND m.role = $%d", argIndex)
		args = append(args, string(filter.Role))
		argIndex++
	}

	if filter.IsActive != nil {
		baseQuery += fmt.Sprintf(" AND m.is_active = $%d", argIndex)
		args = append(args, *filter.IsActive)
		argIndex++
	}

	// Считаем количество
	countQuery := `SELECT COUNT(*) ` + baseQuery
	var total int
	err := r.tx.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count memberships: %w", err)
	}

	// Запрос данных
	dataQuery := `
        SELECT m.id, m.user_id, m.organization_id, m.role, m.is_active, m.joined_at, m.updated_at
    ` + baseQuery + `
        ORDER BY m.joined_at DESC
    `

	if filter.Limit > 0 {
		dataQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, filter.Limit, filter.Offset)
	}

	rows, err := r.tx.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list memberships: %w", err)
	}
	defer rows.Close()

	var memberships []*domain.Membership
	for rows.Next() {
		var m domain.Membership
		err := rows.Scan(
			&m.ID,
			&m.UserID,
			&m.OrganizationID,
			&m.Role,
			&m.IsActive,
			&m.JoinedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan membership: %w", err)
		}
		memberships = append(memberships, &m)
	}

	return memberships, total, nil
}

// Delete удаляет членство
func (r *MembershipRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM auth_memberships WHERE id = $1`

	result, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete membership: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrMembershipNotFound
	}

	return nil
}

// Activate активирует членство
func (r *MembershipRepository) Activate(ctx context.Context, id string) error {
	query := `UPDATE auth_memberships SET is_active = true, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to activate membership: %w", err)
	}

	return nil
}

// Deactivate деактивирует членство
func (r *MembershipRepository) Deactivate(ctx context.Context, id string) error {
	query := `UPDATE auth_memberships SET is_active = false, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to deactivate membership: %w", err)
	}

	return nil
}

// ChangeRole изменяет роль в членстве
func (r *MembershipRepository) ChangeRole(ctx context.Context, id string, newRole domain.OrganizationRole) error {
	query := `UPDATE auth_memberships SET role = $2, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, id, string(newRole))
	if err != nil {
		return fmt.Errorf("failed to change role: %w", err)
	}

	return nil
}
