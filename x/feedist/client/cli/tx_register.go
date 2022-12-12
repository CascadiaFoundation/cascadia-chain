package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v9/x/feedist/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterFeedistProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-distribution [contract-address] [gas-shares] [reward-shares] [epochs]",
		Short: "Submit a proposal to register a contract feedist",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			argContract := args[0]
			argFeeShares := args[1]
			argRewardShares := args[2]
			argEpochs := args[3]

			feeShares, err := sdk.NewDecFromStr(argFeeShares)
			if err != nil {
				return err
			}

			rewardShares, err := sdk.NewDecFromStr(argRewardShares)
			if err != nil {
				return err
			}

			epochs, err := strconv.ParseUint(argEpochs, 10, 32)
			if err != nil {
				return err
			}

			content := types.NewRegisterFeedistProposal(
				title,
				description,
				argContract,
				feeShares,
				rewardShares,
				uint32(epochs),
			)

			from := clientCtx.GetFromAddress()
			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1uCC", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}

func CmdCancelFeedistProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-distribution [contract-address]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to cancel a contract incentive",
		Long:  "Submit a proposal to cancel a contract incentive.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			if !common.IsHexAddress(args[0]) {
				return fmt.Errorf("invalid contract address: %s", args[0])
			}

			contract := args[0]

			from := clientCtx.GetFromAddress()
			content := types.NewCancelFeedistProposal(title, description, contract)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1uCC", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}
