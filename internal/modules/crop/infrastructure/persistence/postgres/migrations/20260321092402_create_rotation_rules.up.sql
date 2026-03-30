-- Правила севооборота
CREATE TABLE crop_crop_rotation_rules
(
    plan_id                  UUID    NOT NULL REFERENCES crop_crop_plans (id) ON DELETE CASCADE,
    predecessor_crop_type_id UUID    NOT NULL REFERENCES crop_crop_types (id),
    min_years                INTEGER NOT NULL,
    recommended              BOOLEAN DEFAULT FALSE,
    notes                    TEXT,

    PRIMARY KEY (plan_id, predecessor_crop_type_id)
);

CREATE INDEX idx_rotation_rules_predecessor ON crop_crop_rotation_rules (predecessor_crop_type_id);