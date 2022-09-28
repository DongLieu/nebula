package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/nebula-labs/nebula/x/launchpad/client/cli"
	"github.com/nebula-labs/nebula/x/launchpad/client/rest"
)

var (
	SetProjectVerifiedProposalHandler = govclient.NewProposalHandler(cli.CmdSubmitSetProjectVerifiedProposal, rest.ProposalSubmitSetProjectVerifiedRESTHandler)
)
