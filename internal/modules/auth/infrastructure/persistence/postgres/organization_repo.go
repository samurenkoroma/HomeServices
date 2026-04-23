package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type OrganizationRepository struct {
	tx *sql.Tx
}

func NewOrganizationRepository(tx *sql.Tx) *OrganizationRepository {
	return &OrganizationRepository{tx: tx}
}

// Save сохраняет организацию
func (r *OrganizationRepository) Save(ctx context.Context, org *domain.Organization) error {
	query := `
        INSERT INTO auth_organizations (
            id, name, tax_id, address, phone, email, is_active, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `

	_, err := r.tx.ExecContext(ctx, query,
		org.ID,
		org.Name,
		org.TaxID,
		org.Address,
		org.Phone,
		org.Email,
		org.IsActive,
		org.CreatedAt,
		org.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save organization: %w", err)
	}

	return nil
}

// Update обновляет организацию
func (r *OrganizationRepository) Update(ctx context.Context, org *domain.Organization) error {
	query := `
        UPDATE auth_organizations SET
            name = $2,
            tax_id = $3,
            address = $4,
            phone = $5,
            email = $6,
            is_active = $7,
            updated_at = $8
        WHERE id = $1
    `

	result, err := r.tx.ExecContext(ctx, query,
		org.ID,
		org.Name,
		org.TaxID,
		org.Address,
		org.Phone,
		org.Email,
		org.IsActive,
		org.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrOrganizationNotFound
	}

	return nil
}

// FindByID находит организацию по ID
func (r *OrganizationRepository) FindByID(ctx context.Context, id string) (*domain.Organization, error) {
	query := `
        SELECT id, name, tax_id, address, phone, email, is_active, created_at, updated_at
        FROM auth_organizations
        WHERE id = $1
    `

	var org domain.Organization

	err := r.tx.QueryRowContext(ctx, query, id).Scan(
		&org.ID,
		&org.Name,
		&org.TaxID,
		&org.Address,
		&org.Phone,
		&org.Email,
		&org.IsActive,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, fmt.Errorf("failed to find organization: %w", err)
	}

	return &org, nil
}

// FindByName находит организацию по названию
func (r *OrganizationRepository) FindByName(ctx context.Context, name string) (*domain.Organization, error) {
	query := `
        SELECT id, name, tax_id, address, phone, email, is_active, created_at, updated_at
        FROM auth_organizations
        WHERE name = $1
    `

	var org domain.Organization

	err := r.tx.QueryRowContext(ctx, query, name).Scan(
		&org.ID,
		&org.Name,
		&org.TaxID,
		&org.Address,
		&org.Phone,
		&org.Email,
		&org.IsActive,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, fmt.Errorf("failed to find organization by name: %w", err)
	}

	return &org, nil
}

// FindByTaxID находит организацию по ИНН
func (r *OrganizationRepository) FindByTaxID(ctx context.Context, taxID string) (*domain.Organization, error) {
	if taxID == "" {
		return nil, domain.ErrOrganizationNotFound
	}

	query := `
        SELECT id, name, tax_id, address, phone, email, is_active, created_at, updated_at
        FROM auth_organizations
        WHERE tax_id = $1
    `

	var org domain.Organization

	err := r.tx.QueryRowContext(ctx, query, taxID).Scan(
		&org.ID,
		&org.Name,
		&org.TaxID,
		&org.Address,
		&org.Phone,
		&org.Email,
		&org.IsActive,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, fmt.Errorf("failed to find organization by tax_id: %w", err)
	}

	return &org, nil
}

// Delete удаляет организацию
func (r *OrganizationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM auth_organizations WHERE id = $1`

	result, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return domain.ErrOrganizationNotFound
	}

	return nil
}

// List возвращает список организаций с фильтрацией
func (r *OrganizationRepository) List(ctx context.Context, filter domain.OrganizationFilter) ([]*domain.Organization, int, error) {
	baseQuery := `
        FROM organizations o
        WHERE 1=1
    `
	var args []interface{}
	argIndex := 1

	if filter.Search != "" {
		baseQuery += fmt.Sprintf(" AND (o.name ILIKE $%d OR o.tax_id ILIKE $%d)", argIndex, argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.IsActive != nil {
		baseQuery += fmt.Sprintf(" AND o.is_active = $%d", argIndex)
		args = append(args, *filter.IsActive)
		argIndex++
	}

	// Считаем количество
	countQuery := `SELECT COUNT(*) ` + baseQuery
	var total int
	err := r.tx.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count organizations: %w", err)
	}

	// Запрос данных
	dataQuery := `
        SELECT o.id, o.name, o.tax_id, o.address, o.phone, o.email, o.is_active, o.created_at, o.updated_at
    ` + baseQuery + `
        ORDER BY o.created_at DESC
    `

	if filter.Limit > 0 {
		dataQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
		args = append(args, filter.Limit, filter.Offset)
	}

	rows, err := r.tx.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list organizations: %w", err)
	}
	defer rows.Close()

	var orgs []*domain.Organization
	for rows.Next() {
		var org domain.Organization
		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.TaxID,
			&org.Address,
			&org.Phone,
			&org.Email,
			&org.IsActive,
			&org.CreatedAt,
			&org.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan organization: %w", err)
		}
		orgs = append(orgs, &org)
	}

	return orgs, total, nil
}

// ListByUser возвращает организации пользователя
func (r *OrganizationRepository) ListByUser(ctx context.Context, userID string) ([]*domain.Organization, error) {
	query := `
        SELECT o.id, o.name, o.tax_id, o.address, o.phone, o.email, o.is_active, o.created_at, o.updated_at
        FROM auth_organizations o
        INNER JOIN auth_memberships m ON m.organization_id = o.id
        WHERE m.user_id = $1 AND m.is_active = true AND o.is_active = true
        ORDER BY o.created_at DESC
    `

	rows, err := r.tx.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list organizations by user: %w", err)
	}
	defer rows.Close()

	var orgs []*domain.Organization
	for rows.Next() {
		var org domain.Organization
		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.TaxID,
			&org.Address,
			&org.Phone,
			&org.Email,
			&org.IsActive,
			&org.CreatedAt,
			&org.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan organization: %w", err)
		}
		orgs = append(orgs, &org)
	}

	return orgs, nil
}

// Activate активирует организацию
func (r *OrganizationRepository) Activate(ctx context.Context, id string) error {
	query := `UPDATE auth_organizations SET is_active = true, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to activate organization: %w", err)
	}

	return nil
}

// Deactivate деактивирует организацию
func (r *OrganizationRepository) Deactivate(ctx context.Context, id string) error {
	query := `UPDATE auth_organizations SET is_active = false, updated_at = NOW() WHERE id = $1`

	_, err := r.tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to deactivate organization: %w", err)
	}

	return nil
}
