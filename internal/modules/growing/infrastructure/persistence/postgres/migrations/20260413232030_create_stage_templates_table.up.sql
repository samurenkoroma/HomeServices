-- ============================================
-- Таблица шаблонов этапов (stage_templates)
-- ============================================

CREATE TABLE IF NOT EXISTS growing_stage_templates (
                                               id SERIAL PRIMARY KEY,
                                               species_key VARCHAR(50) NOT NULL REFERENCES public.growing_crops(key) ON DELETE CASCADE,

    -- Данные этапа
                                               type VARCHAR(50) NOT NULL,
                                               name VARCHAR(200) NOT NULL,
                                               bbch_start INTEGER NOT NULL,
                                               bbch_end INTEGER NOT NULL,
                                               description TEXT,
                                               priority VARCHAR(20) DEFAULT 'medium',
                                               is_required BOOLEAN DEFAULT true,

    -- Сортировка
                                               display_order INTEGER DEFAULT 0,

    -- Системные поля
                                               created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                                               updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),

    -- Уникальность: для одного вида не может быть двух этапов с одинаковым типом и порядком
                                               CONSTRAINT unique_species_stage_type UNIQUE (species_key, type, display_order)
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_stage_templates_species_key ON growing_stage_templates(species_key);
CREATE INDEX IF NOT EXISTS idx_stage_templates_bbch_range ON growing_stage_templates(bbch_start, bbch_end);

COMMENT ON TABLE  growing_stage_templates IS 'Шаблоны этапов для видов культур';
COMMENT ON COLUMN growing_stage_templates.type IS 'Тип этапа: soil_preparation, sowing, fertilization, protection, irrigation, pruning, harvest';
COMMENT ON COLUMN growing_stage_templates.bbch_start IS 'Начало BBCH диапазона (включительно)';
COMMENT ON COLUMN growing_stage_templates.bbch_end IS 'Конец BBCH диапазона (включительно)';
COMMENT ON COLUMN growing_stage_templates.priority IS 'Приоритет: low, medium, high, urgent';
COMMENT ON COLUMN growing_stage_templates.is_required IS 'Обязательный ли этап';