CREATE TABLE growing_tasks
(
    id UUID PRIMARY KEY,

    crop_plan_id UUID NOT NULL REFERENCES growing_crop_plans(id) ON DELETE CASCADE,

    step_id TEXT NOT NULL, -- из snapshot

    status TEXT NOT NULL, -- pending | done

    planned_date DATE,
    completed_at TIMESTAMP
);