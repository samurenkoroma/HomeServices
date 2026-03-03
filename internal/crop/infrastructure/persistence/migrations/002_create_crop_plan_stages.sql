CREATE TABLE crop_plan_stages
(
    plan_id       TEXT NOT NULL,
    stage_order   INT  NOT NULL,
    name          TEXT NOT NULL,
    duration      INT  NOT NULL,
    min_temp      DOUBLE PRECISION,
    max_temp      DOUBLE PRECISION,
    water_per_day DOUBLE PRECISION,

    PRIMARY KEY (plan_id, stage_order),
    FOREIGN KEY (plan_id) REFERENCES crop_plans (id) ON DELETE CASCADE
);