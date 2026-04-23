-- Организации
CREATE TABLE auth_organizations
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    tax_id     VARCHAR(50),
    address    TEXT,
    phone      VARCHAR(50),
    email      VARCHAR(255),
    is_active  BOOLEAN   DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
