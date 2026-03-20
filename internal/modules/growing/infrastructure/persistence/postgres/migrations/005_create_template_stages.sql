-- Этапы роста в шаблоне
CREATE TABLE template_stages
(
    template_id   TEXT    NOT NULL REFERENCES crop_templates (id) ON DELETE CASCADE,
    stage_order   INTEGER NOT NULL,
    name          TEXT    NOT NULL,
    duration      INTEGER NOT NULL, -- дней
    min_temp      DOUBLE PRECISION,
    max_temp      DOUBLE PRECISION,
    min_humidity  DOUBLE PRECISION,
    max_humidity  DOUBLE PRECISION,
    water_per_day DOUBLE PRECISION, -- литров на м²

    PRIMARY KEY (template_id, stage_order)
);