package migrations

import (
	"context"
	"example/pkg/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// A collection of migrations.
var Migrations = migrate.NewMigrations()

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// Create FormModel table
		_, err := db.NewCreateTable().
			Model((*model.FormModel)(nil)).
			IfNotExists(). // Avoid recreating the table if it already exists
			Exec(ctx)
		if err != nil {
			return err
		}

		// Create Todo table
		_, err = db.NewCreateTable().
			Model((*model.Todo)(nil)).
			IfNotExists(). // Avoid recreating the table if it already exists
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*model.Todo)(nil)).IfExists().Exec(ctx)
		if err != nil {
			return err
		}

		_, err = db.NewDropTable().Model((*model.FormModel)(nil)).IfExists().Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})

}
