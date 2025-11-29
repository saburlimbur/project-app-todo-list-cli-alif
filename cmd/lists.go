package cmd

import (
	"fmt"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var listsCmd = &cobra.Command{
	Use:   "lists",
	Short: "List of todos",
	Long:  "Display all active and completed todo items",
	Run: func(cmd *cobra.Command, args []string) {
		todoService := service.NewTodoService()

		todos, err := todoService.ListTodo()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(todos) == 0 {
			fmt.Print("todos 0 found")
		}

		fmt.Println("=== Daftar Todo ===")
		for _, tds := range todos {
			fmt.Printf(
				"ID: %d | Title: %s | Description: %s | Status: %s | Priority: %s\n",
				tds.Id, tds.Title, tds.Description, tds.Status, tds.Priority,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listsCmd)
}
