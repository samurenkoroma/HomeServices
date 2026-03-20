-- История изменений планов
CREATE TABLE crop_plan_history (
                                   id          SERIAL PRIMARY KEY,
                                   plan_id     TEXT NOT NULL REFERENCES crop_plans(id),
                                   version     INTEGER NOT NULL,
                                   snapshot    JSONB NOT NULL,
                                   changed_by  TEXT NOT NULL,
                                   changed_at  TIMESTAMP WITH TIME ZONE NOT NULL,
                                   reason      TEXT
);

CREATE INDEX idx_plan_history_plan ON crop_plan_history(plan_id);
CREATE INDEX idx_plan_history_version ON crop_plan_history(plan_id, version);