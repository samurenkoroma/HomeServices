-- Циклы выращивания
CREATE TABLE crop_cycles
(
    id                TEXT PRIMARY KEY,
    area_id           TEXT                     NOT NULL REFERENCES cultivation_areas (id),
    season_id         TEXT                     NOT NULL REFERENCES seasons (id),
    crop_plan_id      TEXT                     NOT NULL,
    crop_plan_version INTEGER                  NOT NULL,
    status            TEXT                     NOT NULL DEFAULT 'draft',
    started_at        TIMESTAMP WITH TIME ZONE,
    finished_at       TIMESTAMP WITH TIME ZONE,
    yield_actual      DOUBLE PRECISION,
    yield_estimated   DOUBLE PRECISION,
    yield_quality     TEXT,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL,

    version           INTEGER                  NOT NULL DEFAULT 1
);

CREATE INDEX idx_crop_cycles_area ON crop_cycles (area_id);
CREATE INDEX idx_crop_cycles_season ON crop_cycles (season_id);
CREATE INDEX idx_crop_cycles_status ON crop_cycles (status);
CREATE INDEX idx_crop_cycles_crop_plan ON crop_cycles (crop_plan_id);