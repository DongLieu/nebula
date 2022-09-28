package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeSetProjectVerifiedProposal = "SetProjectVerifiedProposal"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeSetProjectVerifiedProposal)
	govtypes.RegisterProposalTypeCodec(&SetProjectVerifiedProposal{}, "nebula/SetProjectVerifiedProposal")
}

var (
	_ govtypes.Content = &SetProjectVerifiedProposal{}
)

func NewSetProjectVerifiedProposal(title, description, projectOwner string, projectId uint64) govtypes.Content {
	return &SetProjectVerifiedProposal{
		Title:        title,
		Description:  description,
		ProjectOwner: projectOwner,
		ProjectId:    projectId,
	}
}

func (p *SetProjectVerifiedProposal) GetTitle() string { return p.Title }

func (p *SetProjectVerifiedProposal) GetDescription() string { return p.Description }

func (p *SetProjectVerifiedProposal) ProposalRoute() string { return RouterKey }

func (p *SetProjectVerifiedProposal) ProposalType() string {
	return ProposalTypeSetProjectVerifiedProposal
}

func (p SetProjectVerifiedProposal) String() string {
	return fmt.Sprintf(`Set Project Verified Proposal:
	Title:       		 %s
	Description: 		 %s
	ProjectOwner: 		 %s
	ProjectId: 			 %d
  `, p.Title, p.Description, p.ProjectOwner, p.ProjectId)
}

func (p *SetProjectVerifiedProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(p.ProjectOwner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	return nil
}
