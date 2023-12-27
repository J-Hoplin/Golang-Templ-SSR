package command

import (
	"cicd/cmd/utility"
	"errors"
	"os/exec"

	"strings"

	"github.com/spf13/cobra"
)

var state string

var StartDB  = &cobra.Command{
	Use: "db",
	Short: "Start redrob database",
	Long: "Start redrob database. Recommend to compile tailwindcss first",
	RunE: func(command *cobra.Command, args []string) error  {
		if value, flagErr := command.Flags().GetString("state"); flagErr !=nil{
			return flagErr
		}else{
			var cmd *exec.Cmd
			switch strings.ToLower(value) {
			case "up":
				command := []string{"docker","compose","up","-d"}
				cmd = utility.ExecuteCommand(command...)
			case "down":
				command := []string{"docker-compose","down"}
				cmd = utility.ExecuteCommand(command...)
			default:
				return errors.New("INVALID_ARGUMENT")
			}
			if err := cmd.Run(); err != nil{
				return err
			}
		}
		return nil
	},
}

func init(){
	// For after sub command
	StartDB.Flags().StringVarP(&state,"state","s","","Initialize database. Argument should be one of 'up' or 'down'")
}