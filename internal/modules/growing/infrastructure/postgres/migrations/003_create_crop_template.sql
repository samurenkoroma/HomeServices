CREATE TABLE crop_templates
(
    plan_id    uuid PRIMARY KEY,
    version    integer     NOT NULL,
    active     boolean     NOT NULL,
    updated_at timestamptz NOT NULL
);