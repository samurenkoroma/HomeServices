-- Добавляем поля для сортов и сезонов в crop_plans
ALTER TABLE crop_plans
    ADD COLUMN species_key VARCHAR(100),
    ADD COLUMN variety_id VARCHAR(100),
    ADD COLUMN crop_family VARCHAR(50),

    -- Поля сезона
    ADD COLUMN season_name VARCHAR(100),
    ADD COLUMN season_start DATE,
    ADD COLUMN season_end DATE,
    ADD COLUMN planting_date DATE,

    -- Агрономические данные
    ADD COLUMN seeds_planted INTEGER DEFAULT 0,
    ADD COLUMN expected_yield NUMERIC(10, 2) DEFAULT 0;

-- Индексы для поиска
CREATE INDEX idx_crop_plans_species ON crop_plans(species_key);
CREATE INDEX idx_crop_plans_season ON crop_plans(season_start, season_end);

COMMENT ON COLUMN crop_plans.species_key IS 'Ключ вида (tomato, eggplant, ...)';
COMMENT ON COLUMN crop_plans.variety_id IS 'ID сорта';
COMMENT ON COLUMN crop_plans.crop_family IS 'Семейство для севооборота';