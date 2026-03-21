-- Шаблоны культур
CREATE TABLE crop_templates
(
    id           TEXT PRIMARY KEY,
    crop_plan_id TEXT                     NOT NULL,
    name         TEXT                     NOT NULL,
    version      INTEGER                  NOT NULL DEFAULT 1,
    status       TEXT                     NOT NULL DEFAULT 'draft',
    requirements JSONB,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL,

    UNIQUE (crop_plan_id, version)
);

CREATE INDEX idx_crop_templates_crop_plan ON crop_templates (crop_plan_id);
CREATE INDEX idx_crop_templates_status ON crop_templates (status);