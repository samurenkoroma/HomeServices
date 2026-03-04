CREATE TABLE crop_cycles
(
    id           uuid PRIMARY KEY,
    plan_id      uuid    NOT NULL,
    plan_version integer NOT NULL,

    facility_id  uuid    NOT NULL,
    bed_id       uuid    NOT NULL,

    status       text    NOT NULL,
    started_at   timestamptz,
    finished_at  timestamptz,

    version      integer NOT NULL
);

