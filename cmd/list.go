package cmd

import (
	"doli/todo"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt  bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list todo items. ",
	Long: `list will print all the items on a todo list 
The priority of a todo is in parenthesis ranging from 1 to 3. 
1 being the highest priority and 2 being the default priority
Todos that are set to done are hidden by default`,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		fmt.Println(viper.GetString("datafile"))
		fmt.Println(fmt.Errorf("%v", err))
	}

	sort.Sort(todo.ByPriority(items))

	// GOOGLE THIS
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		switch priority {
		case 1:
			if allOpt && i.Priority == 1 {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
			if !allOpt && i.Priority == 1 && i.Done == doneOpt {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
		case 2:
			if allOpt && i.Priority == 2 {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
			if !allOpt && i.Priority == 2 && i.Done == doneOpt {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
		case 3:
			if allOpt && i.Priority == 3 {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
			if !allOpt && i.Priority == 3 && i.Done == doneOpt {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
		default:
			if allOpt || i.Done == doneOpt {
				fmt.Fprintln(w, i.GetPosition()+"\t"+i.GetPriority()+"\t"+i.Text+"\t"+i.DisplayDone()+"\t")
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&priority, "priority", "p", 0, "Priority: 1,2,3")
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'done' todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all the todos including the 'done' todos ")
}
