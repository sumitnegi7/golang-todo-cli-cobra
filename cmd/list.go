package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:     "list",
    Aliases: []string{"list"},
    Short:   "List all todos",
    Long:    "List all todos",
    Args:    cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        err := listTodo()
        cobra.CheckErr(err)
        fmt.Printf("List" )
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}