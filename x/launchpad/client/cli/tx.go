package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/nebula-labs/nebula/utils"
	"github.com/nebula-labs/nebula/x/launchpad/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateProject())
	cmd.AddCommand(CmdDeleteProject())
	cmd.AddCommand(CmdWithdrawAllTokens())
	cmd.AddCommand(CmdSubmitSetProjectVerifiedProposal())

	return cmd
}

func CmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [project_title] [project_information]",
		Short: "Create a new project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectTitle := args[0]
			projectInformation := args[1]

			msg := types.NewMsgCreateProjectRequest(clientCtx.GetFromAddress().String(), projectTitle, projectInformation)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-project [project_id]",
		Short: "Delete project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProjectRequest(clientCtx.GetFromAddress().String(), projectId.Uint64())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdWithdrawAllTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-all-tokens [project_id]",
		Short: "Withdraw all tokens of a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawAllTokensRequest(clientCtx.GetFromAddress().String(), projectId.Uint64())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSubmitSetProjectVerifiedProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-project-verified [project-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a set proposal to validate this launchpad project",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			proposal, err := utils.ParseProposalFlags(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse proposal: %w", err)
			}

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			content := types.NewSetProjectVerifiedProposal(proposal.Title, proposal.Description, from.String(), projectId.Uint64())

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")

	return cmd
}
