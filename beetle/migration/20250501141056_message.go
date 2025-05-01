package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upMessage, downMessage)
}

func upMessage(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table messages (
		    create table messages (
		        id uuid
					constraint messages_pk
						primary key default uuid_generate_v4(),,
				content text not null,
				chat_id uuid not null,
				from_user_id uuid not null,
				created_at timestamp default current_timestamp not null,
				updated_at timestamp default current_timestamp not null
		    )
		)`)

	return err
}

func downMessage(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table messages`)

	return err
}
