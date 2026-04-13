-- ============================================
-- Таблица сортов (varieties)
-- ============================================

CREATE TABLE IF NOT EXISTS growing_varieties
(
    id                  VARCHAR(100) PRIMARY KEY,
    name                VARCHAR(200)  NOT NULL,
    species_key         VARCHAR(50)   NOT NULL REFERENCES growing_species (key) ON DELETE CASCADE,
    species_name        VARCHAR(100)  NOT NULL,

    -- Температурные параметры
    base_temperature    DECIMAL(5, 2) NOT NULL   DEFAULT 10.0,
    max_temperature     DECIMAL(5, 2) NOT NULL   DEFAULT 30.0,

    -- Агрономические параметры
    days_to_maturity    INTEGER       NOT NULL,
    yield_potential     DECIMAL(6, 2) NOT NULL,
    plant_height        DECIMAL(5, 2),

    -- Рекомендации
    recommended_seasons TEXT[]                   DEFAULT '{}',
    growing_types       TEXT[]                   DEFAULT '{}',

    -- Характеристики (JSONB для гибкости)
    characteristics     JSONB                    DEFAULT '{}',

    -- Описание
    description         TEXT,

    -- Водные требования (JSONB)
    water_requirement   JSONB                    DEFAULT '{}',

    -- Световые требования (JSONB)
    light_requirement   JSONB                    DEFAULT '{}',

    -- Фазы развития (JSONB)
    phenophase_gdd      JSONB                    DEFAULT '[]',

    -- Нормы высева (JSONB)
    seeding_rates       JSONB                    DEFAULT '{}',

    -- Системные поля
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at          TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Комментарии
COMMENT ON TABLE  growing_varieties IS 'Сорта сельскохозяйственных культур';
COMMENT ON COLUMN growing_varieties.id IS 'Уникальный идентификатор сорта';
COMMENT ON COLUMN growing_varieties.name IS 'Название сорта';
COMMENT ON COLUMN growing_varieties.species_key IS 'Ссылка на вид культуры';
COMMENT ON COLUMN growing_varieties.base_temperature IS 'Базовая температура для GDD расчета (°C)';
COMMENT ON COLUMN growing_varieties.max_temperature IS 'Максимальная температура для GDD расчета (°C)';
COMMENT ON COLUMN growing_varieties.days_to_maturity IS 'Дней до созревания';
COMMENT ON COLUMN growing_varieties.yield_potential IS 'Потенциальная урожайность (кг/м²)';
COMMENT ON COLUMN growing_varieties.recommended_seasons IS 'Рекомендуемые сезоны для посадки';
COMMENT ON COLUMN growing_varieties.growing_types IS 'Типы выращивания (open_ground, greenhouse)';
COMMENT ON COLUMN growing_varieties.water_requirement IS 'Потребность в воде по фазам (JSON)';
COMMENT ON COLUMN growing_varieties.light_requirement IS 'Потребность в освещении (JSON)';
COMMENT ON COLUMN growing_varieties.phenophase_gdd IS 'Фазы развития с GDD требованиями (JSON)';
COMMENT ON COLUMN growing_varieties.seeding_rates IS 'Нормы высева для разных типов выращивания (JSON)';

-- Индексы
CREATE INDEX IF NOT EXISTS idx_varieties_species_key ON growing_varieties (species_key);
CREATE INDEX IF NOT EXISTS idx_varieties_days_to_maturity ON growing_varieties (days_to_maturity);
CREATE INDEX IF NOT EXISTS idx_varieties_growing_types ON growing_varieties USING GIN (growing_types);
CREATE INDEX IF NOT EXISTS idx_varieties_recommended_seasons ON growing_varieties USING GIN (recommended_seasons);

-- JSONB индексы для эффективного поиска
CREATE INDEX IF NOT EXISTS idx_varieties_characteristics ON growing_varieties USING GIN (characteristics);
CREATE INDEX IF NOT EXISTS idx_varieties_water_requirement ON growing_varieties USING GIN (water_requirement);
CREATE INDEX IF NOT EXISTS idx_varieties_light_requirement ON growing_varieties USING GIN (light_requirement);
CREATE INDEX IF NOT EXISTS idx_varieties_phenophase_gdd ON growing_varieties USING GIN (phenophase_gdd);