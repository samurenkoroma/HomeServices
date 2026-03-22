-- Циклы выращивания
CREATE TABLE crop_cycles
(
    id                UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    template_id       UUID                     NOT NULL REFERENCES crop_templates (id),
    area_id           UUID                     NOT NULL REFERENCES cultivation_areas (id),
    season_id         UUID                     NOT NULL REFERENCES seasons (id),
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

CREATE INDEX idx_crop_cycles_area ON crop_cycles (area_id);
CREATE INDEX idx_crop_cycles_season ON crop_cycles (season_id);
CREATE INDEX idx_crop_cycles_status ON crop_cycles (status);
CREATE INDEX idx_crop_cycles_crop_plan ON crop_cycles (crop_plan_id);