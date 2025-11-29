package cmd

import (
	"fmt"
	"strconv"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var (
	id int
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete single todo by id and status",
	Long:  "Delete a specific todo item by its ID",
	// ambil 1 argumen
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		// parsing id
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: ID must be a valid number.")
			return
		}

		todoService := service.NewTodoService()

		err = todoService.DeleteTodo(id)

		if err != nil {
			fmt.Println("Error deleting todo:", err)
			return
		}

		fmt.Printf("Todo with ID %d successfully deleted.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
