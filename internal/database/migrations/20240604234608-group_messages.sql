-- +migrate Up
CREATE TABLE group_messages
(
    id         uuid      DEFAULT Uuid_generate_v4() primary key,
    group_id   uuid not null,
    user_id    uuid not null,
    message    text not null,
    created_at timestamp DEFAULT now(),

    CONSTRAINT fk_group_id FOREIGN KEY (group_id) REFERENCES groups (id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +migrate Down
DROP TABLE IF EXISTS group_messages;