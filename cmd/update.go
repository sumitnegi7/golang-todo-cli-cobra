package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:     "update",
    Aliases: []string{"update"},
    Short:   "Update todo task",
    Long:    "Update todo task description",
    Args:    cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string)  {
        err := updateTask(args[0], args[1]);  
        cobra.CheckErr(err) 
        fmt.Println("Task updated successfully")
    },
}


func init() {
    // divideCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
    rootCmd.AddCommand(updateCmd)
}