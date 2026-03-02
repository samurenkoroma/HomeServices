CREATE TABLE land_structure
(
    id         UUID PRIMARY KEY,
    root_id    UUID NOT NULL ,
    parent_id  UUID NULL REFERENCES land_structure(id),

    unit_type  TEXT NOT NULL, 

    name       TEXT NOT NULL,

    length     NUMERIC NOT NULL,
    width      NUMERIC NOT NULL,

    created_at TIMESTAMP NOT NULL
);
