package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v9/x/feedist/types"
)

func (k Keeper) RegisterFeedist(
	ctx sdk.Context,
	contract common.Address,
	feeShares sdk.Dec,
	rewardShares sdk.Dec,
	epochs uint32,
) (*types.Feedist, error) {
	params := k.GetParams(ctx)
	if !params.EnableFeedist {
		return nil, types.ErrFeedistEnable
	}

	// contract must already be deployed, to avoid spam registrations
	contractAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, contract)

	if contractAccount == nil || !contractAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRevenueNoContractDeployed,
			"no contract code found at address %s", contract,
		)
	}

	feedist := types.Feedist{
		Index:        "feedist",
		Contract:     contract.String(),
		Feeshares:    feeShares,
		Rewardshares: rewardShares,
		Epochs:       epochs,
		StartTime:    ctx.BlockTime(),
	}
	k.SetFeedist(ctx, feedist)

	return &feedist, nil
}

func (k Keeper) CancelFeedist(
	ctx sdk.Context,
	contract common.Address,
) error {
	// Check if the Incentives are globally enabled
	params := k.GetParams(ctx)
	if !params.EnableFeedist {
		return sdkerrors.Wrap(
			types.ErrFeedistEnable,
			"incentives are currently disabled by governance",
		)
	}

	_, found := k.GetFeedist(ctx, "feedist")
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidAddress,
			"unmatching contract '%s' ", contract,
		)
	}

	k.RemoveFeedist(ctx, "feedist")

	return nil
}
