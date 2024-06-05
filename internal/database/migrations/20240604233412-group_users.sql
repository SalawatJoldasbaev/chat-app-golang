-- +migrate Up
CREATE TABLE group_users
(
    group_id   UUID NOT NULL,
    user_id    UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +migrate Down
DROP TABLE IF EXISTS group_users;