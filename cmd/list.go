package cmd

import (
	"fmt"
	"go-cli-app/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "list todo items",
	Long: `list will print all the items on a todo list`,
	Run: lsitRun,
}

func lsitRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	fmt.Println(items)
}

func init() {
	rootCmd.AddCommand(listCmd)
}