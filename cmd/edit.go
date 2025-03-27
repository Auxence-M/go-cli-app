package cmd

import (
	"fmt"
	"go-cli-app/todo"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
Use: "edit",
Short: "edits a todo item",
Long: `edit will edit a todo item value. 
It takes only two argument the label or position of the todo item and its new value`,
Run: editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i> 0 && i <= len(items) {
		prevValue := items[i-1].Text
		items[i-1].Text = args[1]
		fmt.Println(prevValue, "changed to", args[1])

		sort.Sort(todo.ByPriority(items))
		todo.SaveItems(viper.GetString("datafile"), items)

	} else {
		
		log.Println(i, "does not match any items")
	}



}

func init() {
	rootCmd.AddCommand(editCmd)
}