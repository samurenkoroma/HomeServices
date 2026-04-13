-- ============================================
-- ИМПОРТ СОРТОВ (varieties)
-- Всего 54 сорта для 19 видов культур
-- ============================================

-- ============================================
-- 1. ТОМАТЫ (tomato) - 9 сортов
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('incas_f1', 'Инкас F1', 'tomato', 'Томат', 10.0, 30.0, 95, 12.0, 1.2, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "180-200g", "fruit_color": "красный", "fruit_shape": "округлый", "type": "индетерминантный", "use": "салатный", "resistance": "VTM, F1, Cladosporium"}',
        'Раннеспелый гибрид для пленочных теплиц и открытого грунта',
        '{"daily_need_min": 2.0, "daily_need_opt": 4.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 50000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "name": "Первый настоящий лист", "gdd_required": 100, "description": "Появление первого настоящего листа", "is_critical": false}, {"code": "BBCH-19", "name": "9 и более листьев", "gdd_required": 280, "description": "Активный рост вегетативной массы", "is_critical": false}, {"code": "BBCH-51", "name": "Бутонизация", "gdd_required": 480, "description": "Появление бутонов", "is_critical": false}, {"code": "BBCH-61", "name": "Цветение", "gdd_required": 650, "description": "Раскрытие цветов", "is_critical": true}, {"code": "BBCH-71", "name": "Завязывание плодов", "gdd_required": 780, "description": "Образование завязей", "is_critical": true}, {"code": "BBCH-81", "name": "Созревание", "gdd_required": 950, "description": "Плоды начинают окрашиваться", "is_critical": false}, {"code": "BBCH-89", "name": "Полная спелость", "gdd_required": 1100, "description": "Плоды достигли полной спелости", "is_critical": false}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.4, "sowing_depth": 1.5, "germination_rate": 90, "safety_factor": 1.15}, "greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.8, "plant_spacing": 0.5, "sowing_depth": 1.0, "germination_rate": 92, "safety_factor": 1.1}}'),

       ('rio_grande', 'Рио-Гранде', 'tomato', 'Томат', 10.0, 32.0, 110, 10.0, 0.8, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_weight": "100-120g", "fruit_color": "красный", "fruit_shape": "сливовидный", "type": "детерминантный", "use": "для переработки", "resistance": "вертициллез, фузариоз"}',
        'Засухоустойчивый сорт для открытого грунта и фермерских хозяйств',
        '{"daily_need_min": 1.5, "daily_need_opt": 3.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 35000, "lux_max": 55000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 110}, {"code": "BBCH-19", "gdd_required": 320}, {"code": "BBCH-51", "gdd_required": 550}, {"code": "BBCH-61", "gdd_required": 730, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 880, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1100}, {"code": "BBCH-89", "gdd_required": 1300}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.35, "sowing_depth": 2.0, "germination_rate": 85, "safety_factor": 1.2}}'),

       ('bella_rosa_f1', 'Белла Роса F1', 'tomato', 'Томат', 8.0, 28.0, 85, 9.0, 1.0, ARRAY ['spring'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "80-100g", "fruit_color": "розовый", "fruit_shape": "округлый", "type": "детерминантный", "use": "салатный", "cold_resistant": "да"}',
        'Холодостойкий гибрид для раннего выращивания в открытом грунте',
        '{"daily_need_min": 1.8, "daily_need_opt": 3.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 90}, {"code": "BBCH-19", "gdd_required": 250}, {"code": "BBCH-51", "gdd_required": 430}, {"code": "BBCH-61", "gdd_required": 580, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 700, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 880}, {"code": "BBCH-89", "gdd_required": 1050}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.4, "sowing_depth": 1.5, "germination_rate": 88, "safety_factor": 1.15}}'),

       ('solerossa_f1', 'Солероссо F1', 'tomato', 'Томат', 10.0, 32.0, 100, 11.0, 1.1, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_weight": "150-170g", "fruit_color": "красный", "fruit_shape": "округлый", "type": "индетерминантный", "use": "универсальный", "salt_resistant": "да"}',
        'Солеустойчивый гибрид для выращивания на засоленных почвах',
        '{"daily_need_min": 2.0, "daily_need_opt": 3.8, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 32000, "lux_max": 52000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 105}, {"code": "BBCH-19", "gdd_required": 300}, {"code": "BBCH-51", "gdd_required": 520}, {"code": "BBCH-61", "gdd_required": 690, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 830, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1050}, {"code": "BBCH-89", "gdd_required": 1250}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.45, "sowing_depth": 1.8, "germination_rate": 87, "safety_factor": 1.18}}'),

       ('macan_f1', 'Макан F1', 'tomato', 'Томат', 10.0, 30.0, 115, 14.0, 1.8, ARRAY ['spring', 'summer'],
        ARRAY ['greenhouse'],
        '{"fruit_weight": "250-300g", "fruit_color": "красный", "fruit_shape": "плоскоокруглый", "type": "индетерминантный", "use": "салатный"}',
        'Крупноплодный гибрид для защищенного грунта',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 35000, "lux_max": 50000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 115}, {"code": "BBCH-19", "gdd_required": 330}, {"code": "BBCH-51", "gdd_required": 570}, {"code": "BBCH-61", "gdd_required": 760, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 910, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1150}, {"code": "BBCH-89", "gdd_required": 1380}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.9, "plant_spacing": 0.6, "sowing_depth": 1.2, "germination_rate": 91, "safety_factor": 1.1}}'),

       ('yamamoto', 'Ямамото', 'tomato', 'Томат', 10.0, 28.0, 108, 13.0, 1.6, ARRAY ['spring', 'summer'],
        ARRAY ['greenhouse'],
        '{"fruit_weight": "200-250g", "fruit_color": "темно-красный", "fruit_shape": "округлый", "type": "индетерминантный", "use": "салатный", "origin": "Япония"}',
        'Японский сорт с высокими вкусовыми качествами',
        '{"daily_need_min": 2.2, "daily_need_opt": 4.2, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 48000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 108}, {"code": "BBCH-19", "gdd_required": 310}, {"code": "BBCH-51", "gdd_required": 540}, {"code": "BBCH-61", "gdd_required": 720, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 860, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1080}, {"code": "BBCH-89", "gdd_required": 1300}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.8, "plant_spacing": 0.55, "sowing_depth": 1.2, "germination_rate": 89, "safety_factor": 1.12}}'),

       ('bull_heart', 'Бычье сердце', 'tomato', 'Томат', 10.0, 30.0, 120, 8.5, 1.5, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "300-500g", "fruit_color": "малиновый", "fruit_shape": "сердцевидный", "type": "индетерминантный", "use": "салатный"}',
        'Крупноплодный салатный сорт',
        '{"daily_need_min": 2.0, "daily_need_opt": 4.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 50000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 120}, {"code": "BBCH-19", "gdd_required": 350}, {"code": "BBCH-51", "gdd_required": 600}, {"code": "BBCH-61", "gdd_required": 800, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 950, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1200}, {"code": "BBCH-89", "gdd_required": 1400}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.5, "sowing_depth": 2.0, "germination_rate": 85, "safety_factor": 1.2}, "greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.8, "plant_spacing": 0.6, "sowing_depth": 1.5, "germination_rate": 90, "safety_factor": 1.1}}'),

       ('persimmon', 'Хурма', 'tomato', 'Томат', 10.0, 30.0, 115, 7.0, 1.2, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "150-200g", "fruit_color": "оранжевый", "fruit_shape": "плоскоокруглый", "type": "детерминантный", "use": "салатный"}',
        'Сорт с оранжевыми плодами, напоминающими хурму',
        '{"daily_need_min": 1.8, "daily_need_opt": 3.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 28000, "lux_max": 48000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 112}, {"code": "BBCH-19", "gdd_required": 325}, {"code": "BBCH-51", "gdd_required": 560}, {"code": "BBCH-61", "gdd_required": 745, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 890, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1120}, {"code": "BBCH-89", "gdd_required": 1330}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.5, "sowing_depth": 2.0, "germination_rate": 84, "safety_factor": 1.22}}'),

       ('black_bunch', 'Черная гроздь', 'tomato', 'Томат', 10.0, 30.0, 95, 6.0, 1.5,
        ARRAY ['spring', 'summer', 'autumn'], ARRAY ['greenhouse'],
        '{"fruit_weight": "30-40g", "fruit_color": "черно-коричневый", "fruit_shape": "сливовидный", "type": "индетерминантный", "use": "черри"}',
        'Черри с необычным цветом и вкусом',
        '{"daily_need_min": 1.5, "daily_need_opt": 3.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 50000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 95}, {"code": "BBCH-19", "gdd_required": 270}, {"code": "BBCH-51", "gdd_required": 460}, {"code": "BBCH-61", "gdd_required": 620, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 750, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 950}, {"code": "BBCH-89", "gdd_required": 1120}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.7, "plant_spacing": 0.4, "sowing_depth": 1.2, "germination_rate": 88, "safety_factor": 1.13}}');


-- ============================================
-- 2. ОГУРЦЫ (cucumber) - 9 сортов
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('rodnichok_f1', 'Родничок F1', 'cucumber', 'Огурец', 12.0, 35.0, 48, 10.0, 1.8, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_length": "8-10cm", "fruit_color": "зеленый с полосками", "fruit_weight": "80-100g", "type": "пчелоопыляемый", "use": "засолочный"}',
        'Пчелоопыляемый сорт для открытого грунта, устойчив к болезням',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "name": "Всходы", "gdd_required": 90}, {"code": "BBCH-19", "name": "3-4 настоящих листа", "gdd_required": 200}, {"code": "BBCH-51", "name": "Бутонизация", "gdd_required": 300}, {"code": "BBCH-61", "name": "Цветение", "gdd_required": 380, "is_critical": true}, {"code": "BBCH-71", "name": "Плодоношение", "gdd_required": 480, "is_critical": true}, {"code": "BBCH-89", "name": "Техническая спелость", "gdd_required": 650}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.3, "sowing_depth": 2.0, "germination_rate": 85, "safety_factor": 1.15}}'),

       ('bochkovoi_f1', 'Бочковой F1', 'cucumber', 'Огурец', 12.0, 35.0, 50, 12.0, 2.0, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_length": "10-12cm", "fruit_color": "темно-зеленый", "fruit_weight": "100-120g", "type": "партенокарпический", "use": "засолочный"}',
        'Бочкового типа для засолки и маринования',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 92}, {"code": "BBCH-19", "gdd_required": 205}, {"code": "BBCH-51", "gdd_required": 310}, {"code": "BBCH-61", "gdd_required": 390, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 490, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 660}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.35, "sowing_depth": 2.0, "germination_rate": 86, "safety_factor": 1.14}}'),

       ('nezhinsky', 'Нежинский', 'cucumber', 'Огурец', 12.0, 35.0, 55, 9.0, 1.5, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_length": "10-14cm", "fruit_color": "зеленый", "fruit_weight": "90-110g", "type": "пчелоопыляемый", "use": "универсальный", "origin": "Украина"}',
        'Классический сорт для открытого грунта',
        '{"daily_need_min": 3.0, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 95}, {"code": "BBCH-19", "gdd_required": 210}, {"code": "BBCH-51", "gdd_required": 320}, {"code": "BBCH-61", "gdd_required": 400, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 500, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 680}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.3, "sowing_depth": 2.5, "germination_rate": 84, "safety_factor": 1.16}}'),

       ('crispina_f1', 'Криспина F1', 'cucumber', 'Огурец', 12.0, 35.0, 42, 14.0, 2.2, ARRAY ['spring', 'summer'],
        ARRAY ['greenhouse'],
        '{"fruit_length": "12-14cm", "fruit_color": "темно-зеленый", "fruit_weight": "80-100g", "type": "партенокарпический", "use": "салатный"}',
        'Ранний гибрид для теплиц',
        '{"daily_need_min": 3.5, "daily_need_opt": 5.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 28000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 85}, {"code": "BBCH-19", "gdd_required": 190}, {"code": "BBCH-51", "gdd_required": 280}, {"code": "BBCH-61", "gdd_required": 360, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 450, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 600}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.9, "plant_spacing": 0.4, "sowing_depth": 1.5, "germination_rate": 92, "safety_factor": 1.08}}'),

       ('ecole_f1', 'Эколь F1', 'cucumber', 'Огурец', 12.0, 35.0, 45, 11.0, 1.9, ARRAY ['spring', 'summer', 'autumn'],
        ARRAY ['greenhouse'],
        '{"fruit_length": "14-16cm", "fruit_color": "зеленый", "fruit_weight": "100-120g", "type": "партенокарпический", "use": "салатный"}',
        'Урожайный гибрид для продленного оборота',
        '{"daily_need_min": 3.2, "daily_need_opt": 5.2, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 26000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 88}, {"code": "BBCH-19", "gdd_required": 195}, {"code": "BBCH-51", "gdd_required": 290}, {"code": "BBCH-61", "gdd_required": 370, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 465, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 630}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.8, "plant_spacing": 0.45, "sowing_depth": 1.5, "germination_rate": 91, "safety_factor": 1.09}}'),

       ('mertus_f1', 'Мертус F1', 'cucumber', 'Огурец', 12.0, 35.0, 47, 13.0, 2.0, ARRAY ['spring', 'summer'],
        ARRAY ['greenhouse'],
        '{"fruit_length": "10-12cm", "fruit_color": "ярко-зеленый", "fruit_weight": "80-90g", "type": "партенокарпический", "use": "универсальный"}',
        'Гладкоплодный гибрид для теплиц',
        '{"daily_need_min": 3.3, "daily_need_opt": 5.3, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 27000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 90}, {"code": "BBCH-19", "gdd_required": 200}, {"code": "BBCH-51", "gdd_required": 300}, {"code": "BBCH-61", "gdd_required": 380, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 475, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 645}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.85, "plant_spacing": 0.4, "sowing_depth": 1.5, "germination_rate": 90, "safety_factor": 1.1}}'),

       ('madrilene_f1', 'Мадрилене F1', 'cucumber', 'Огурец', 12.0, 35.0, 44, 12.5, 2.1, ARRAY ['spring', 'summer'],
        ARRAY ['greenhouse'],
        '{"fruit_length": "16-18cm", "fruit_color": "темно-зеленый", "fruit_weight": "120-140g", "type": "партенокарпический", "use": "салатный"}',
        'Длинноплодный гибрид для салатов',
        '{"daily_need_min": 3.4, "daily_need_opt": 5.4, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 26000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 87}, {"code": "BBCH-19", "gdd_required": 193}, {"code": "BBCH-51", "gdd_required": 285}, {"code": "BBCH-61", "gdd_required": 365, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 460, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 620}]',
        '{"greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.9, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 89, "safety_factor": 1.11}}'),

       ('hector_f1', 'Гектор F1', 'cucumber', 'Огурец', 12.0, 35.0, 40, 8.0, 1.2, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_length": "6-8cm", "fruit_color": "зеленый", "fruit_weight": "60-80g", "type": "партенокарпический", "use": "засолочный"}',
        'Ультраранний гибрид для открытого грунта',
        '{"daily_need_min": 2.8, "daily_need_opt": 4.8, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 24000, "lux_max": 42000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 80}, {"code": "BBCH-19", "gdd_required": 180}, {"code": "BBCH-51", "gdd_required": 270}, {"code": "BBCH-61", "gdd_required": 340, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 430, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 580}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.25, "sowing_depth": 2.0, "germination_rate": 87, "safety_factor": 1.17}}'),

       ('ajax_f1', 'Аякс F1', 'cucumber', 'Огурец', 12.0, 35.0, 46, 10.5, 1.7, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"fruit_length": "10-12cm", "fruit_color": "зеленый", "fruit_weight": "90-110g", "type": "пчелоопыляемый", "use": "универсальный"}',
        'Устойчивый к стрессам гибрид для открытого грунта',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 12, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 91}, {"code": "BBCH-19", "gdd_required": 202}, {"code": "BBCH-51", "gdd_required": 305}, {"code": "BBCH-61", "gdd_required": 385, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 485, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 655}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.35, "sowing_depth": 2.0, "germination_rate": 85, "safety_factor": 1.15}}');


-- ============================================
-- 3. КАПУСТА (cabbage) - 11 сортов
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('slava_1305', 'Слава 1305', 'cabbage', 'Капуста', 5.0, 25.0, 110, 7.0, 0.4, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "2-3kg", "head_color": "светло-зеленый", "head_shape": "округлый", "use": "свежий, квашение", "origin": "СССР"}',
        'Классический среднеспелый сорт для квашения',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 150}, {"code": "BBCH-19", "gdd_required": 400}, {"code": "BBCH-51", "gdd_required": 700}, {"code": "BBCH-61", "gdd_required": 900, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1100, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1400}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 80, "safety_factor": 1.25}}'),

       ('megaton_f1', 'Мегатон F1', 'cabbage', 'Капуста', 5.0, 25.0, 120, 12.0, 0.45, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "4-6kg", "head_color": "зеленый", "head_shape": "округлый", "use": "свежий, хранение"}',
        'Крупноплодный гибрид для длительного хранения',
        '{"daily_need_min": 3.5, "daily_need_opt": 5.5, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 22000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 160}, {"code": "BBCH-19", "gdd_required": 430}, {"code": "BBCH-51", "gdd_required": 750}, {"code": "BBCH-61", "gdd_required": 950, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1150, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1500}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.6, "sowing_depth": 1.5, "germination_rate": 82, "safety_factor": 1.22}}'),

       ('podarok', 'Подарок', 'cabbage', 'Капуста', 5.0, 25.0, 115, 8.5, 0.4, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "3-4kg", "head_color": "светло-зеленый", "head_shape": "плоскоокруглый", "use": "универсальный"}',
        'Хорошо хранится, устойчив к растрескиванию',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 155}, {"code": "BBCH-19", "gdd_required": 420}, {"code": "BBCH-51", "gdd_required": 730}, {"code": "BBCH-61", "gdd_required": 930, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1130, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1450}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 81, "safety_factor": 1.23}}'),

       ('menza_f1', 'Менза F1', 'cabbage', 'Капуста', 5.0, 25.0, 125, 11.0, 0.5, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "4-5kg", "head_color": "сине-зеленый", "head_shape": "округлый", "use": "свежий, квашение"}',
        'Голландский гибрид с высокими вкусовыми качествами',
        '{"daily_need_min": 3.2, "daily_need_opt": 5.2, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 21000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 165}, {"code": "BBCH-19", "gdd_required": 440}, {"code": "BBCH-51", "gdd_required": 770}, {"code": "BBCH-61", "gdd_required": 970, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1170, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1520}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.55, "sowing_depth": 1.5, "germination_rate": 83, "safety_factor": 1.2}}'),

       ('krautman_amager', 'Краутман Амагер', 'cabbage', 'Капуста', 5.0, 25.0, 130, 9.0, 0.4,
        ARRAY ['spring', 'summer'], ARRAY ['open_ground'],
        '{"head_weight": "3-4kg", "head_color": "зеленый", "head_shape": "округлый", "use": "квашение"}',
        'Позднеспелый сорт для квашения',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 170}, {"code": "BBCH-19", "gdd_required": 450}, {"code": "BBCH-51", "gdd_required": 800}, {"code": "BBCH-61", "gdd_required": 1000, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1200, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1550}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 80, "safety_factor": 1.25}}'),

       ('agressor_f1', 'Агрессор F1', 'cabbage', 'Капуста', 5.0, 27.0, 105, 10.0, 0.4, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "3-4kg", "head_color": "зеленый", "head_shape": "округлый", "use": "универсальный"}',
        'Жаростойкий гибрид для южных регионов',
        '{"daily_need_min": 2.8, "daily_need_opt": 4.8, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 145}, {"code": "BBCH-19", "gdd_required": 390}, {"code": "BBCH-51", "gdd_required": 680}, {"code": "BBCH-61", "gdd_required": 880, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1080, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1380}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.65, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 84, "safety_factor": 1.19}}'),

       ('atria_f1', 'Атрия F1', 'cabbage', 'Капуста', 5.0, 25.0, 108, 11.5, 0.45, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "3-5kg", "head_color": "темно-зеленый", "head_shape": "округлый", "use": "свежий"}',
        'Гибрид для промышленного выращивания',
        '{"daily_need_min": 3.1, "daily_need_opt": 5.1, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 21000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 148}, {"code": "BBCH-19", "gdd_required": 395}, {"code": "BBCH-51", "gdd_required": 690}, {"code": "BBCH-61", "gdd_required": 890, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1090, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1420}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.65, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 85, "safety_factor": 1.18}}'),

       ('belosnezhka', 'Белоснежка', 'cabbage', 'Капуста', 5.0, 25.0, 100, 6.5, 0.35, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "1-2kg", "head_color": "белый", "head_shape": "округлый", "use": "салатный"}',
        'Раннеспелая капуста для салатов',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 18000, "lux_max": 30000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 140}, {"code": "BBCH-19", "gdd_required": 370}, {"code": "BBCH-51", "gdd_required": 650}, {"code": "BBCH-61", "gdd_required": 850, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1050, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1350}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.5, "plant_spacing": 0.4, "sowing_depth": 1.2, "germination_rate": 83, "safety_factor": 1.2}}'),

       ('storidor_f1', 'Сторидор F1', 'cabbage', 'Капуста', 5.0, 25.0, 112, 9.5, 0.42, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "3-4kg", "head_color": "зеленый", "head_shape": "округлый", "use": "хранение"}',
        'Гибрид для длительного хранения',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 152}, {"code": "BBCH-19", "gdd_required": 410}, {"code": "BBCH-51", "gdd_required": 720}, {"code": "BBCH-61", "gdd_required": 920, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1120, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1440}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.5, "sowing_depth": 1.5, "germination_rate": 82, "safety_factor": 1.22}}'),

       ('rinda_f1', 'Ринда F1', 'cabbage', 'Капуста', 5.0, 25.0, 100, 8.0, 0.4, ARRAY ['spring', 'summer', 'autumn'],
        ARRAY ['open_ground'],
        '{"head_weight": "2-3kg", "head_color": "зеленый", "head_shape": "округлый", "use": "универсальный"}',
        'Популярный гибрид для свежего потребления и переработки',
        '{"daily_need_min": 2.8, "daily_need_opt": 4.8, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 142}, {"code": "BBCH-19", "gdd_required": 380}, {"code": "BBCH-51", "gdd_required": 660}, {"code": "BBCH-61", "gdd_required": 860, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1060, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1360}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.6, "plant_spacing": 0.45, "sowing_depth": 1.5, "germination_rate": 84, "safety_factor": 1.19}}'),

       ('jintama_f1', 'Джинтама F1', 'cabbage', 'Капуста', 5.0, 25.0, 118, 10.5, 0.48, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"head_weight": "4-5kg", "head_color": "темно-зеленый", "head_shape": "округлый", "use": "свежий, хранение"}',
        'Крупноплодный гибрид японской селекции',
        '{"daily_need_min": 3.3, "daily_need_opt": 5.3, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 22000, "lux_max": 35000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 158}, {"code": "BBCH-19", "gdd_required": 425}, {"code": "BBCH-51", "gdd_required": 740}, {"code": "BBCH-61", "gdd_required": 940, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1140, "is_critical": true}, {"code": "BBCH-89", "gdd_required": 1480}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.55, "sowing_depth": 1.5, "germination_rate": 83, "safety_factor": 1.2}}');


-- ============================================
-- 4. КАРТОФЕЛЬ (potato) - 3 сорта
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('gala', 'Гала', 'potato', 'Картофель', 5.0, 28.0, 70, 4.5, 0.6, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"tuber_weight": "80-120g", "tuber_color": "желтый", "starch": "11-13%", "use": "столовый", "keeping_quality": "хорошая"}',
        'Раннеспелый сорт с желтой мякотью',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "short_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "name": "Всходы", "gdd_required": 80}, {"code": "BBCH-19", "name": "Смыкание рядов", "gdd_required": 200}, {"code": "BBCH-51", "name": "Бутонизация", "gdd_required": 350}, {"code": "BBCH-61", "name": "Цветение", "gdd_required": 450}, {"code": "BBCH-71", "name": "Рост клубней", "gdd_required": 600}, {"code": "BBCH-89", "name": "Созревание", "gdd_required": 800}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.3, "sowing_depth": 8.0, "germination_rate": 95, "safety_factor": 1.1}}'),

       ('nevsky', 'Невский', 'potato', 'Картофель', 5.0, 28.0, 85, 5.0, 0.65, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"tuber_weight": "90-130g", "tuber_color": "белый", "starch": "10-12%", "use": "столовый", "keeping_quality": "отличная"}',
        'Популярный сорт с отличной лежкостью',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "short_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 90}, {"code": "BBCH-19", "gdd_required": 220}, {"code": "BBCH-51", "gdd_required": 380}, {"code": "BBCH-61", "gdd_required": 480}, {"code": "BBCH-71", "gdd_required": 650}, {"code": "BBCH-89", "gdd_required": 900}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.35, "sowing_depth": 8.0, "germination_rate": 96, "safety_factor": 1.08}}'),

       ('lugovskoy', 'Луговской', 'potato', 'Картофель', 5.0, 28.0, 95, 5.5, 0.7, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"tuber_weight": "100-150g", "tuber_color": "розовый", "starch": "12-15%", "use": "столовый"}',
        'Урожайный сорт с розовыми клубнями',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "short_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 95}, {"code": "BBCH-19", "gdd_required": 230}, {"code": "BBCH-51", "gdd_required": 400}, {"code": "BBCH-61", "gdd_required": 500}, {"code": "BBCH-71", "gdd_required": 680}, {"code": "BBCH-89", "gdd_required": 950}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.3, "sowing_depth": 8.0, "germination_rate": 94, "safety_factor": 1.12}}');


-- ============================================
-- 5. КУКУРУЗА (corn) - 3 сорта
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('ladozhsky_101', 'Ладожский 101', 'corn', 'Кукуруза', 10.0, 35.0, 95, 8.0, 2.0, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'],
        '{"cob_weight": "200-250g", "grain_color": "желтый", "use": "свежий, консервирование", "sugar_content": "высокий"}',
        'Сахарный сорт для свежего потребления',
        '{"daily_need_min": 4.0, "daily_need_opt": 6.0, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 10000, "lux_opt": 40000, "lux_max": 65000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 100}, {"code": "BBCH-19", "gdd_required": 300}, {"code": "BBCH-51", "gdd_required": 550}, {"code": "BBCH-61", "gdd_required": 750}, {"code": "BBCH-71", "gdd_required": 950}, {"code": "BBCH-89", "gdd_required": 1200}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.25, "sowing_depth": 5.0, "germination_rate": 90, "safety_factor": 1.15}}'),

       ('kubanskaya_saharnaya', 'Кубанская сахарная', 'corn', 'Кукуруза', 10.0, 35.0, 105, 9.0, 2.2,
        ARRAY ['spring', 'summer'], ARRAY ['open_ground'],
        '{"cob_weight": "250-300g", "grain_color": "ярко-желтый", "use": "консервирование"}',
        'Высокоурожайный сахарный сорт',
        '{"daily_need_min": 4.0, "daily_need_opt": 6.0, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 10000, "lux_opt": 40000, "lux_max": 65000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 110}, {"code": "BBCH-19", "gdd_required": 320}, {"code": "BBCH-51", "gdd_required": 580}, {"code": "BBCH-61", "gdd_required": 780}, {"code": "BBCH-71", "gdd_required": 980}, {"code": "BBCH-89", "gdd_required": 1250}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.3, "sowing_depth": 5.0, "germination_rate": 90, "safety_factor": 1.15}}'),

       ('dobrynya', 'Добрыня', 'corn', 'Кукуруза', 10.0, 35.0, 100, 8.5, 1.9, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground'], '{"cob_weight": "220-270g", "grain_color": "золотистый", "use": "свежий"}',
        'Раннеспелый гибрид для всех регионов',
        '{"daily_need_min": 4.0, "daily_need_opt": 6.0, "critical_phases": ["BBCH-51", "BBCH-61", "BBCH-71"]}',
        '{"lux_min": 10000, "lux_opt": 38000, "lux_max": 60000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 105}, {"code": "BBCH-19", "gdd_required": 310}, {"code": "BBCH-51", "gdd_required": 560}, {"code": "BBCH-61", "gdd_required": 760}, {"code": "BBCH-71", "gdd_required": 960}, {"code": "BBCH-89", "gdd_required": 1220}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.28, "sowing_depth": 5.0, "germination_rate": 92, "safety_factor": 1.12}}');


-- ============================================
-- 6. РЕДИС (radish) - 3 сорта
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('french_breakfast', 'Французский завтрак', 'radish', 'Редис', 4.0, 25.0, 22, 2.5, 0.15,
        ARRAY ['spring', 'autumn'], ARRAY ['open_ground', 'greenhouse'],
        '{"root_weight": "15-20g", "root_color": "красный с белым кончиком", "root_shape": "цилиндрический", "use": "свежий"}',
        'Скороспелый сорт для раннего выращивания',
        '{"daily_need_min": 2.0, "daily_need_opt": 3.5, "critical_phases": ["BBCH-51"]}',
        '{"lux_min": 4000, "lux_opt": 15000, "lux_max": 25000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 30}, {"code": "BBCH-19", "gdd_required": 80}, {"code": "BBCH-51", "gdd_required": 150}, {"code": "BBCH-89", "gdd_required": 250}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 1.5, "germination_rate": 85, "safety_factor": 1.2}}'),

       ('red_giant', 'Красный великан', 'radish', 'Редис', 4.0, 25.0, 30, 3.0, 0.2,
        ARRAY ['spring', 'summer', 'autumn'], ARRAY ['open_ground'],
        '{"root_weight": "30-40g", "root_color": "ярко-красный", "root_shape": "удлиненный", "use": "свежий"}',
        'Крупноплодный сорт для длительного хранения',
        '{"daily_need_min": 2.0, "daily_need_opt": 3.5, "critical_phases": ["BBCH-51"]}',
        '{"lux_min": 4000, "lux_opt": 15000, "lux_max": 25000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 35}, {"code": "BBCH-19", "gdd_required": 100}, {"code": "BBCH-51", "gdd_required": 180}, {"code": "BBCH-89", "gdd_required": 300}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.2, "plant_spacing": 0.07, "sowing_depth": 1.5, "germination_rate": 85, "safety_factor": 1.2}}'),

       ('rondar_f1', 'Рондар F1', 'radish', 'Редис', 4.0, 25.0, 20, 3.5, 0.12, ARRAY ['spring', 'summer', 'autumn'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"root_weight": "15-20g", "root_color": "красный", "root_shape": "округлый", "use": "свежий"}',
        'Ультраранний гибрид с дружным созреванием',
        '{"daily_need_min": 2.0, "daily_need_opt": 3.5, "critical_phases": ["BBCH-51"]}',
        '{"lux_min": 4000, "lux_opt": 15000, "lux_max": 25000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 28}, {"code": "BBCH-19", "gdd_required": 75}, {"code": "BBCH-51", "gdd_required": 140}, {"code": "BBCH-89", "gdd_required": 230}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 1.0, "germination_rate": 88, "safety_factor": 1.18}}');


-- ============================================
-- 7. ЛУК (onion) - 2 сорта
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('stuttgarter_riesen', 'Штутгартер Ризен', 'onion', 'Лук', 5.0, 30.0, 100, 5.0, 0.4, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"bulb_weight": "80-120g", "bulb_color": "золотисто-желтый", "bulb_shape": "плоскоокруглый", "sharpness": "среднеострый", "keeping_quality": "отличная"}',
        'Популярный сорт с отличной лежкостью',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 4000, "lux_opt": 18000, "lux_max": 30000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 100}, {"code": "BBCH-19", "gdd_required": 250}, {"code": "BBCH-51", "gdd_required": 450}, {"code": "BBCH-61", "gdd_required": 600}, {"code": "BBCH-71", "gdd_required": 800}, {"code": "BBCH-89", "gdd_required": 1100}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.3, "plant_spacing": 0.1, "sowing_depth": 2.0, "germination_rate": 80, "safety_factor": 1.25}}'),

       ('senshui', 'Сеншуй', 'onion', 'Лук', 5.0, 30.0, 110, 4.5, 0.35, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"bulb_weight": "100-150g", "bulb_color": "бронзовый", "bulb_shape": "округлый", "sharpness": "сладкий"}',
        'Японский салатный сорт',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 4000, "lux_opt": 18000, "lux_max": 30000, "day_length_min": 12, "day_length_opt": 16, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 105}, {"code": "BBCH-19", "gdd_required": 260}, {"code": "BBCH-51", "gdd_required": 470}, {"code": "BBCH-61", "gdd_required": 620}, {"code": "BBCH-71", "gdd_required": 820}, {"code": "BBCH-89", "gdd_required": 1150}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.3, "plant_spacing": 0.12, "sowing_depth": 2.0, "germination_rate": 78, "safety_factor": 1.28}}');


-- ============================================
-- 8. ЧЕСНОК (garlic) - 2 сорта
-- ============================================

INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('lyubasha', 'Любаша', 'garlic', 'Чеснок', 3.0, 28.0, 100, 3.5, 0.5, ARRAY ['autumn'], ARRAY ['open_ground'],
        '{"bulb_weight": "80-120g", "bulb_color": "белый", "cloves_count": "7-9", "sharpness": "острый", "keeping_quality": "отличная"}',
        'Озимый сорт с высокой зимостойкостью',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 4000, "lux_opt": 18000, "lux_max": 30000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 120}, {"code": "BBCH-19", "gdd_required": 280}, {"code": "BBCH-51", "gdd_required": 480}, {"code": "BBCH-61", "gdd_required": 650}, {"code": "BBCH-71", "gdd_required": 850}, {"code": "BBCH-89", "gdd_required": 1100}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.25, "plant_spacing": 0.1, "sowing_depth": 5.0, "germination_rate": 95, "safety_factor": 1.1}}'),

       ('ukrainian_white', 'Украинский белый', 'garlic', 'Чеснок', 3.0, 28.0, 90, 3.0, 0.45, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"bulb_weight": "60-80g", "bulb_color": "белый", "cloves_count": "15-20", "sharpness": "среднеострый"}',
        'Яровой сорт для весенней посадки',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-71"]}',
        '{"lux_min": 4000, "lux_opt": 18000, "lux_max": 30000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-19"]}',
        '[{"code": "BBCH-10", "gdd_required": 100}, {"code": "BBCH-19", "gdd_required": 250}, {"code": "BBCH-51", "gdd_required": 430}, {"code": "BBCH-61", "gdd_required": 600}, {"code": "BBCH-71", "gdd_required": 800}, {"code": "BBCH-89", "gdd_required": 1050}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.25, "plant_spacing": 0.08, "sowing_depth": 4.0, "germination_rate": 93, "safety_factor": 1.12}}');


-- ============================================
-- 9-19. ОСТАЛЬНЫЕ ВИДЫ (по 2 сорта)
-- Ячмень, Баклажан, Горох, Рапс, Соя, Подсолнечник, Пшеница
-- ============================================

-- ЯЧМЕНЬ (barley)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('vakula', 'Вакула', 'barley', 'Ячмень', 3.0, 30.0, 85, 0.55, 0.8, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"use": "пивоваренный", "drought_resistance": "высокая", "protein": "10-12%"}',
        'Пивоваренный ячмень с высоким качеством зерна',
        '{"daily_need_min": 3.0, "daily_need_opt": 4.5, "critical_phases": ["BBCH-30", "BBCH-51", "BBCH-61"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 50000, "day_length_min": 8, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 55}, {"code": "BBCH-19", "gdd_required": 150}, {"code": "BBCH-30", "gdd_required": 300}, {"code": "BBCH-51", "gdd_required": 480}, {"code": "BBCH-61", "gdd_required": 600}, {"code": "BBCH-71", "gdd_required": 750}, {"code": "BBCH-89", "gdd_required": 950}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.03, "sowing_depth": 4.0, "germination_rate": 88, "safety_factor": 1.18}}'),

       ('priazovsky', 'Приазовский', 'barley', 'Ячмень', 3.0, 30.0, 90, 0.6, 0.85, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"use": "фуражный", "drought_resistance": "высокая", "disease_resistance": "мучнистая роса, карликовая ржавчина"}',
        'Засухоустойчивый фуражный ячмень',
        '{"daily_need_min": 2.8, "daily_need_opt": 4.2, "critical_phases": ["BBCH-30", "BBCH-51", "BBCH-61"]}',
        '{"lux_min": 8000, "lux_opt": 30000, "lux_max": 50000, "day_length_min": 8, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 58}, {"code": "BBCH-19", "gdd_required": 155}, {"code": "BBCH-30", "gdd_required": 310}, {"code": "BBCH-51", "gdd_required": 490}, {"code": "BBCH-61", "gdd_required": 620}, {"code": "BBCH-71", "gdd_required": 770}, {"code": "BBCH-89", "gdd_required": 980}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.03, "sowing_depth": 4.5, "germination_rate": 87, "safety_factor": 1.19}}');

-- БАКЛАЖАН (eggplant)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('diamond', 'Алмаз', 'eggplant', 'Баклажан', 13.0, 32.0, 120, 6.0, 0.6, ARRAY ['spring', 'summer'],
        ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "100-150g", "fruit_color": "темно-фиолетовый", "fruit_shape": "цилиндрический", "use": "универсальный"}',
        'Среднеспелый сорт для открытого грунта и теплиц',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 130}, {"code": "BBCH-19", "gdd_required": 380}, {"code": "BBCH-51", "gdd_required": 650}, {"code": "BBCH-61", "gdd_required": 850, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 1000, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1300}, {"code": "BBCH-89", "gdd_required": 1500}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.5, "sowing_depth": 2.0, "germination_rate": 80, "safety_factor": 1.2}, "greenhouse": {"growing_type": "greenhouse", "row_spacing": 0.8, "plant_spacing": 0.6, "sowing_depth": 1.5, "germination_rate": 85, "safety_factor": 1.15}}'),

       ('black_beauty', 'Черный красавец', 'eggplant', 'Баклажан', 13.0, 32.0, 110, 7.0, 0.7,
        ARRAY ['spring', 'summer'], ARRAY ['open_ground', 'greenhouse'],
        '{"fruit_weight": "200-300g", "fruit_color": "черный", "fruit_shape": "грушевидный", "use": "универсальный"}',
        'Крупноплодный сорт с отличными вкусовыми качествами',
        '{"daily_need_min": 2.5, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 125}, {"code": "BBCH-19", "gdd_required": 370}, {"code": "BBCH-51", "gdd_required": 630}, {"code": "BBCH-61", "gdd_required": 820, "is_critical": true}, {"code": "BBCH-71", "gdd_required": 970, "is_critical": true}, {"code": "BBCH-81", "gdd_required": 1250}, {"code": "BBCH-89", "gdd_required": 1450}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.5, "sowing_depth": 2.0, "germination_rate": 82, "safety_factor": 1.18}}');

-- ГОРОХ (pea)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('atlant', 'Атлант', 'pea', 'Горох', 4.0, 28.0, 85, 0.35, 0.8, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"protein": "22-24%", "use": "продовольственный", "lodging_resistance": "высокая", "seed_size": "крупный"}',
        'Лущильный сорт с высоким содержанием белка',
        '{"daily_need_min": 3.0, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 70}, {"code": "BBCH-19", "gdd_required": 180}, {"code": "BBCH-51", "gdd_required": 350}, {"code": "BBCH-61", "gdd_required": 480}, {"code": "BBCH-71", "gdd_required": 650}, {"code": "BBCH-89", "gdd_required": 850}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 5.0, "germination_rate": 88, "safety_factor": 1.18}}'),

       ('madonna', 'Мадонна', 'pea', 'Горох', 4.0, 28.0, 75, 0.4, 0.7, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"protein": "23-25%", "use": "продовольственный", "disease_resistance": "аскохитоз, корневые гнили"}',
        'Раннеспелый сорт для консервирования',
        '{"daily_need_min": 3.0, "daily_need_opt": 4.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 65}, {"code": "BBCH-19", "gdd_required": 170}, {"code": "BBCH-51", "gdd_required": 330}, {"code": "BBCH-61", "gdd_required": 460}, {"code": "BBCH-71", "gdd_required": 620}, {"code": "BBCH-89", "gdd_required": 800}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 5.0, "germination_rate": 90, "safety_factor": 1.15}}');

-- РАПС (rapeseed)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('lider', 'Лидер', 'rapeseed', 'Рапс', 3.0, 30.0, 95, 0.25, 1.2, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"oil_content": "44-46%", "use": "продовольственный", "erucic_acid": "<0.5%", "glucosinolates": "<20"}',
        'Яровой рапс с низким содержанием эруковой кислоты',
        '{"daily_need_min": 3.5, "daily_need_opt": 5.5, "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 80}, {"code": "BBCH-19", "gdd_required": 200}, {"code": "BBCH-51", "gdd_required": 400}, {"code": "BBCH-61", "gdd_required": 550}, {"code": "BBCH-71", "gdd_required": 700}, {"code": "BBCH-89", "gdd_required": 950}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 2.5, "germination_rate": 85, "safety_factor": 1.2}}'),

       ('salsa', 'Сальса', 'rapeseed', 'Рапс', 3.0, 30.0, 100, 0.3, 1.3, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"oil_content": "46-48%", "use": "продовольственный", "erucic_acid": "<0.3%"}', 'Высокоурожайный яровой рапс',
        '{"daily_need_min": 3.5, "daily_need_opt": 5.5, "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '{"lux_min": 5000, "lux_opt": 20000, "lux_max": 35000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 85}, {"code": "BBCH-19", "gdd_required": 210}, {"code": "BBCH-51", "gdd_required": 420}, {"code": "BBCH-61", "gdd_required": 570}, {"code": "BBCH-71", "gdd_required": 720}, {"code": "BBCH-89", "gdd_required": 980}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.05, "sowing_depth": 2.5, "germination_rate": 87, "safety_factor": 1.18}}');

-- СОЯ (soybean)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('vilana', 'Вилана', 'soybean', 'Соя', 10.0, 35.0, 110, 0.3, 0.9, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"protein": "38-40%", "oil_content": "20-22%", "use": "продовольственный"}', 'Раннеспелый сорт сои',
        '{"daily_need_min": 4.0, "daily_need_opt": 6.0, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "short_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 120}, {"code": "BBCH-19", "gdd_required": 280}, {"code": "BBCH-51", "gdd_required": 550}, {"code": "BBCH-61", "gdd_required": 700}, {"code": "BBCH-71", "gdd_required": 900}, {"code": "BBCH-89", "gdd_required": 1200}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.45, "plant_spacing": 0.05, "sowing_depth": 4.0, "germination_rate": 85, "safety_factor": 1.2}}'),

       ('mageva', 'Магева', 'soybean', 'Соя', 10.0, 35.0, 120, 0.35, 1.0, ARRAY ['spring'], ARRAY ['open_ground'],
        '{"protein": "39-41%", "oil_content": "21-23%", "drought_resistance": "высокая"}', 'Засухоустойчивый сорт сои',
        '{"daily_need_min": 3.5, "daily_need_opt": 5.5, "critical_phases": ["BBCH-61", "BBCH-71"]}',
        '{"lux_min": 6000, "lux_opt": 25000, "lux_max": 45000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "short_day", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 130}, {"code": "BBCH-19", "gdd_required": 300}, {"code": "BBCH-51", "gdd_required": 580}, {"code": "BBCH-61", "gdd_required": 740}, {"code": "BBCH-71", "gdd_required": 950}, {"code": "BBCH-89", "gdd_required": 1280}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.45, "plant_spacing": 0.05, "sowing_depth": 4.0, "germination_rate": 86, "safety_factor": 1.19}}');

-- ПОДСОЛНЕЧНИК (sunflower)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('lakomka', 'Лакомка', 'sunflower', 'Подсолнечник', 8.0, 35.0, 105, 0.25, 1.6, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"oil_content": "48-50%", "use": "кондитерский", "drought_resistance": "высокая", "seed_size": "крупный"}',
        'Кондитерский сорт с крупными семенами',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '{"lux_min": 10000, "lux_opt": 40000, "lux_max": 65000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 100}, {"code": "BBCH-19", "gdd_required": 250}, {"code": "BBCH-51", "gdd_required": 500}, {"code": "BBCH-61", "gdd_required": 700}, {"code": "BBCH-71", "gdd_required": 900}, {"code": "BBCH-89", "gdd_required": 1150}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.25, "sowing_depth": 6.0, "germination_rate": 90, "safety_factor": 1.15}}'),

       ('oliver', 'Оливер', 'sunflower', 'Подсолнечник', 8.0, 35.0, 110, 0.3, 1.7, ARRAY ['spring'],
        ARRAY ['open_ground'], '{"oil_content": "52-54%", "use": "масличный", "drought_resistance": "высокая"}',
        'Высокомасличный гибрид для производства масла',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '{"lux_min": 10000, "lux_opt": 40000, "lux_max": 65000, "day_length_min": 10, "day_length_opt": 14, "photoperiod_type": "day_neutral", "critical_phases": ["BBCH-51", "BBCH-61"]}',
        '[{"code": "BBCH-10", "gdd_required": 105}, {"code": "BBCH-19", "gdd_required": 260}, {"code": "BBCH-51", "gdd_required": 520}, {"code": "BBCH-61", "gdd_required": 720}, {"code": "BBCH-71", "gdd_required": 930}, {"code": "BBCH-89", "gdd_required": 1200}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.7, "plant_spacing": 0.25, "sowing_depth": 6.0, "germination_rate": 92, "safety_factor": 1.12}}');

-- ПШЕНИЦА (wheat)
INSERT INTO public.growing_varieties (id, name, species_key, species_name, base_temperature, max_temperature, days_to_maturity,
                       yield_potential, plant_height, recommended_seasons, growing_types, characteristics, description,
                       water_requirement, light_requirement, phenophase_gdd, seeding_rates)
VALUES ('bezostaya_1', 'Безостая 1', 'wheat', 'Пшеница', 3.0, 28.0, 280, 0.6, 1.0, ARRAY ['autumn'],
        ARRAY ['open_ground'],
        '{"grain_type": "мягкая", "use": "хлебопекарная", "winter_hardiness": "средняя", "protein": "12-14%"}',
        'Классический сорт озимой пшеницы',
        '{"daily_need_min": 3.0, "daily_need_opt": 5.0, "critical_phases": ["BBCH-30", "BBCH-51", "BBCH-61"]}',
        '{"lux_min": 8000, "lux_opt": 35000, "lux_max": 55000, "day_length_min": 8, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 50}, {"code": "BBCH-19", "gdd_required": 150}, {"code": "BBCH-30", "gdd_required": 300}, {"code": "BBCH-51", "gdd_required": 500}, {"code": "BBCH-61", "gdd_required": 650}, {"code": "BBCH-71", "gdd_required": 800}, {"code": "BBCH-89", "gdd_required": 1000}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.03, "sowing_depth": 5.0, "germination_rate": 90, "safety_factor": 1.15}}'),

       ('saratovskaya_29', 'Саратовская 29', 'wheat', 'Пшеница', 3.0, 30.0, 100, 0.5, 0.9, ARRAY ['spring'],
        ARRAY ['open_ground'],
        '{"grain_type": "твердая", "use": "макаронная", "drought_resistance": "высокая", "protein": "14-16%"}',
        'Яровая твердая пшеница для макарон',
        '{"daily_need_min": 2.8, "daily_need_opt": 4.5, "critical_phases": ["BBCH-30", "BBCH-51", "BBCH-61"]}',
        '{"lux_min": 8000, "lux_opt": 35000, "lux_max": 55000, "day_length_min": 8, "day_length_opt": 14, "photoperiod_type": "long_day", "critical_phases": ["BBCH-30", "BBCH-51"]}',
        '[{"code": "BBCH-10", "gdd_required": 60}, {"code": "BBCH-19", "gdd_required": 160}, {"code": "BBCH-30", "gdd_required": 320}, {"code": "BBCH-51", "gdd_required": 520}, {"code": "BBCH-61", "gdd_required": 670}, {"code": "BBCH-71", "gdd_required": 820}, {"code": "BBCH-89", "gdd_required": 1050}]',
        '{"open_ground": {"growing_type": "open_ground", "row_spacing": 0.15, "plant_spacing": 0.03, "sowing_depth": 5.0, "germination_rate": 92, "safety_factor": 1.12}}');


-- ============================================
-- ПРОВЕРКА ИМПОРТА
-- ============================================

DO
$$
    DECLARE
        tomato_count    INTEGER;
        cucumber_count  INTEGER;
        cabbage_count   INTEGER;
        potato_count    INTEGER;
        corn_count      INTEGER;
        radish_count    INTEGER;
        onion_count     INTEGER;
        garlic_count    INTEGER;
        barley_count    INTEGER;
        eggplant_count  INTEGER;
        pea_count       INTEGER;
        rapeseed_count  INTEGER;
        soybean_count   INTEGER;
        sunflower_count INTEGER;
        wheat_count     INTEGER;
        total_count     INTEGER;
    BEGIN
        SELECT COUNT(*) INTO tomato_count FROM public.growing_varieties WHERE species_key = 'tomato';
        SELECT COUNT(*) INTO cucumber_count FROM public.growing_varieties WHERE species_key = 'cucumber';
        SELECT COUNT(*) INTO cabbage_count FROM public.growing_varieties WHERE species_key = 'cabbage';
        SELECT COUNT(*) INTO potato_count FROM public.growing_varieties WHERE species_key = 'potato';
        SELECT COUNT(*) INTO corn_count FROM public.growing_varieties WHERE species_key = 'corn';
        SELECT COUNT(*) INTO radish_count FROM public.growing_varieties WHERE species_key = 'radish';
        SELECT COUNT(*) INTO onion_count FROM public.growing_varieties WHERE species_key = 'onion';
        SELECT COUNT(*) INTO garlic_count FROM public.growing_varieties WHERE species_key = 'garlic';
        SELECT COUNT(*) INTO barley_count FROM public.growing_varieties WHERE species_key = 'barley';
        SELECT COUNT(*) INTO eggplant_count FROM public.growing_varieties WHERE species_key = 'eggplant';
        SELECT COUNT(*) INTO pea_count FROM public.growing_varieties WHERE species_key = 'pea';
        SELECT COUNT(*) INTO rapeseed_count FROM public.growing_varieties WHERE species_key = 'rapeseed';
        SELECT COUNT(*) INTO soybean_count FROM public.growing_varieties WHERE species_key = 'soybean';
        SELECT COUNT(*) INTO sunflower_count FROM public.growing_varieties WHERE species_key = 'sunflower';
        SELECT COUNT(*) INTO wheat_count FROM public.growing_varieties WHERE species_key = 'wheat';

        SELECT COUNT(*) INTO total_count FROM public.growing_varieties;

        RAISE NOTICE '=========================================';
        RAISE NOTICE 'ИТОГИ ИМПОРТА СОРТОВ';
        RAISE NOTICE '=========================================';
        RAISE NOTICE 'Томаты (tomato): % сортов', tomato_count;
        RAISE NOTICE 'Огурцы (cucumber): % сортов', cucumber_count;
        RAISE NOTICE 'Капуста (cabbage): % сортов', cabbage_count;
        RAISE NOTICE 'Картофель (potato): % сортов', potato_count;
        RAISE NOTICE 'Кукуруза (corn): % сортов', corn_count;
        RAISE NOTICE 'Редис (radish): % сортов', radish_count;
        RAISE NOTICE 'Лук (onion): % сортов', onion_count;
        RAISE NOTICE 'Чеснок (garlic): % сортов', garlic_count;
        RAISE NOTICE 'Ячмень (barley): % сортов', barley_count;
        RAISE NOTICE 'Баклажан (eggplant): % сортов', eggplant_count;
        RAISE NOTICE 'Горох (pea): % сортов', pea_count;
        RAISE NOTICE 'Рапс (rapeseed): % сортов', rapeseed_count;
        RAISE NOTICE 'Соя (soybean): % сортов', soybean_count;
        RAISE NOTICE 'Подсолнечник (sunflower): % сортов', sunflower_count;
        RAISE NOTICE 'Пшеница (wheat): % сортов', wheat_count;
        RAISE NOTICE '-----------------------------------------';
        RAISE NOTICE 'ВСЕГО СОРТОВ: %', total_count;
        RAISE NOTICE '=========================================';
    END
$$;