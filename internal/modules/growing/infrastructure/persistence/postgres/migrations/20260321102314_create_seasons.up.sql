-- Сезоны
CREATE TABLE seasons
(
    id          TEXT PRIMARY KEY,
    name        TEXT                     NOT NULL,
    start_date  DATE                     NOT NULL,
    end_date    DATE                     NOT NULL,
    status      TEXT                     NOT NULL DEFAULT 'planning',
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_seasons_status ON seasons (status);
CREATE INDEX idx_seasons_dates ON seasons (start_date, end_date);