package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multiplyCmd = &cobra.Command{
    Use:     "mark as done",
    Aliases: []string{"mark-done", "mark-completed"},
    Short:   "Mark as done",
    Long:    "Mark as done",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
            err := markDone(args[0])
            cobra.CheckErr(err)
            fmt.Println("Task marked as done")
    },
}

func init() {
    multiplyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
    rootCmd.AddCommand(multiplyCmd)
}