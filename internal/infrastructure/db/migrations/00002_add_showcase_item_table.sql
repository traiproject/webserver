-- +goose Up
-- +goose StatementBegin
CREATE TABLE showcase_item (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE showcase_item;
-- +goose StatementEnd
