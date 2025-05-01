-- +goose Up
-- +goose StatementBegin
create table chats (
   id uuid
       constraint chats_pk
           primary key default uuid_generate_v4(),
   user_a uuid not null,
   user_b uuid not null,
   created_at timestamp default current_timestamp not null,
   updated_at timestamp default current_timestamp not null,
   unique (user_a, user_b)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
