-- +migrate Up
CREATE TABLE groups
(
    id         UUID      DEFAULT Uuid_generate_v4() PRIMARY KEY,
    owner_id   UUID         NOT NULL,
    name       VARCHAR(255) NOT NULL,
    max_users  INT       DEFAULT 6,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_groups_owner_id FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +migrate Down
DROP TABLE IF EXISTS groups;