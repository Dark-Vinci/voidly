package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upChat, downChat)
}

func upChat(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
	 create table chats (
		id uuid
			constraint users_pk
				primary key default uuid_generate_v4(),,
		user_a uuid not null,
		user_b uuid not null,
		created_at timestamp default current_timestamp not null,
		updated_at timestamp default current_timestamp not null,
		unique (user_a, user_b)
	)`)

	return err
}

func downChat(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table chats`)

	return err
}
