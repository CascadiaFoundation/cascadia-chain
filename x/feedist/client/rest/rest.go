package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/evmos/evmos/v9/x/feedist/types"
)

// RegisterFeedistProposalRequest defines a request for a new register a
// contract Feedist.
type RegisterFeedistProposalRequest struct {
	BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title           string       `json:"title" yaml:"title"`
	Description     string       `json:"description" yaml:"description"`
	Deposit         sdk.Coins    `json:"deposit" yaml:"deposit"`
	ContractAddress string       `json:"contract_address" yaml:"contract_address"`
	FeeShares       sdk.Dec      `json:"fee_shares" yaml:"feeshares"`
	RewardShares    sdk.Dec      `json:"reward_shares" yaml:"rewardshares"`
	Epochs          uint32       `json:"epochs" yaml:"epochs"`
}

// CancelFeedistProposalRequest defines a request for a new register a
// contract Feedist.
type CancelFeedistProposalRequest struct {
	BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title           string       `json:"title" yaml:"title"`
	Description     string       `json:"description" yaml:"description"`
	Deposit         sdk.Coins    `json:"deposit" yaml:"deposit"`
	ContractAddress string       `json:"contract_address" yaml:"contract_address"`
}

func RegisterFeedistProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterFeedistProposalHandler(clientCtx),
	}
}

func CancelFeedistProposalRequestRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newCancelFeedistProposalRequestHandler(clientCtx),
	}
}

func newRegisterFeedistProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterFeedistProposalRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		contract := req.ContractAddress

		content := types.NewRegisterFeedistProposal(req.Title, req.Description, contract, req.FeeShares, req.RewardShares, req.Epochs)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newCancelFeedistProposalRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CancelFeedistProposalRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		contract := req.ContractAddress

		content := types.NewCancelFeedistProposal(req.Title, req.Description, contract)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
