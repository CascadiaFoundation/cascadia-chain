package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/evmos/evmos/v9/x/feedist/client/cli"
	"github.com/evmos/evmos/v9/x/feedist/client/rest"
)

var (
	RegisterFeedistProposalHandler = govclient.NewProposalHandler(cli.CmdRegisterFeedistProposal, rest.RegisterFeedistProposalRESTHandler)
	CancelFeedistProposalHandler   = govclient.NewProposalHandler(cli.CmdCancelFeedistProposal, rest.CancelFeedistProposalRequestRESTHandler)
)
