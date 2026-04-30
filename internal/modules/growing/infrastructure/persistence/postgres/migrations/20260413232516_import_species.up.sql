-- ============================================
-- Импорт видов культур (species)
-- ============================================

-- Очищаем таблицу перед импортом (опционально)
-- TRUNCATE species CASCADE;

-- Вставка данных
INSERT INTO public.growing_crops (key, name, family, category, image_url, description)
VALUES ('tomato', 'Томат', 'nightshade', 'Овощные', 'https://images.example.com/crops/tomato.jpg',
        'Овощная культура семейства пасленовых. Плоды — сочные ягоды, используемые в кулинарии.'),
       ('cucumber', 'Огурец', 'cucurbits', 'Овощные', 'https://images.example.com/crops/cucumber.jpg',
        'Овощная культура семейства тыквенных. Плоды используются в свежем виде, для засолки и маринования.'),
       ('cabbage', 'Капуста', 'brassicaceae', 'Овощные', 'https://images.example.com/crops/cabbage.jpg',
        'Овощная культура семейства капустных. Кочанная капуста богата витаминами и клетчаткой.'),
       ('potato', 'Картофель', 'nightshade', 'Овощные', 'https://images.example.com/crops/potato.jpg',
        'Клубнеплодная культура семейства пасленовых. Один из основных продуктов питания.'),
       ('corn', 'Кукуруза', 'poaceae', 'Зерновые', 'https://images.example.com/crops/corn.jpg',
        'Зерновая культура семейства злаковых. Используется в пищу, на корм скоту и для переработки.'),
       ('radish', 'Редис', 'brassicaceae', 'Овощные', 'https://images.example.com/crops/radish.jpg',
        'Корнеплодная культура семейства капустных. Скороспелый овощ с островатым вкусом.'),
       ('onion', 'Лук', 'amaryllidaceae', 'Овощные', 'https://images.example.com/crops/onion.jpg',
        'Луковичная культура семейства амариллисовых. Широко используется в кулинарии.'),
       ('garlic', 'Чеснок', 'amaryllidaceae', 'Овощные', 'https://images.example.com/crops/garlic.jpg',
        'Луковичная культура семейства амариллисовых. Ценная пряность и лекарственное растение.'),
       ('dill', 'Укроп', 'apiaceae', 'Зеленные', 'https://images.example.com/crops/dill.jpg',
        'Пряная зелень семейства зонтичных. Используется в свежем и сушеном виде.'),
       ('parsley', 'Петрушка', 'apiaceae', 'Зеленные', 'https://images.example.com/crops/parsley.jpg',
        'Двулетнее растение семейства зонтичных. Листовая и корневая петрушка.'),
       ('lettuce', 'Салат', 'asteraceae', 'Зеленные', 'https://images.example.com/crops/lettuce.jpg',
        'Листовая овощная культура семейства астровых. Богат витаминами.'),
       ('basil', 'Базилик', 'lamiaceae', 'Зеленные', 'https://images.example.com/crops/basil.jpg',
        'Пряная культура семейства яснотковых. Ценный источник эфирных масел.'),
       ('barley', 'Ячмень', 'poaceae', 'Зерновые', 'https://images.example.com/crops/barley.jpg',
        'Зерновая культура семейства злаковых. Используется для производства крупы, пива и кормов.'),
       ('eggplant', 'Баклажан', 'nightshade', 'Овощные', 'https://images.example.com/crops/eggplant.jpg',
        'Овощная культура семейства пасленовых. Плоды используются в кулинарии.'),
       ('pea', 'Горох', 'fabaceae', 'Бобовые', 'https://images.example.com/crops/pea.jpg',
        'Бобовая культура семейства бобовых. Ценный источник белка.'),
       ('rapeseed', 'Рапс', 'brassicaceae', 'Масличные', 'https://images.example.com/crops/rapeseed.jpg',
        'Масличная культура семейства капустных. Используется для производства масла и биодизеля.'),
       ('soybean', 'Соя', 'fabaceae', 'Бобовые', 'https://images.example.com/crops/soybean.jpg',
        'Бобовая культура семейства бобовых. Ценный источник растительного белка и масла.'),
       ('sunflower', 'Подсолнечник', 'asteraceae', 'Масличные', 'https://images.example.com/crops/sunflower.jpg',
        'Масличная культура семейства астровых. Семена используются для производства масла.'),
       ('wheat', 'Пшеница', 'poaceae', 'Зерновые', 'https://images.example.com/crops/wheat.jpg',
        'Зерновая культура семейства злаковых. Основной хлебный злак.')
ON CONFLICT (key) DO UPDATE SET name        = EXCLUDED.name,
                                family      = EXCLUDED.family,
                                category    = EXCLUDED.category,
                                image_url   = EXCLUDED.image_url,
                                description = EXCLUDED.description,
                                updated_at  = NOW();

-- Проверка импорта
DO
$$
    DECLARE
        species_count INTEGER;
    BEGIN
        SELECT COUNT(*) INTO species_count FROM public.growing_crops;
        RAISE NOTICE 'Импортировано видов: %', species_count;
    END
$$;