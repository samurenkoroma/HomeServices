CREATE TABLE crop_plans
(
    id           TEXT PRIMARY KEY,
    crop_type_id TEXT      NOT NULL,
    variety_id   TEXT NULL,
    name         TEXT      NOT NULL,
    duration     INT       NOT NULL,
    version      INT       NOT NULL,
    status       TEXT      NOT NULL,

    min_temp     DOUBLE PRECISION,
    max_temp     DOUBLE PRECISION,
    min_humidity DOUBLE PRECISION,
    max_humidity DOUBLE PRECISION,
    min_ph       DOUBLE PRECISION,
    max_ph       DOUBLE PRECISION,

    nitrogen     DOUBLE PRECISION,
    phosphorus   DOUBLE PRECISION,
    potassium    DOUBLE PRECISION,

    created_at   TIMESTAMP NOT NULL DEFAULT NOW()
);