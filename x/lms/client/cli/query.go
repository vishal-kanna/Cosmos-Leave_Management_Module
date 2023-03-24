package cli

import (
	"clms/x/lms/types"
	"log"

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
		ListStudentLeaveStatus(),
		Admin(),
	)

	return cmd
}
func ListAllTheLeavesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listallleaves",
		Short: "List all the leaves",
		Long: `List all the leaves which are accepted or rejected by the admin,
		`,
		Example: "./simd query lms listallleaves",

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.ListAllTheLeavesRequest{}
			res, err := queryClient.ListAllTheLeaves(cmd.Context(), params)
			// fmt.Println("the result is", res)
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
		Example: "./simd query lms listallstudents",
		RunE: func(cmd *cobra.Command, args []string) error {
			// panic("called0")
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.ListAllTheStudentRequest{}
			res, err := queryClient.ListAllTheStudent(cmd.Context(), params)
			if err != nil {
				log.Println("the error is", err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func ListStudentLeaveStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "studentleavestatus",
		Short: "List all the students",
		Long: `List all the students which are added by admin,
		`,
		Example: "./simd query lms studentleavestatus [studentid]",

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			studentaddress := args[0]
			leaveid := args[1]
			params := &types.GetLeaveStatusRequest{
				Studentaddress: studentaddress,
				Leaveid:        leaveid,
			}
			res, err := queryClient.GetLeaveStatus(cmd.Context(), params)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func Admin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Lists the admin ",
		Long: `Lists the admin details,
		[adminaddress]
		`,
		Example: "./simd query lms admin",

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			req := types.GetadminRequest{}
			res, err := queryClient.GetAdmin(cmd.Context(), &req)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// ./simd tx leavemanagementsystem addstudent cosmos12f70wvzjtwlyrralxat3xqc86fz7rjvuk3jrvx
// cosmos1dmcr5va59frvjazgsvj3w8fcf83u06wp3fgd8l vinni  b172263 cosmos1u7z2zw86ryycpj4y0jp0p8qdlye8a3zcme7p6e --from validator-key --chain-id testnet
