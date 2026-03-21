-- Места выращивания
CREATE TABLE cultivation_areas
(
    id          TEXT PRIMARY KEY,
    farm_ref_id TEXT                     NOT NULL,
    type        TEXT                     NOT NULL, -- field, block, greenhouse, bed
    name        TEXT                     NOT NULL,
    geometry    GEOMETRY(Geometry, 4326) NOT NULL,
    area        DOUBLE PRECISION         NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,

    UNIQUE (farm_ref_id, type)
);

CREATE INDEX idx_cultivation_areas_farm_ref ON cultivation_areas (farm_ref_id);
CREATE INDEX idx_cultivation_areas_type ON cultivation_areas (type);
CREATE INDEX idx_cultivation_areas_geometry ON cultivation_areas USING GIST (geometry);