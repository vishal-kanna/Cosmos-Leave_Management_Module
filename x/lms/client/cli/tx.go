package cli

import (
	"clms/x/lms/types"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/bank/types"
)

//NewTxCmd returns a root CLI command handler for all x/lms transaction commands.
func NewTxCmd() *cobra.Command {
	// panic("txn")
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Leave Management commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewRegisterAdminCmd(),
		// NewMultiSendTxCmd(),
		NewAddStudentRequestCmd(),
		NewApplyLeaveRequestCmd(),
		NewAcceptLeaveRequestCmd(),
	)

	return txCmd
}

//NewRegisterAdminCmd returns CLI command handler for creating a RegisterAdmin method

func NewRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registeradmin",
		Args:  cobra.ExactArgs(3),
		Short: "Admin can register",
		Long: `Admin can register using address and name ,
		

		In order to Register admin first specify the address and next specify name of the admin
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			signer := args[0]
			name := args[1]
			address := args[2]
			msg := types.NewRegisterAdminRequest(signer, address, name)
			// return nil
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

//NewAddStudentRequestCmd returns CLI Command Handler for Adding the student
func NewAddStudentRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addstudent",
		Short: "Admin can add the student ",
		Long: `Admin need to register first in order to add the student and pass the student details which are need to added,
		

		the format that need to be given in command line is 
		first specify the admin address and 
		then pass the list of students{
			Name==args[1]
			Id=args[2]
			Address=args[3]
		}
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			signer := args[0]
			adminaddress := args[1]

			//create an array of students to store the students details
			// arrayOfStudents := *[]types.Student{}
			students := []*types.Student{}
			for i := 0; i < (len(args)-1)/3; i++ {

				student := &types.Student{
					Name:    args[3*i+2],
					Id:      args[3*i+3],
					Address: args[3*i+4],
				}
				students = append(students, student)
			}

			msg := types.NewAddStudentRequest(signer, adminaddress, students)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

//NewApplyLeaveRequestCmd returns a cobra Command for Handling the Applying leaves
func NewApplyLeaveRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "applyleave",
		Short: "Added student can apply the leave",
		Long: `Student can apply the leave,
		address | Reason | from | to
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			var format string = "2006-Jan-06"
			from, _ := time.Parse(format, args[3])
			to, _ := time.Parse(format, args[4])
			Address, _ := sdk.AccAddressFromBech32(args[1])
			Reason := args[2]
			From := &from
			To := &to
			Signer := args[0]

			msg := types.NewApplyLeaveRequest(Signer, Address.String(), Reason, *From, *To)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

//NewAcceptLeaveRequestCmd returns a CLI command that handles the Accepting the leaves
func NewAcceptLeaveRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acceptleave",
		Short: "Accept the leave",
		Long: `admin accepts the leave req
		Admin address | Student address		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			Signer := args[0]
			AdminAddress := args[1]
			StudentAddress := args[2]
			msg := types.NewAcceptLeaveRequest(Signer, AdminAddress, StudentAddress)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
