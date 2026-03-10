CREATE TABLE land_structure
(
    id          UUID PRIMARY KEY,

    root_id     UUID NOT NULL REFERENCES land_structure(id),
    parent_id   UUID NULL REFERENCES land_structure(id),

    unit_type   TEXT NOT NULL,
    name        TEXT NOT NULL,

    properties  JSONB,
    geom        GEOMETRY,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    delete_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
