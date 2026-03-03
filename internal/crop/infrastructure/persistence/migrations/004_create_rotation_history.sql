-- Таблица истории севооборота
CREATE TABLE rotation_history
(
    id           SERIAL PRIMARY KEY,
    bed_id       VARCHAR(36)              NOT NULL,

    crop_name    VARCHAR(255)             NOT NULL,
    crop_family  VARCHAR(50)              NOT NULL,
    species_key  VARCHAR(100),
    variety_id   VARCHAR(100),

    -- Даты
    planted_at   DATE                     NOT NULL,
    harvested_at DATE,
    season_name  VARCHAR(50),
    year         INTEGER                  NOT NULL,

    -- Дополнительно
    notes        TEXT,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    -- Связи
    CONSTRAINT fk_rotation_history_bed FOREIGN KEY (bed_id) REFERENCES land_structure (id),
    CONSTRAINT uk_rotation_history UNIQUE (bed_id, planted_at)
);

CREATE INDEX idx_rotation_history_bed_id ON rotation_history (bed_id);
CREATE INDEX idx_rotation_history_year ON rotation_history (year);
CREATE INDEX idx_rotation_history_family ON rotation_history (crop_family);

COMMENT ON TABLE rotation_history IS 'История посадок для севооборота';