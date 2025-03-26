package cmd

import (
	"fmt"
	"go-cli-app/todo"
	"os"
	"sort"
	"text/tabwriter"

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

	sort.Sort(todo.ByPriority(items))

	// GOOGLE THIS
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t")
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}