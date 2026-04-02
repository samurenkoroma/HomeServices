-- Up migration: create_translations
-- Created: Чт 02 апр 2026 15:14:59 EEST

CREATE TABLE IF NOT EXISTS translations (
    entity text,
    latin text,
    ru text,
    PRIMARY KEY (entity, latin)
);

