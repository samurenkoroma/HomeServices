-- Таблица заданий для агрономов
CREATE TABLE tasks
(
    id                 VARCHAR(36) PRIMARY KEY,
    plan_id            VARCHAR(36),
    bed_id             VARCHAR(36)              NOT NULL,
    assigned_to        VARCHAR(100),

    type               VARCHAR(50)              NOT NULL,
    status             VARCHAR(50)              NOT NULL DEFAULT 'pending',
    priority           VARCHAR(20)              NOT NULL DEFAULT 'medium',

    title              VARCHAR(255)             NOT NULL,
    description        TEXT,
    instructions       TEXT,

    scheduled_date     DATE                     NOT NULL,
    due_date           DATE                     NOT NULL,
    completed_at       TIMESTAMP WITH TIME ZONE,

    estimated_duration INTEGER,
    actual_duration    INTEGER,

    -- JSON поля для гибких данных
    location           JSONB,
    weather            JSONB,
    photos             JSONB                             DEFAULT '[]',
    comments           JSONB                             DEFAULT '[]',

    created_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    -- Связи
    CONSTRAINT fk_tasks_plan FOREIGN KEY (plan_id) REFERENCES crop_plans (id) ON DELETE SET NULL,
    CONSTRAINT fk_tasks_bed FOREIGN KEY (bed_id) REFERENCES land_structure (id)
);

CREATE INDEX idx_tasks_assigned_to ON tasks (assigned_to);
CREATE INDEX idx_tasks_scheduled_date ON tasks (scheduled_date);
CREATE INDEX idx_tasks_status ON tasks (status);
CREATE INDEX idx_tasks_priority ON tasks (priority);

COMMENT ON TABLE tasks IS 'Ежедневные задания для агрономов';