package cmd

import (
	"github.com/spf13/cobra"
)

// var shouldRoundUp bool
var deleteCmd = &cobra.Command{
    Use:     "delete",
    Aliases: []string{"delete", "del"},
    Short:   "Delete task",
    Long:    "Delete task",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
            err := deleteTask(args[0])
            cobra.CheckErr(err)
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}