-- Сезоны
CREATE TABLE growing_seasons
(
    id          UUID PRIMARY KEY,
    name        TEXT                     NOT NULL,
    start_date  DATE                     NOT NULL,
    end_date    DATE                     NOT NULL,
    status      TEXT                     NOT NULL DEFAULT 'planning',
    created_by  TEXT                     NOT NULL,
    is_active   BOOLEAN                           DEFAULT FALSE,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_seasons_status ON growing_seasons (status);
CREATE INDEX idx_seasons_dates ON growing_seasons (start_date, end_date);