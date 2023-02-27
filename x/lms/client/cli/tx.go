package cli

import (
	"clms/x/lms/types"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func NewTxCmd() *cobra.Command {
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
		NewAddStudentRequest(),
		// NewAddStudentRequest
		NewApplyLeaveRequest(),
	)

	return txCmd
}

//NewRegisterAdmin returns CLI command handler for creating a RegisterAdmin method or transaction

func NewRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "this is used to register the admin ",
		Short: "Admin can register",
		Long: `Admin can register using address and name ,
		

		In order to Register admin first specify the address and next specify name of the admin
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			address := sdk.AccAddress(args[1])
			types.NewRegisterAdminRequest(address, name)
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAddStudentRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "this is used to by the admin to add the student ",
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
			adminaddress := args[0]
			// address := sdk.AccAddress(args[1])
			// types.NewRegisterAdminRequest(address, name)

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
			types.NewAddStudentRequest(adminaddress, students)
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewApplyLeaveRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "this cli is used to apply the leave",
		Short: "Added student can apply the leave",
		Long: `Student can apply the leave,
		address | Reason | from | to
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var format string = "2006-Jan-06"
			from, _ := time.Parse(format, args[2])
			to, _ := time.Parse(format, args[3])
			Address := args[0]
			Reason := args[1]
			From := &from
			To := &to
			types.NewApplyLeaveRequest(Address, Reason, *From, *To)
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAcceptLeaveRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "This is used to accept the leave",
		Short: "Accept the leave",
		Long: `admin accepts the leave req
		Admin address | Student address		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			AdminAddress := args[0]
			StudentAddress := args[1]
			types.NewAcceptLeaveRequest(AdminAddress, StudentAddress)
			return nil
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
