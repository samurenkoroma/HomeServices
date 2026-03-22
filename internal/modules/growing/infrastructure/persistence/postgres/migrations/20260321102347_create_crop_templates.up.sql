-- Шаблоны выращивания
CREATE TABLE crop_templates
(
    id           UUID PRIMARY KEY                  DEFAULT,
    crop_plan_id UUID                     NOT NULL,
    name         TEXT                     NOT NULL,
    version      INTEGER                  NOT NULL DEFAULT 1,
    status       TEXT                     NOT NULL DEFAULT 'draft',
    requirements JSONB,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    published_at TIMESTAMP WITH TIME ZONE,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    UNIQUE (crop_plan_id, version)
);

CREATE INDEX idx_crop_templates_crop_plan ON crop_templates (crop_plan_id);
CREATE INDEX idx_crop_templates_status ON crop_templates (status);