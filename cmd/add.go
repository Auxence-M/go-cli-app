package cmd

import (
	"go-cli-app/todo"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add a new todo",
	Long: `add will create a new todo item to the list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)
		items  = append(items, item)
	}
	
	if err := todo.SaveItems(viper.GetString("datafile"), items); 
	err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
}