-- +goose Up
CREATE TABLE mapel (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(32),
    password VARCHAR(255)
);

-- +goose Down
DROP TABLE mapel;
DROP TABLE user;
