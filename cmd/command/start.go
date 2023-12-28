package command

import (
	"cicd/cmd/utility"
	"cicd/internal"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var StartServer  = &cobra.Command{
	Use: "start",
	Short: "Start redrob server",
	Long: "Start redrob server. Recommend to compile tailwindcss first",
	RunE: func(command *cobra.Command, args []string) error  {
		var baseCommand = []string{"templ","generate"}
		var tailwindCommand = []string{"npx", "tailwindcss", "-i", "static/config.css", "-o", "static/tailwind.css"}
		var cmd *exec.Cmd
		var cmd2 *exec.Cmd
		cmd = utility.ExecuteCommand(baseCommand...)
		cmd2 = utility.ExecuteCommand(tailwindCommand...)

		if err := cmd.Run(); err != nil {
			return nil
		}else{
			log.Println("Templ components compiled!")
			
		}

		if err := cmd2.Run();err != nil {
			return nil
		}else{
			log.Println("Tailwind CSS compiled!")
		}
		internal.InitServer()
		return nil
	},
}

func init(){
	// For after sub command
}
