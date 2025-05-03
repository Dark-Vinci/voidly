-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages
    ADD COLUMN deleted_at TIMESTAMP NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE messages
DROP COLUMN deleted_at;
-- +goose StatementEnd
