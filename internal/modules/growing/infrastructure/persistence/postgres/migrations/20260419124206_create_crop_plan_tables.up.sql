CREATE TABLE growing_crop_plans
(
    id                        UUID PRIMARY KEY,

    crop_key                  TEXT      NOT NULL,
    variety_id                UUID,

    season_id                 UUID      NOT NULL REFERENCES growing_seasons (id) ON DELETE RESTRICT,
    area_id                   UUID      NOT NULL REFERENCES growing_cultivation_areas (id) ON DELETE RESTRICT, -- если есть география
    owner_id                  UUID      NOT NULL REFERENCES auth_organizations (id) ON DELETE RESTRICT,

    start_date                DATE      NOT NULL,
    status                    TEXT      NOT NULL DEFAULT 'draft',

    cultivation_plan_id       UUID      NOT NULL,
    cultivation_plan_version  INT       NOT NULL,
    cultivation_plan_snapshot JSONB     NOT NULL,

    created_at                TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at                TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE growing_crop_plan_snapshot
(
    crop_plan_id UUID PRIMARY KEY REFERENCES growing_crop_plans (id) ON DELETE CASCADE,
    data         JSONB NOT NULL
);