-- Членство пользователей в организациях
CREATE TABLE auth_memberships
(
    id              VARCHAR(36) PRIMARY KEY,
    user_id         UUID        NOT NULL REFERENCES auth_users (id) ON DELETE CASCADE,
    organization_id UUID        NOT NULL REFERENCES auth_organizations (id) ON DELETE CASCADE,
    role            VARCHAR(50) NOT NULL,
    is_active       BOOLEAN   DEFAULT true,
    joined_at       TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    UNIQUE (user_id, organization_id)
);
