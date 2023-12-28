package cmd

import (
	"cicd/cmd/command"
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "redrob",
	Short: "Redrob service CLI",
	Long:  `Utility CLI command tool for development`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Utility CLI")
	},
}

func init() {
	RootCmd.AddCommand(command.StartServer)
	RootCmd.AddCommand(command.CompileTailwind)
	RootCmd.AddCommand(command.StartDB)
	RootCmd.AddCommand(command.GenerateDummy)
}

func Execute() {
	RootCmd.Execute()
}
