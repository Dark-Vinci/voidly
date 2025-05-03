-- +goose Up
-- +goose StatementBegin
ALTER TABLE chats
    ADD COLUMN deleted_at TIMESTAMP NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE chats
DROP COLUMN deleted_at;
-- +goose StatementEnd
