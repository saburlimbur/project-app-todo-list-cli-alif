package utils

import (
	"fmt"
	"os"
	"todolist_cli/model"

	"github.com/olekukonko/tablewriter"
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
	fmt.Println("==========================  List Tasks  ==========================")

	table := tablewriter.NewWriter(os.Stdout)

	// Header manual
	table.Append([]string{"No", "Task", "Status", "Priority"})

	// Garis pembatas manual (harus append juga)
	table.Append([]string{"---", "----------------", "--------------", "----------"})

	for i, t := range todos {
		row := []string{
			fmt.Sprintf("%d", i+1),
			t.Title,
			ColorStatus(t.Status),
			string(t.Priority),
		}
		table.Append(row)
	}

	table.Render()

	fmt.Println("==================================================================")
}
