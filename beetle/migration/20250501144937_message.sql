-- +goose Up
-- +goose StatementBegin
create table messages (
    id uuid
        constraint messages_pk
            primary key default uuid_generate_v4(),
    content text not null,
    chat_id uuid not null,
    from_user_id uuid not null,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
-- +goose StatementEnd
