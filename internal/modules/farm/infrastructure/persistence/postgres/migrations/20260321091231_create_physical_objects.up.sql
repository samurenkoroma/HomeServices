CREATE TABLE farm_physical_objects
(
    id          UUID PRIMARY KEY,
    type        TEXT NOT NULL,

    name        TEXT NOT NULL,
    geometry    geometry(Geometry, 4326) NOT NULL ,
    area        NUMERIC ,

    status      TEXT NOT NULL DEFAULT 'active',
    owner_id    uuid REFERENCES auth_organizations (id) ON DELETE CASCADE ,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    delete_at   TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    attributes  JSONB
);

CREATE INDEX idx_objects_type ON farm_physical_objects(type);
CREATE INDEX idx_objects_geometry ON farm_physical_objects USING GIST(geometry);
CREATE INDEX idx_objects_status ON farm_physical_objects(status);