package cmd

import (
	"go-cli-app/todo"
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add a new todo",
	Long: `add will create a new todo item to the list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		items  = append(items, 
		todo.Item{Text: x})
	}
	
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}