-- +goose Up
-- +goose StatementBegin
INSERT INTO showcase_item (id, title) 
VALUES 
    (uuid_generate_v4(), 'item 1'),
    (uuid_generate_v4(), 'item 2'),
    (uuid_generate_v4(), 'item 3'),
    (uuid_generate_v4(), 'item 4');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM showcase_item WHERE title IN ('item 1', 'item 2', 'item 3', 'item 4');
-- +goose StatementEnd
