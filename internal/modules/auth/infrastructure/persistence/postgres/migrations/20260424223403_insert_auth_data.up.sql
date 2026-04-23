INSERT INTO auth_users (id, email, username, password, first_name, last_name, role, status)
VALUES ('00000000-0000-0000-0000-000000000001',
        'admin@example.com',
        'admin',
        '$2a$10$N9qo8uLOickgx2ZMRZoMy.MrJkq6dVg8qKjG7QwYqZ5YtYtYtYtY', -- admin123
        'Admin',
        'User',
        'super_admin',
        'active')
ON CONFLICT (id) DO NOTHING;

-- ============================================
-- Создание тестовой организации
-- ============================================

INSERT INTO auth_organizations (id, name, tax_id, address, phone, email, is_active)
VALUES ('11111111-1111-1111-1111-111111111111',
        'Тестовая ферма',
        '1234567890',
        'г. Москва, ул. Тестовая, д. 1',
        '+7 (999) 123-45-67',
        'info@testfarm.ru',
        true)
ON CONFLICT (id) DO NOTHING;

-- ============================================
-- Создание членства для супер-админа
-- ============================================

INSERT INTO auth_memberships (id, user_id, organization_id, role, is_active)
VALUES ('22222222-2222-2222-2222-222222222222',
        '00000000-0000-0000-0000-000000000001',
        '11111111-1111-1111-1111-111111111111',
        'owner',
        true)
ON CONFLICT (user_id, organization_id) DO NOTHING;