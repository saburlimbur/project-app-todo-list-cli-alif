package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"todolist_cli/model"
)

var (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func ColorStatus(status model.TodoStatus) string {
	switch status {
	case model.StatusComplete:
		return Green + string(status) + Reset
	case model.StatusProgress:
		return Yellow + string(status) + Reset
	case model.StatusPending:
		return Blue + string(status) + Reset
	default:
		return string(status)
	}
}

func PrintTodoTable(todos []model.Todos) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Println("=================================  List Tasks  =================================")

	fmt.Fprintf(w, "No\tTask\tStatus\tPriority\t\n")
	fmt.Fprintf(w, "----\t----------------\t--------------\t-----------\t\n")

	for i, t := range todos {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t\n",
			i+1,
			t.Title,
			ColorStatus(t.Status),
			t.Priority,
		)
	}

	w.Flush()
	fmt.Println("================================================================================")
}
