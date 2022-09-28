package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nebula-labs/nebula/x/ido/types"
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

	cmd.AddCommand(GetEnableIDOCmd())
	cmd.AddCommand(GetCommitParticipationCmd())

	return cmd
}

func GetEnableIDOCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "enable-ido [project_id] [tokens] [token_listing_price] [allocation_limit]",
		Short:     "Enable IDO for a project",
		Args:      cobra.ExactArgs(4),
		ValidArgs: []string{"start_time"},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			tokens, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			token_listing_price, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			allocationLimit, err := types.ParseAllocationLimitArrayFromString(args[3])

			startTimeStr, err := cmd.Flags().GetString("start-time")
			if err != nil {
				return err
			}

			startTime, err := time.Parse(time.RFC3339, startTimeStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnableIDORequest(clientCtx.GetFromAddress().String(), uint64(projectId), tokens, token_listing_price, allocationLimit, startTime)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("start-time", time.Now().Add(time.Hour*24).Format(time.RFC3339), "add custom start time to this ido, if not specified, then it will start 1 day from now")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCommitParticipationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-participation [project-id] [tokens]",
		Short: "Commit coins to an IDO project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			projectId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			tokens, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCommitParticipationRequest(clientCtx.GetFromAddress().String(), uint64(projectId), tokens)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
