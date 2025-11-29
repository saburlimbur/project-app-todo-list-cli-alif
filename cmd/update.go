package cmd

import (
	"fmt"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var todoID int
var todoStatus string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status of a todo by its ID",
	// Args:  cobra.ExactArgs(2), // 2 argument
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// parsing id
		if todoID <= 0 {
			fmt.Println("Error: ID must be a valid positive number.")
			return
		}

		id := todoID
		status := todoStatus

		todoService := service.NewTodoService()

		if err := todoService.UpdateTodoStatus(id, status); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Todo updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&todoID, "id", "i", 0, "ID of the todo item to update")
	updateCmd.Flags().StringVarP(&todoStatus, "status", "s", "", "New status (pending, in-progress, complete)")

	// wajib dgn flag
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("status")
}
