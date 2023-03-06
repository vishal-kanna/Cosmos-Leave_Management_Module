package cli

import (
	"clms/x/lms/types"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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
		Args:  cobra.ExactArgs(2),
		Short: "Admin can register",
		Long: `Admin can register using address and name ,
		[adminname]|[address]

		In order to Register admin first specify the address and next specify name of the admin
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			signer := clientCtx.GetFromAddress()
			fmt.Println("the signer address", signer)
			name := args[0]
			address := args[1]

			msg := types.NewRegisterAdminRequest(signer.String(), address, name)
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
		
		[adminaddress]|[studentname] |[id]|[address]
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
			signer := clientCtx.GetFromAddress()
			adminaddress := args[0]

			//create an array of students to store the students details
			// arrayOfStudents := *[]types.Student{}
			students := []*types.Student{}
			for i := 0; i < (len(args)-1)/3; i++ {

				student := &types.Student{
					Name:    args[3*i+1],
					Id:      args[3*i+2],
					Address: args[3*i+3],
				}
				students = append(students, student)
			}

			msg := types.NewAddStudentRequest(signer.String(), adminaddress, students)
			// panic("called")

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
			from, _ := time.Parse(format, args[2])
			to, _ := time.Parse(format, args[3])
			Address := args[0]
			Reason := args[1]
			From := &from
			To := &to
			sender := clientCtx.GetFromAddress()

			msg := types.NewApplyLeaveRequest(sender.String(), Address, Reason, *From, *To)
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
		[signer] | [adminaddress] | [studentid]
		Admin address | Student address		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			Signer := clientCtx.GetFromAddress()
			AdminAddress := args[0]
			Studentid := args[1]
			msg := types.NewAcceptLeaveRequest(Signer.String(), AdminAddress, Studentid)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
