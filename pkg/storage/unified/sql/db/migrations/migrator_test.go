package migrations

import (
	"context"
	"testing"

	"github.com/grafana/grafana/pkg/setting"
	"github.com/stretchr/testify/assert"
)

func TestMigrateResourceStore_SkipMigrations(t *testing.T) {
	// Test case: skip_migrations = true should return nil without running migrations
	cfg := setting.NewCfg()
	cfg.Raw.NewSection("database")
	cfg.Raw.Section("database").NewKey("skip_migrations", "true")

	err := MigrateResourceStore(context.Background(), nil, cfg)
	assert.NoError(t, err, "MigrateResourceStore should return nil when skip_migrations is true")
}