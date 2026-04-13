-- ============================================
-- Таблица видов культур (species)
-- ============================================

CREATE TABLE IF NOT EXISTS growing_species
(
    key         VARCHAR(50) PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    family      VARCHAR(50)  NOT NULL,
    category    VARCHAR(50)  NOT NULL,
    image_url   TEXT,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- Комментарии к таблице и полям
COMMENT ON TABLE  growing_species IS 'Виды сельскохозяйственных культур';
COMMENT ON COLUMN growing_species.key IS 'Уникальный ключ вида (tomato, cucumber, etc.)';
COMMENT ON COLUMN growing_species.name IS 'Название вида на русском';
COMMENT ON COLUMN growing_species.family IS 'Ботаническое семейство';
COMMENT ON COLUMN growing_species.category IS 'Категория: Овощные, Зерновые, Бобовые, Масличные, Зеленные';
COMMENT ON COLUMN growing_species.image_url IS 'URL изображения культуры';
COMMENT ON COLUMN growing_species.description IS 'Описание культуры';

-- Индексы
CREATE INDEX IF NOT EXISTS idx_species_category ON growing_species (category);
CREATE INDEX IF NOT EXISTS idx_species_family ON growing_species (family);