-- Этапы шаблона
CREATE TABLE template_stages
(
    template_id   UUID    NOT NULL REFERENCES crop_templates (id) ON DELETE CASCADE,
    stage_order   INTEGER NOT NULL,
    name          TEXT    NOT NULL,
    duration      INTEGER NOT NULL,
    min_temp      DOUBLE PRECISION,
    max_temp      DOUBLE PRECISION,
    optimal_temp  DOUBLE PRECISION,
    water_per_day DOUBLE PRECISION,
    description   TEXT,

    PRIMARY KEY (template_id, stage_order)
);

CREATE INDEX idx_template_stages_order ON template_stages (template_id, stage_order);