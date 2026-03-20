-- Этапы роста в плане
CREATE TABLE crop_plan_stages (
                                  plan_id         TEXT NOT NULL REFERENCES crop_plans(id) ON DELETE CASCADE,
                                  stage_order     INTEGER NOT NULL,
                                  name            TEXT NOT NULL,
                                  duration        INTEGER NOT NULL,
                                  min_temp        DOUBLE PRECISION,
                                  max_temp        DOUBLE PRECISION,
                                  optimal_temp    DOUBLE PRECISION,
                                  water_per_day   DOUBLE PRECISION,
                                  nitrogen_req    DOUBLE PRECISION,
                                  phosphorus_req  DOUBLE PRECISION,
                                  potassium_req   DOUBLE PRECISION,

                                  PRIMARY KEY (plan_id, stage_order)
);

CREATE INDEX idx_plan_stages_order ON crop_plan_stages(plan_id, stage_order);