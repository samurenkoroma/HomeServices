-- Конфигурации мест по сезонам
CREATE TABLE area_season_configs
(
    area_id      TEXT                     NOT NULL REFERENCES cultivation_areas (id) ON DELETE CASCADE,
    season_id    TEXT                     NOT NULL REFERENCES seasons (id) ON DELETE CASCADE,
    name         TEXT                     NOT NULL,
    geometry     GEOMETRY(Geometry, 4326) NOT NULL,
    area         DOUBLE PRECISION         NOT NULL,
    crop_plan_id TEXT,
    block_ids    TEXT[], -- для полей с поликультурой
    metadata     JSONB,
    valid_from   TIMESTAMP WITH TIME ZONE NOT NULL,
    valid_until  TIMESTAMP WITH TIME ZONE,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    PRIMARY KEY (area_id, season_id)
);

CREATE INDEX idx_area_configs_season ON area_season_configs (season_id);
CREATE INDEX idx_area_configs_crop_plan ON area_season_configs (crop_plan_id);
CREATE INDEX idx_area_configs_geometry ON area_season_configs USING GIST (geometry);