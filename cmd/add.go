package cmd

import (
	"fmt"
	"todolist_cli/service"

	"github.com/spf13/cobra"
)

var (
	title    string
	desc     string
	status   string
	priority string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  "Add a new todo item to the list",
	Run: func(cmd *cobra.Command, args []string) {
		todoService := service.NewTodoService()

		err := todoService.AddTodo(title, desc, status, priority)

		if err != nil {
			fmt.Println("Error:", err)
			return // stop print kalo ada validasi error service AddTodo
		}

		fmt.Println("Todo created succesfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&title, "title", "", "Judul tugas")
	addCmd.Flags().StringVar(&desc, "desc", "", "Deskripsi tugas")

	// default status = pending (karena service mengizinkan pending, in-progress, complete)
	addCmd.Flags().StringVar(&status, "status", "pending", "Status tugas (pending | in-progress | complete)")

	// default priority = medium
	addCmd.Flags().StringVar(&priority, "priority", "medium", "Priority tugas (low | medium | high)")
}
