CREATE TABLE growing_cultivation_plans
(
    id         UUID  NOT NULL,
    version    INT   NOT NULL,

    name       TEXT  NOT NULL,
    crop_key   TEXT  NOT NULL,
    variety_id UUID,
    steps      JSONB NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id, version)
);