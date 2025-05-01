-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table users (
   id uuid
       constraint users_pk
           primary key default uuid_generate_v4(),
   username TEXT NOT NULL,
   email varchar unique not null,
   password text not null,
   created_at timestamp default current_timestamp not null,
   updated_at timestamp default current_timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
