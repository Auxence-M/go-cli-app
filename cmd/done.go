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

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "set todo items to done",
	Long: `done will set a todo item on a todo list to done.
The label of a todo item is its position ranging from 1 to n`,
	Run: doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

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
	rootCmd.AddCommand(doneCmd)
}
