package cmd

import (
	"go-cli-app/todo"
	"log"

	"github.com/spf13/cobra"
)

var priority int

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
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items  = append(items, item)
	}
	
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
}