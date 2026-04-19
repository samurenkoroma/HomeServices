-- Сезоны
CREATE TABLE growing_seasons
(
    id          UUID PRIMARY KEY,
    name        TEXT NOT NULL,
    start_date  DATE NOT NULL,
    end_date    DATE NOT NULL,
    status      TEXT NOT NULL            DEFAULT 'planning',
    created_by  UUID NOT NULL,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX idx_seasons_status ON growing_seasons (status);
CREATE INDEX idx_seasons_dates ON growing_seasons (start_date, end_date);