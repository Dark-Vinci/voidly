package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUser, downUser)
}

func upUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table users (
			id uuid
				constraint users_pk
					primary key default uuid_generate_v4(),,
			username TEXT NOT NULL,
			email varchar unique not null,
			password text not null,
			created_at timestamp default current_timestamp not null,
			updated_at timestamp default current_timestamp not null
		);
	`)

	return err
}

func downUser(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`drop table users`)

	return err
}
