-- Планы выращивания
CREATE TABLE crop_plans
(
    id           TEXT PRIMARY KEY,
    crop_type_id TEXT                     NOT NULL REFERENCES crop_types (id),
    variety_id   TEXT REFERENCES varieties (id),
    name         TEXT                     NOT NULL,
    description  TEXT,
    duration     INTEGER                  NOT NULL,
    version      INTEGER                  NOT NULL DEFAULT 1,
    status       TEXT                     NOT NULL DEFAULT 'draft',
    environment  JSONB,
    nutrients    JSONB,
    created_by   TEXT                     NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE,

    UNIQUE (crop_type_id, variety_id, version)
);

CREATE INDEX idx_crop_plans_crop_type ON crop_plans (crop_type_id);
CREATE INDEX idx_crop_plans_status ON crop_plans (status);
CREATE INDEX idx_crop_plans_name ON crop_plans (name);