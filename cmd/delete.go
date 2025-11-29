package cmd

import (
	"fmt"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var (
	todoId int
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete single todo by id and status",
	Long:  "Delete a specific todo item by its ID",
	// ambil 1 argumen
	// Args: cobra.ExactArgs(1),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// parsing id
		if todoId <= 0 {
			fmt.Println("Error: ID must be a valid positive number.")
			return
		}

		id := todoId
		todoService := service.NewTodoService()

		err := todoService.DeleteTodo(id)

		if err != nil {
			fmt.Println("Error deleting todo:", err)
			return
		}

		fmt.Printf("Todo with ID %d successfully deleted.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&todoId, "id", "i", 0, "ID of the todo item to deleted")

	// wajib dgn flag
	updateCmd.MarkFlagRequired("id")
}
