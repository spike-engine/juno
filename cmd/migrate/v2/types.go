package v2

import (
	"time"

	loggingconfig "github.com/spike-engine/juno/logging/config"
	"github.com/spike-engine/juno/modules/pruning"
	"github.com/spike-engine/juno/modules/telemetry"
	nodeconfig "github.com/spike-engine/juno/node/config"
	pricefeedconfig "github.com/spike-engine/juno/pricefeed"
	"github.com/spike-engine/juno/types/config"
)

type Config struct {
	Chain    config.ChainConfig   `yaml:"chain"`
	Node     nodeconfig.Config    `yaml:"node"`
	Parser   ParserConfig         `yaml:"parsing"`
	Database DatabaseConfig       `yaml:"database"`
	Logging  loggingconfig.Config `yaml:"logging"`

	// The following are there to support modules which config are present if they are enabled

	Telemetry *telemetry.Config       `yaml:"telemetry,omitempty"`
	Pruning   *pruning.Config         `yaml:"pruning,omitempty"`
	PriceFeed *pricefeedconfig.Config `yaml:"pricefeed,omitempty"`
}

type ParserConfig struct {
	Workers         int64  `yaml:"workers"`
	ParseNewBlocks  bool   `yaml:"listen_new_blocks"`
	ParseOldBlocks  bool   `yaml:"parse_old_blocks"`
	GenesisFilePath string `yaml:"genesis_file_path,omitempty"`
	ParseGenesis    bool   `yaml:"parse_genesis"`
	StartHeight     int64  `yaml:"start_height"`
	FastSync        bool   `yaml:"fast_sync,omitempty"`

	// Following there are the new fields that have been added into v3. We use pointers and the "omitempty" clause
	// to make sure that if they are not already specified, then we get nil as values

	AvgBlockTime *time.Duration `yaml:"average_block_time,omitempty"`
}

type DatabaseConfig struct {
	Name               string `yaml:"name"`
	Host               string `yaml:"host"`
	Port               int64  `yaml:"port"`
	User               string `yaml:"user"`
	Password           string `yaml:"password"`
	SSLMode            string `yaml:"ssl_mode,omitempty"`
	Schema             string `yaml:"schema,omitempty"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`

	// Following there are the new fields that have been added into v3. We use pointers and the "omitempty" clause
	// to make sure that if they are not already specified, then we get nil as values

	PartitionSize      *int64 `yaml:"partition_size,omitempty"`
	PartitionBatchSize *int64 `yaml:"partition_batch,omitempty"`
}
