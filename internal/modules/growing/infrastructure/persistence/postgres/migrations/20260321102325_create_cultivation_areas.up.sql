-- Места выращивания
CREATE TABLE growing_cultivation_areas
(
    id          UUID PRIMARY KEY,
    farm_ref_id UUID                     NOT NULL,
    type        TEXT                     NOT NULL CHECK (type IN ('field', 'block', 'greenhouse', 'bed')),
    name        TEXT                     NOT NULL,
    geometry    GEOMETRY(Geometry, 4326) NOT NULL,
    area        DOUBLE PRECISION         NOT NULL,
    parent_id   UUID REFERENCES growing_cultivation_areas (id),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    attributes  JSONB
);

CREATE INDEX idx_cultivation_areas_farm_ref ON growing_cultivation_areas (farm_ref_id);
CREATE INDEX idx_cultivation_areas_type ON growing_cultivation_areas (type);
CREATE INDEX idx_cultivation_areas_parent ON growing_cultivation_areas (parent_id);
CREATE INDEX idx_cultivation_areas_geometry ON growing_cultivation_areas USING GIST (geometry);

