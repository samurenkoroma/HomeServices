package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/shared/domain/translation"
)

type translationRepository struct {
	tx *sql.Tx
}

func NewTranslationsRepository(tx *sql.Tx) translation.Repository {
	return &translationRepository{tx: tx}
}

func (t *translationRepository) Save(ctx context.Context, entity, latin, ru string) error {
	query := `
        INSERT INTO translations (
            entity, latin, ru
        ) VALUES ($1, $2, $3)
        ON CONFLICT (entity, latin) DO UPDATE SET
            ru = EXCLUDED.ru
    `
	_, err := t.tx.ExecContext(ctx, query, entity, latin, ru)
	return err

}
