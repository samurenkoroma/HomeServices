-- Циклы выращивания
CREATE TABLE growing_crop_cycles
(
    id                UUID PRIMARY KEY,
    template_id       UUID                     NOT NULL REFERENCES growing_crop_templates (id),
    area_id           UUID                     NOT NULL REFERENCES growing_cultivation_areas (id),
    season_id         UUID                     NOT NULL REFERENCES growing_seasons (id),
    crop_plan_id      UUID                     NOT NULL,
    crop_plan_version INTEGER                  NOT NULL,
    status            TEXT                     NOT NULL DEFAULT 'draft',
    started_at        TIMESTAMP WITH TIME ZONE,
    finished_at       TIMESTAMP WITH TIME ZONE,
    yield_actual      DOUBLE PRECISION,
    yield_estimated   DOUBLE PRECISION,
    yield_quality     TEXT,
    yield_notes       TEXT,
    version           INTEGER                  NOT NULL DEFAULT 1,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT check_status CHECK (status IN
                                   ('draft', 'active', 'growing', 'harvested', 'completed', 'failed', 'cancelled'))
);

CREATE INDEX idx_crop_cycles_area ON growing_crop_cycles (area_id);
CREATE INDEX idx_crop_cycles_season ON growing_crop_cycles (season_id);
CREATE INDEX idx_crop_cycles_status ON growing_crop_cycles (status);
CREATE INDEX idx_crop_cycles_crop_plan ON growing_crop_cycles (crop_plan_id);