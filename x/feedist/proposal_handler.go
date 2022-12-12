package feedist

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/evmos/v9/x/feedist/keeper"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

// NewFeedistProposalHandler creates a governance handler to manage new
// proposal types.
func NewFeedistProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterFeedistProposal:
			return handleRegisterFeedistProposal(ctx, k, c)
		case *types.CancelFeedistProposal:
			return handleCancelFeedistProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"unrecognized %s proposal content type: %T", types.ModuleName, c,
			)
		}
	}
}

func handleRegisterFeedistProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterFeedistProposal) error {
	in, err := k.RegisterFeedist(ctx, common.HexToAddress(p.Contract), p.Feeshares, p.Rewardshares, p.Epochs)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterFeedist,
			sdk.NewAttribute(types.AttributeKeyContract, in.Contract),
			sdk.NewAttribute(
				types.AttributeKeyEpochs,
				strconv.FormatUint(uint64(in.Epochs), 10),
			),
		),
	)
	return nil
}

func handleCancelFeedistProposal(ctx sdk.Context, k *keeper.Keeper, p *types.CancelFeedistProposal) error {
	err := k.CancelFeedist(ctx, common.HexToAddress(p.Contract))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCancelFeedist,
			sdk.NewAttribute(types.AttributeKeyContract, p.Contract),
		),
	)
	return nil
}
