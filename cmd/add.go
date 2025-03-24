package cmd

import (
	"fmt"
	"go-cli-app/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add a new todo",
	Long: `add will create a new tofo item to the list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items := []todo.Item{}

	for _, x := range args {
		items  = append(items, 
		todo.Item{Text: x})
	}
	
	err := todo.SaveItems("todo.json", items)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}