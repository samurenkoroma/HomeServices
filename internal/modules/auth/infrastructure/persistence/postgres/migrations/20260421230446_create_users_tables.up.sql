-- ============================================
-- Таблица пользователей (users)
-- ============================================

CREATE TABLE IF NOT EXISTS auth_users
(
    id                      UUID PRIMARY KEY,
    email                   VARCHAR(255)             NOT NULL UNIQUE,
    username                VARCHAR(100)             NOT NULL UNIQUE,
    password                VARCHAR(255)             NOT NULL,
    first_name              VARCHAR(100),
    last_name               VARCHAR(100),
    phone                   VARCHAR(50),

    current_organization_id UUID REFERENCES auth_organizations (id),

    role                    VARCHAR(50)              NOT NULL DEFAULT 'client',
    status                  VARCHAR(50)              NOT NULL DEFAULT 'active',

    last_login              TIMESTAMP WITH TIME ZONE,
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_user_status CHECK (status IN ('active', 'inactive', 'blocked'))
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_users_email ON auth_users (email);
CREATE INDEX IF NOT EXISTS idx_users_username ON auth_users (username);
CREATE INDEX IF NOT EXISTS idx_users_role ON auth_users (role);
CREATE INDEX IF NOT EXISTS idx_users_status ON auth_users (status);

