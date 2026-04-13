-- ============================================
-- Вставка шаблонов этапов для всех видов
-- ============================================

-- ТОМАТ (tomato)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('tomato', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение базовых удобрений', 'high', true, 1),
                                                                                                                                   ('tomato', 'sowing', 'Посев семян', 0, 9, 'Посев семян в грунт или на рассаду', 'high', true, 2),
                                                                                                                                   ('tomato', 'pruning', 'Пикировка', 10, 19, 'Рассаживание растений по отдельным емкостям', 'medium', false, 3),
                                                                                                                                   ('tomato', 'fertilization', 'Подкормка азотом', 19, 39, 'Внесение азотных удобрений для роста', 'medium', true, 4),
                                                                                                                                   ('tomato', 'pruning', 'Формирование куста', 30, 39, 'Удаление пасынков, формирование стебля', 'medium', true, 5),
                                                                                                                                   ('tomato', 'protection', 'Обработка от вредителей', 50, 69, 'Опрыскивание от тли, клещей и других вредителей', 'high', true, 6),
                                                                                                                                   ('tomato', 'fertilization', 'Калийная подкормка', 70, 79, 'Внесение калийных удобрений для плодоношения', 'high', true, 7),
                                                                                                                                   ('tomato', 'harvest', 'Сбор урожая', 80, 89, 'Сбор спелых плодов', 'high', true, 8);

-- ОГУРЕЦ (cucumber)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('cucumber', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение удобрений', 'high', true, 1),
                                                                                                                                   ('cucumber', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('cucumber', 'fertilization', 'Подкормка', 19, 39, 'Комплексная подкормка', 'medium', true, 3),
                                                                                                                                   ('cucumber', 'protection', 'Обработка от вредителей', 50, 69, 'Защита от вредителей', 'high', true, 4),
                                                                                                                                   ('cucumber', 'harvest', 'Сбор урожая', 80, 89, 'Сбор плодов', 'high', true, 5);

-- КАПУСТА (cabbage)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('cabbage', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение удобрений', 'high', true, 1),
                                                                                                                                   ('cabbage', 'sowing', 'Посев на рассаду', 0, 9, 'Посев семян на рассаду', 'high', true, 2),
                                                                                                                                   ('cabbage', 'fertilization', 'Подкормка азотом', 19, 39, 'Внесение азотных удобрений', 'medium', true, 3),
                                                                                                                                   ('cabbage', 'protection', 'Обработка от вредителей', 50, 69, 'Защита от крестоцветной блошки', 'high', true, 4),
                                                                                                                                   ('cabbage', 'harvest', 'Сбор урожая', 80, 89, 'Срезка кочанов', 'high', true, 5);

-- КАРТОФЕЛЬ (potato)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('potato', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, нарезка гребней', 'high', true, 1),
                                                                                                                                   ('potato', 'sowing', 'Посадка клубней', 0, 9, 'Посадка клубней в почву', 'high', true, 2),
                                                                                                                                   ('potato', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('potato', 'protection', 'Обработка от колорадского жука', 50, 69, 'Защита от колорадского жука', 'high', true, 4),
                                                                                                                                   ('potato', 'harvest', 'Уборка урожая', 80, 89, 'Выкопка клубней', 'high', true, 5);

-- КУКУРУЗА (corn)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('corn', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, боронование', 'high', true, 1),
                                                                                                                                   ('corn', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('corn', 'fertilization', 'Подкормка азотом', 19, 39, 'Внесение азотных удобрений', 'medium', true, 3),
                                                                                                                                   ('corn', 'harvest', 'Уборка урожая', 80, 89, 'Сбор початков', 'high', true, 4);

-- РЕДИС (radish)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('radish', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Рыхление почвы', 'medium', true, 1),
                                                                                                                                   ('radish', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('radish', 'harvest', 'Уборка урожая', 80, 89, 'Сбор корнеплодов', 'high', true, 3);

-- ЛУК (onion)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('onion', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение удобрений', 'high', true, 1),
                                                                                                                                   ('onion', 'sowing', 'Посев севка', 0, 9, 'Посадка лука-севка', 'high', true, 2),
                                                                                                                                   ('onion', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('onion', 'harvest', 'Уборка урожая', 80, 89, 'Сбор луковиц', 'high', true, 4);

-- ЧЕСНОК (garlic)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('garlic', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение удобрений', 'high', true, 1),
                                                                                                                                   ('garlic', 'sowing', 'Посадка зубков', 0, 9, 'Посадка зубков чеснока', 'high', true, 2),
                                                                                                                                   ('garlic', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('garlic', 'harvest', 'Уборка урожая', 80, 89, 'Выкопка чеснока', 'high', true, 4);

-- ПШЕНИЦА (wheat)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('wheat', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, культивация', 'high', true, 1),
                                                                                                                                   ('wheat', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('wheat', 'fertilization', 'Подкормка азотом', 19, 39, 'Внесение азотных удобрений', 'medium', true, 3),
                                                                                                                                   ('wheat', 'protection', 'Обработка от сорняков', 30, 39, 'Гербицидная обработка', 'medium', true, 4),
                                                                                                                                   ('wheat', 'harvest', 'Уборка урожая', 80, 89, 'Комбайнирование', 'high', true, 5);

-- ЯЧМЕНЬ (barley)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('barley', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, боронование', 'high', true, 1),
                                                                                                                                   ('barley', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('barley', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('barley', 'harvest', 'Уборка урожая', 80, 89, 'Комбайнирование', 'high', true, 4);

-- ПОДСОЛНЕЧНИК (sunflower)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('sunflower', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, культивация', 'high', true, 1),
                                                                                                                                   ('sunflower', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('sunflower', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('sunflower', 'harvest', 'Уборка урожая', 80, 89, 'Сбор корзинок', 'high', true, 4);

-- СОЯ (soybean)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('soybean', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, культивация', 'high', true, 1),
                                                                                                                                   ('soybean', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('soybean', 'protection', 'Обработка от сорняков', 19, 39, 'Гербицидная обработка', 'medium', true, 3),
                                                                                                                                   ('soybean', 'harvest', 'Уборка урожая', 80, 89, 'Сбор бобов', 'high', true, 4);

-- ГОРОХ (pea)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('pea', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, культивация', 'high', true, 1),
                                                                                                                                   ('pea', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('pea', 'harvest', 'Уборка урожая', 80, 89, 'Сбор бобов', 'high', true, 3);

-- РАПС (rapeseed)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('rapeseed', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, культивация', 'high', true, 1),
                                                                                                                                   ('rapeseed', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('rapeseed', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('rapeseed', 'harvest', 'Уборка урожая', 80, 89, 'Комбайнирование', 'high', true, 4);

-- БАКЛАЖАН (eggplant)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('eggplant', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Вспашка, внесение удобрений', 'high', true, 1),
                                                                                                                                   ('eggplant', 'sowing', 'Посев', 0, 9, 'Посев семян на рассаду', 'high', true, 2),
                                                                                                                                   ('eggplant', 'fertilization', 'Подкормка', 19, 39, 'Внесение удобрений', 'medium', true, 3),
                                                                                                                                   ('eggplant', 'protection', 'Обработка от вредителей', 50, 69, 'Защита от колорадского жука', 'high', true, 4),
                                                                                                                                   ('eggplant', 'harvest', 'Уборка урожая', 80, 89, 'Сбор плодов', 'high', true, 5);

-- УКРОП (dill)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('dill', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Рыхление почвы', 'low', true, 1),
                                                                                                                                   ('dill', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('dill', 'harvest', 'Срезка зелени', 19, 39, 'Срезка листьев', 'high', true, 3);

-- ПЕТРУШКА (parsley)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('parsley', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Рыхление почвы', 'low', true, 1),
                                                                                                                                   ('parsley', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('parsley', 'harvest', 'Срезка зелени', 19, 39, 'Срезка листьев', 'high', true, 3);

-- САЛАТ (lettuce)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('lettuce', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Рыхление почвы', 'low', true, 1),
                                                                                                                                   ('lettuce', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('lettuce', 'harvest', 'Срезка листьев', 19, 39, 'Срезка розетки', 'high', true, 3);

-- БАЗИЛИК (basil)
INSERT INTO public.growing_stage_templates (species_key, type, name, bbch_start, bbch_end, description, priority, is_required, display_order) VALUES
                                                                                                                                   ('basil', 'soil_preparation', 'Подготовка почвы', 0, 9, 'Рыхление почвы', 'low', true, 1),
                                                                                                                                   ('basil', 'sowing', 'Посев', 0, 9, 'Посев семян', 'high', true, 2),
                                                                                                                                   ('basil', 'harvest', 'Срезка зелени', 19, 39, 'Срезка листьев', 'high', true, 3);

-- Проверка
DO $$
    DECLARE
        template_count INTEGER;
    BEGIN
        SELECT COUNT(*) INTO template_count FROM public.growing_stage_templates;
        RAISE NOTICE 'Добавлено шаблонов этапов: %', template_count;
    END $$;