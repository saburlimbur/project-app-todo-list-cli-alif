package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "Todolist CLI Application",
	Long:  "Todolist CLI is a simple CLI tool to manage your tasks, supporting add, list, mark done, and delete operations.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
