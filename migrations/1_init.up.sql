CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE,
    message_status TEXT DEFAULT 'waiting_payload',
    payload TEXT,
    operations INT NOT NULL DEFAULT 5,
    file_name TEXT DEFAULT ''
);