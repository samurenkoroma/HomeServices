-- Этапы роста в плане
CREATE TABLE crop_crop_plan_stages
(
    plan_id         UUID    NOT NULL REFERENCES crop_crop_plans (id) ON DELETE CASCADE,
    stage_order     INTEGER NOT NULL,
    name            TEXT    NOT NULL,
    duration        INTEGER NOT NULL,
    recommendations JSONB,

    PRIMARY KEY (plan_id, stage_order)
);

CREATE INDEX idx_plan_stages_order ON crop_crop_plan_stages (plan_id, stage_order);