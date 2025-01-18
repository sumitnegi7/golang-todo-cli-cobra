package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:     "add [task] [status]",
    Aliases: []string{"add"},
    Short:   "Add a todo",
    Long:    "Carry out addition operation on a todo",
    Args:    cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // fmt.Printf("Arguments received: %s, %s\n", args[0],)
        id, err := addTodo(args[0], "false")
        cobra.CheckErr(err) // Automatically prints the error and exit
        fmt.Println("Todo added successfully.",id)
    },
}


func init() {
    rootCmd.AddCommand(addCmd)
}