CREATE DATABASE panda IF NOT EXISTS;

USE panda;

CREATE TABLE users (
    id BIGINT,
    username TEXT,
    password_hash TEXT,
    password_salt TEXT,
    email TEXT,
    name TEXT,
    bio TEXT,
    profile_photo_b64 MEDIUMTEXT,
    primary_color int,
    secondary_color int,
    is_bitcoin_baller BOOLEAN,
    links_json TEXT,
    created_at TEXT
);

UPDATE users SET username = ?, SET email = ?, SET name = ?, SET bio = ?, SET profile_photo_b64 = ?, SET primary_color = ?, SET secondary_color = ?, SET links_json = ? WHERE id = ?
