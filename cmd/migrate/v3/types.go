package v3

import (
	databaseconfig "github.com/spike-engine/juno/database/config"
	loggingconfig "github.com/spike-engine/juno/logging/config"
	"github.com/spike-engine/juno/modules/pruning"
	"github.com/spike-engine/juno/modules/telemetry"
	nodeconfig "github.com/spike-engine/juno/node/config"
	parserconfig "github.com/spike-engine/juno/parser/config"
	pricefeedconfig "github.com/spike-engine/juno/pricefeed"
	"github.com/spike-engine/juno/types/config"
)

// Config defines all necessary juno configuration parameters.
type Config struct {
	Chain    config.ChainConfig    `yaml:"chain"`
	Node     nodeconfig.Config     `yaml:"node"`
	Parser   parserconfig.Config   `yaml:"parsing"`
	Database databaseconfig.Config `yaml:"database"`
	Logging  loggingconfig.Config  `yaml:"logging"`

	// The following are there to support modules which config are present if they are enabled

	Telemetry *telemetry.Config       `yaml:"telemetry,omitempty"`
	Pruning   *pruning.Config         `yaml:"pruning,omitempty"`
	PriceFeed *pricefeedconfig.Config `yaml:"pricefeed,omitempty"`
}
