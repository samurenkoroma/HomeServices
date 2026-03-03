CREATE TABLE crop_rotation_rules
(
    plan_id                  TEXT NOT NULL,
    predecessor_crop_type_id TEXT NOT NULL,
    min_years                INT  NOT NULL,

    PRIMARY KEY (plan_id, predecessor_crop_type_id),
    FOREIGN KEY (plan_id) REFERENCES crop_plans (id) ON DELETE CASCADE
);