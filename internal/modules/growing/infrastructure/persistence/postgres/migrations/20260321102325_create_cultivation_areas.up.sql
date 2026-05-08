-- Места выращивания
CREATE TABLE growing_cultivation_areas
(
    id          UUID PRIMARY KEY,
    farm_ref_id UUID                     NOT NULL,
    type        TEXT                     NOT NULL CHECK (type IN ('field', 'bed')),
    name        TEXT                     NOT NULL,
    area        NUMERIC(10, 2)           NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    attributes  JSONB
);


