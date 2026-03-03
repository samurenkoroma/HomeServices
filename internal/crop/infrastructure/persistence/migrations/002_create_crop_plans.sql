-- Таблица планов выращивания
CREATE TABLE crop_plans
(
    id           VARCHAR(36) PRIMARY KEY,
    bed_id       VARCHAR(36)              NOT NULL,
    name         VARCHAR(255)             NOT NULL,
    crop_name    VARCHAR(255)             NOT NULL,
    status       VARCHAR(50)              NOT NULL DEFAULT 'draft',
    stages       JSONB                    NOT NULL DEFAULT '[]',
    -- Даты
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    started_at   TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,

    -- Урожай
    harvest_kg   NUMERIC(10, 2)                    DEFAULT 0,

    -- Индексы
    CONSTRAINT fk_crop_plans_bed FOREIGN KEY (bed_id) REFERENCES land_structure (id)
);

CREATE INDEX idx_crop_plans_bed_id ON crop_plans (bed_id);
CREATE INDEX idx_crop_plans_status ON crop_plans (status);
CREATE INDEX idx_crop_plans_dates ON crop_plans (started_at, completed_at);

COMMENT
ON TABLE crop_plans IS 'Планы выращивания культур на грядках';
COMMENT
ON COLUMN crop_plans.status IS 'draft, active, completed, cancelled';