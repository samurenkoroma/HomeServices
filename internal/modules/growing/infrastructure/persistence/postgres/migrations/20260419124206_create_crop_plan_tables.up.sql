CREATE TABLE IF NOT EXISTS growing_crop_plans
(
    id            UUID PRIMARY KEY,
    name          TEXT  NOT NULL,
    area_id       UUID REFERENCES growing_cultivation_areas (id),
    variety_id    UUID REFERENCES growing_varieties (id),
    season_id     UUID REFERENCES growing_seasons (id),
    assigned_to   UUID,
    harvest_kg    DECIMAL(10, 2)           DEFAULT 0,
    status        text,

    metadata      JSONB NOT NULL           DEFAULT '[]',
    stages        JSONB NOT NULL           DEFAULT '[]',

    planting_date DATE                     DEFAULT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    started_at    TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    completed_at  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

