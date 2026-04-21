-- +goose Up
INSERT INTO pengguna(nama, email, telepon, password) VALUES
('Budiman', 'budiman@gmail.com', '0851234123', '$2a$10$.4ZkjAGsq7vrrsHLtlgXZObXIbuAUBpKbrXMi8pkEm/PP6/MmuRcS'),
('Clara', 'clair@gmail.com', '0882123123', '$2a$10$B4mcLhU6jBpr9ZZSCt2Vk.hb6mjx2QLp2cqya7YmJks2l4S1XOWNK'),
('Clara', 'admin@admin', '01234567890', '$2a$10$..Tm/xSwohhkqacM71rgPOdP4ttE2kOdEHU.yQq0VJ2oqiBO1KIw.');

-- +goose Down
TRUNCATE pengguna;
