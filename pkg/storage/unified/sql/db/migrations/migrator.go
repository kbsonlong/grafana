package migrations

import (
	"context"

	"xorm.io/xorm"

	"github.com/grafana/grafana/pkg/services/sqlstore/migrator"
	"github.com/grafana/grafana/pkg/setting"
)

func MigrateResourceStore(ctx context.Context, engine *xorm.Engine, cfg *setting.Cfg) error {
	// Check skip_migrations parameter
	sec := cfg.Raw.Section("database")
	if sec.Key("skip_migrations").MustBool(false) {
		return nil
	}

	mg := migrator.NewScopedMigrator(engine, cfg, "resource")
	mg.AddCreateMigration()

	initResourceTables(mg)

	return mg.RunMigrations(
		ctx,
		sec.Key("migration_locking").MustBool(true),
		sec.Key("locking_attempt_timeout_sec").MustInt(),
	)
}
