package cli

import (
	"clms/x/lms/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the LMS module",
		// DisableFlagParsing:         true,
		// SuggestionsMinimumDistance: 2,
		RunE: client.ValidateCmd,
	}

	cmd.AddCommand(
		ListAllTheLeavesCmd(),
		ListAllTheStudentCmd(),
	)

	return cmd
}
func ListAllTheLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listallleaves",
		Short: "List all the leaves",
		Long: `List all the leaves which are accepted or rejected by the admin,
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.ListAllTheLeavesRequest{}
			res, err := queryClient.ListAllTheLeaves(cmd.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func ListAllTheStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listallstudents",
		Short: "List all the students",
		Long: `List all the students which are added by admin,
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.ListAllTheStudentRequest{}
			res, err := queryClient.ListAllTheStudent(cmd.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
