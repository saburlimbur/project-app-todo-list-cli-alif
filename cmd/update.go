package cmd

import (
	"fmt"
	"strconv"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status of a todo by its ID",
	Args:  cobra.ExactArgs(2), // 2 argument
	Run: func(cmd *cobra.Command, args []string) {
		// parsing id
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: ID must be a valid number.")
			return
		}

		status := args[1]

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
}
