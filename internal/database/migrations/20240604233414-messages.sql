-- +migrate Up
CREATE TABLE messages
(
    id          uuid      DEFAULT Uuid_generate_v4() primary key,
    sender_id   uuid not null,
    receiver_id uuid not null,
    message     text not null,
    created_at  timestamp DEFAULT now(),

    CONSTRAINT fk_sender_id FOREIGN KEY (sender_id) REFERENCES users (id),
    CONSTRAINT fk_receiver_id FOREIGN KEY (receiver_id) REFERENCES users (id)
);
-- +migrate Down
DROP TABLE IF EXISTS messages;