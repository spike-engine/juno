package messages_test

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	ibctypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

func (suite *ModuleSuite) TestParseMsgValue() {
	msg := ibctypes.MsgTransfer{
		SourcePort: "port",
	}

	anyMsg, err := codectypes.NewAnyWithValue(&msg)
	suite.Require().NoError(err)

	bz, err := suite.module.ParseMsgValue(anyMsg)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(bz)
}
