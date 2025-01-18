/*
The root.go file uses the Cobra package to create the root command
and an accessor function to execute the command.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "Todo app: Add, Update, and Delete tasks",
	Long: `The application should run from the command line, accept user actions and inputs as arguments, and store the tasks in a JSON file. The user should be able to:
	Add, Update, and Delete tasks`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}



