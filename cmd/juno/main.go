package main

import (
	"os"

	"github.com/spike-engine/juno/cmd/parse/types"

	"github.com/spike-engine/juno/modules/messages"
	"github.com/spike-engine/juno/modules/registrar"

	"github.com/spike-engine/juno/cmd"
)

func main() {
	// JunoConfig the runner
	config := cmd.NewConfig("juno").
		WithParseConfig(types.NewConfig().
			WithRegistrar(registrar.NewDefaultRegistrar(
				messages.CosmosMessageAddressesParser,
			)),
		)

	// Run the commands and panic on any error
	exec := cmd.BuildDefaultExecutor(config)
	err := exec.Execute()
	if err != nil {
		os.Exit(1)
	}
}
