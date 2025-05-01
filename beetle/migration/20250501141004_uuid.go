package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUuid, downUuid)
}

func upUuid(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`)

	return err
}

func downUuid(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
