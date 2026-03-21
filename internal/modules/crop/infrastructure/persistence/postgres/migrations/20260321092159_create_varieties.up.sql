-- Сорта культур
CREATE TABLE varieties
(
    id                  TEXT PRIMARY KEY,
    crop_type_id        TEXT                     NOT NULL REFERENCES crop_types (id),
    name                TEXT                     NOT NULL,
    description         TEXT,
    vegetation_days     INTEGER                  NOT NULL,
    yield_potential     DOUBLE PRECISION,
    disease_resistance  TEXT[],
    recommended_regions TEXT[],
    planting_density    INTEGER,
    seed_rate           DOUBLE PRECISION,
    breeder             TEXT,
    year_released       INTEGER,
    is_active           BOOLEAN DEFAULT TRUE,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL,

    UNIQUE (crop_type_id, name)
);

CREATE INDEX idx_varieties_crop_type ON varieties (crop_type_id);
CREATE INDEX idx_varieties_active ON varieties (is_active);