package command

import (
	"cicd/cmd/utility"
	"os/exec"

	"github.com/spf13/cobra"
)

var flag = true
var CompileTailwind = &cobra.Command{
	Use:   "style",
	Short: "Compile tailwind-css",
	Long:  "Compile tailwind-css. Add -w option for watch mode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if value, flagErr := cmd.Flags().GetBool("watch"); flagErr != nil {
			return flagErr
		} else {
			var baseCommand = []string{"npx", "tailwindcss", "-i", "static/config.css", "-o", "static/tailwind.css"}
			var cmd *exec.Cmd
			if value {
				baseCommand = append(baseCommand, "--watch")
				cmd = utility.ExecuteCommand(baseCommand...)
			} else {
				cmd = utility.ExecuteCommand(baseCommand...)
			}
			if err := cmd.Run(); err != nil {
				return nil
			}
		}
		return nil
	},
}

func init() {
	CompileTailwind.Flags().BoolVarP(&flag, "watch", "w", false, "If it's watch mode")
}
