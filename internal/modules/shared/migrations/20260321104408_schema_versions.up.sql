CREATE TABLE schema_migrations
(
    id         SERIAL PRIMARY KEY,
    module     TEXT                     NOT NULL,
    version    BIGINT                   NOT NULL,
    applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    dirty      BOOLEAN                  NOT NULL DEFAULT FALSE,
    UNIQUE (module, version)
);