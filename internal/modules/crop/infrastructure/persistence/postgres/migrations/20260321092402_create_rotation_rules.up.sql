-- Правила севооборота
CREATE TABLE crop_rotation_rules
(
    plan_id                  TEXT    NOT NULL REFERENCES crop_plans (id) ON DELETE CASCADE,
    predecessor_crop_type_id TEXT    NOT NULL REFERENCES crop_types (id),
    min_years                INTEGER NOT NULL,
    recommended              BOOLEAN DEFAULT FALSE,
    notes                    TEXT,

    PRIMARY KEY (plan_id, predecessor_crop_type_id)
);

CREATE INDEX idx_rotation_rules_predecessor ON crop_rotation_rules (predecessor_crop_type_id);