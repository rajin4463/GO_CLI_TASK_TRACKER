/*
Copyright © 2025 NAME HERE rajin4463
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/rajin4463/GO_CLI_TASK_TRACKER/utils"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task [title]",
	Short: "Add a task",
	Long:  `Add a task with a title and optional priority.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the task title from the arguments
		title := args[0]

		// Get the priority flag
		priority, _ := cmd.Flags().GetString("priority")
		status, _ := cmd.Flags().GetString("status")
		due, _ := cmd.Flags().GetString("due")
		assign, _ := cmd.Flags().GetString("assign")

		newTask := []string{strconv.Itoa(getNextTaskID()), title, priority, status, due, assign}

		data := utils.CsvRead("tasks.csv")
		task := append(data[1:], newTask)

		// Write the new task to the file
		utils.CsvWrite("tasks.csv", task)
		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)

	// Add a priority flag to the task command
	taskCmd.Flags().StringP("priority", "p", " ", "Set the priority of the task (Low, Medium, High)")
	taskCmd.Flags().StringP("status", "s", "Not Started", "Set the status of task (Not Started, In Progress, Done, etc.)")
	taskCmd.Flags().StringP("due", "d", " ", "Set task Due Date")
	taskCmd.Flags().StringP("assign", "a", " ", "Assign task to someone")
}

// Helper function to get the next task ID
func getNextTaskID() int {
	// Use the CsvRead function to read the CSV file
	data := utils.CsvRead("tasks.csv")

	// Skip the header row and find the maximum ID
	maxID := 0
	for i := 1; i < len(data); i++ { // Start from 1 to skip the header
		if len(data[i]) > 0 {
			id, err := strconv.Atoi(data[i][0]) // Convert the ID to an integer
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}

	return maxID + 1
}
