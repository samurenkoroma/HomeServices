-- Операции в цикле
CREATE TABLE growing_cycle_operations
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    cycle_id     UUID                     NOT NULL REFERENCES growing_crop_cycles (id) ON DELETE CASCADE,
    type         TEXT                     NOT NULL,
    description  TEXT,
    amount       DOUBLE PRECISION,
    unit         TEXT,
    performed_by TEXT,
    notes        TEXT,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT check_type CHECK (type IN
                                 ('planting', 'watering', 'fertilizing', 'pest_control', 'weeding', 'harvesting',
                                  'other'))
);

CREATE INDEX idx_cycle_operations_cycle ON growing_cycle_operations (cycle_id);
CREATE INDEX idx_cycle_operations_type ON growing_cycle_operations (type);