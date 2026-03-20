-- Операции в цикле выращивания
CREATE TABLE cycle_operations
(
    id           TEXT PRIMARY KEY,
    cycle_id     TEXT                     NOT NULL REFERENCES crop_cycles (id) ON DELETE CASCADE,
    type         TEXT                     NOT NULL,
    description  TEXT,
    amount       DOUBLE PRECISION,
    unit         TEXT,
    performed_by TEXT,
    notes        TEXT,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_cycle_operations_cycle ON cycle_operations (cycle_id);
CREATE INDEX idx_cycle_operations_type ON cycle_operations (type);