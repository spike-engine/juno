package messages_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/forbole/juno/v3/modules/messages"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestModuleSuite(t *testing.T) {
	suite.Run(t, new(ModuleSuite))
}

type ModuleSuite struct {
	suite.Suite

	cdc    codec.Codec
	module *messages.Module
}

func (suite *ModuleSuite) SetupSuite() {
	encodingConfig := simapp.MakeTestEncodingConfig()
	suite.module = messages.NewModule(encodingConfig.Marshaler, nil)
}
