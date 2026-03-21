-- Типы культур
CREATE TABLE crop_types
(
    id              TEXT PRIMARY KEY,
    name            TEXT                     NOT NULL UNIQUE,
    scientific_name TEXT,
    category        TEXT                     NOT NULL,
    description     TEXT,
    root_depth      INTEGER,
    is_perennial    BOOLEAN DEFAULT FALSE,
    vegetation_days INTEGER                  NOT NULL,
    default_yield   DOUBLE PRECISION,
    market_price    DOUBLE PRECISION,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_crop_types_category ON crop_types (category);
CREATE INDEX idx_crop_types_name ON crop_types (name);