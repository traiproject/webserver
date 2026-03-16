-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, email) 
VALUES 
    (uuid_generate_v4(), 'Alice Admin', 'alice@example.com'),
    (uuid_generate_v4(), 'Bob Builder', 'bob@example.com'),
    (uuid_generate_v4(), 'Cloe Coach', 'cloe@example.com'),
    (uuid_generate_v4(), 'Denis Dentist', 'denis@example.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE email IN ('alice@example.com', 'bob@example.com', 'cloe@example.com', 'denis@example.com');
-- +goose StatementEnd
