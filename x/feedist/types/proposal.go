package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ethermint "github.com/evmos/ethermint/types"
)

// constants
const (
	ProposalTypeRegisterFeedist string = "RegisterFeedist"
	ProposalTypeCancelFeedist   string = "CancelFeedist"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterFeedistProposal{}
	_ govtypes.Content = &CancelFeedistProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterFeedist)
	govtypes.RegisterProposalType(ProposalTypeCancelFeedist)
	govtypes.RegisterProposalTypeCodec(&RegisterFeedistProposal{}, "incentives/RegisterFeedistProposal")
	govtypes.RegisterProposalTypeCodec(&CancelFeedistProposal{}, "incentives/CancelFeedistProposal")
}

// NewRegisterIncentiveProposal returns new instance of RegisterIncentiveProposal
func NewRegisterFeedistProposal(
	title, description, contract string,
	feeShares sdk.Dec,
	rewardShares sdk.Dec,
	epochs uint32,
) govtypes.Content {
	return &RegisterFeedistProposal{
		Title:        title,
		Description:  description,
		Contract:     contract,
		Feeshares:    feeShares,
		Rewardshares: rewardShares,
		Epochs:       epochs,
	}
}

// ProposalRoute returns router key for this proposal
func (*RegisterFeedistProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*RegisterFeedistProposal) ProposalType() string {
	return ProposalTypeRegisterFeedist
}

// ValidateBasic performs a stateless check of the proposal fields
func (rfp *RegisterFeedistProposal) ValidateBasic() error {
	if err := ethermint.ValidateAddress(rfp.Contract); err != nil {
		return err
	}

	if err := validateShares(rfp.Feeshares); err != nil {
		return err
	}

	if err := validateShares(rfp.Rewardshares); err != nil {
		return err
	}

	if err := validateEpochs(rfp.Epochs); err != nil {
		return err
	}

	return govtypes.ValidateAbstract(rfp)
}

func validateShares(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("invalid parameter: nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("value cannot be negative: %T", i)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("value cannot be greater than 1: %T", i)
	}

	return nil
}

func validateEpochs(epochs uint32) error {
	if epochs == 0 {
		return fmt.Errorf("epochs value (%d) cannot be 0", epochs)
	}
	return nil
}

// NewCancelFeedistProposal returns new instance of RegisterFeedistProposal
func NewCancelFeedistProposal(
	title, description, contract string,
) govtypes.Content {
	return &CancelFeedistProposal{
		Title:       title,
		Description: description,
		Contract:    contract,
	}
}

// ProposalRoute returns router key for this proposal
func (*CancelFeedistProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*CancelFeedistProposal) ProposalType() string {
	return ProposalTypeCancelFeedist
}

// ValidateBasic performs a stateless check of the proposal fields
func (rip *CancelFeedistProposal) ValidateBasic() error {
	if err := ethermint.ValidateAddress(rip.Contract); err != nil {
		return err
	}

	return govtypes.ValidateAbstract(rip)
}
