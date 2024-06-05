-- +migrate Up
CREATE TABLE IF NOT EXISTS sessions
(
    id         UUID DEFAULT Uuid_generate_v4() PRIMARY KEY,
    user_id    UUID      NOT NULL,
    expires    TIMESTAMP NOT NULL,
    data       TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_sessions_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT uc_sessions_id UNIQUE (id)
);
-- +migrate Down
DROP TABLE IF EXISTS sessions;