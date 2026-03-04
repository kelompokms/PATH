-- +goose Up
CREATE TABLE user(
    id INTEGER AUTO_INCREMENT PRIMARY KEY
);

-- +goose Down
DROP TABLE user;
