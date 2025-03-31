package cmd

import (
	"doli/todo"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
	Use:   "edit",
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

	if i > 0 && i <= len(items) {
		prevValue := items[i-1].Text

		if len(args) > 1 {
			items[i-1].Text = args[1]
			fmt.Println(prevValue, "changed to", args[1])
		}

		if priority > 0 && priority != items[i-1].Priority {
			items[i-1].SetPriority(priority)
			fmt.Println("Todo item number", args[0], "priority changed to", priority)
		}

		sort.Sort(todo.ByPriority(items))
		err := todo.SaveItems(viper.GetString("datafile"), items)
		if err != nil {
			log.Printf("%v", err)
		}

	} else {

		log.Println(i, "does not match any items")
	}

}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().IntVarP(&priority, "priority", "p", 0, "Priority: 1,2,3")
}
