package adapter

import "github.com/grafana/grafana/pkg/modules"

const (
	// BackgroundServices is the module name for all background services.
	// Individual background services are registered as dependencies of this module.
	BackgroundServices = "background-services"

	// Core is the core module that includes critical services like the API server.
	Core = "core"

	// All is the top-level module that depends on both core and background services.
	// This module represents the complete system startup.
	All = "all"
)

// dependencyMap returns the module dependency relationships for the background service system.
// It defines the startup order and dependencies between different module groups:
//   - GrafanaAPIServer has no dependencies (starts first)
//   - Core depends on GrafanaAPIServer
//   - BackgroundServices has no explicit dependencies (populated dynamically)
//   - All depends on both Core and BackgroundServices (starts last)
//
// Background services are automatically added as dependencies to the BackgroundServices module
// unless they are explicitly listed in this map with custom dependencies.
func dependencyMap() map[string][]string {
	return map[string][]string{
		modules.GrafanaAPIServer: {},
		Core:                     {modules.GrafanaAPIServer},
		BackgroundServices:       {},
		All:                      {Core, BackgroundServices},
	}
}
