-- Сорта культур
CREATE TABLE varieties
(
    id                  TEXT PRIMARY KEY,
    crop_type_id        TEXT                     NOT NULL REFERENCES crop_types(id),
    name                TEXT                     NOT NULL,
    description         TEXT,

    attributes  JSONB,
    is_active           BOOLEAN DEFAULT TRUE,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL,

    UNIQUE (crop_type_id, name)
);

CREATE INDEX idx_varieties_name_per_crop ON varieties (crop_type_id, name);
CREATE INDEX idx_varieties_active ON varieties (is_active);