-- Типы культур
CREATE TABLE crop_crop_types
(
    id           UUID PRIMARY KEY,
    name         TEXT                     NOT NULL UNIQUE,
    category     TEXT                     NOT NULL,
    description  TEXT,
    is_perennial BOOLEAN DEFAULT FALSE,
    is_active    BOOLEAN DEFAULT FALSE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_crop_types_category ON crop_crop_types (category);
CREATE INDEX idx_crop_types_name ON crop_crop_types (name);