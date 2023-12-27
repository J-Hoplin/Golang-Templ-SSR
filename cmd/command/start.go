package command

import (
	"cicd/internal"

	"github.com/spf13/cobra"
)

var StartServer  = &cobra.Command{
	Use: "start",
	Short: "Start redrob server",
	Long: "Start redrob server. Recommend to compile tailwindcss first",
	Run: func(command *cobra.Command, args []string)  {
		internal.InitServer()
	},
}

func init(){
	// For after sub command
}
