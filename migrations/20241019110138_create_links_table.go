package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateLinksTable, downCreateLinksTable)
}

func upCreateLinksTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE link (
			id SERIAL PRIMARY KEY,
			title TEXT,
			full_link TEXT,
			shorten_link_code VARCHAR(8),
			follows BIGINT DEFAULT 0,
			shorten_link text
		);
	`)
	return err
}

func downCreateLinksTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE link;")
	return err
}
