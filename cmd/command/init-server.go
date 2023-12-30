package command

import (
	"cicd/internal"

	"github.com/spf13/cobra"
)

var InitServer = &cobra.Command{
	Use:   "server",
	Short: "Start redrob server",
	Long:  "Start redrob server. Recommend to compile tailwindcss first",
	RunE: func(command *cobra.Command, args []string) error {
		internal.InitServer()
		return nil
	},
}

func init() {
	// For after sub command
}
