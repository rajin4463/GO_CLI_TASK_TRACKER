/*
Copyright © 2025 NAME HERE rajin4463
*/
package cmd

import (
	"strconv"

	"github.com/rajin4463/GO_CLI_TASK_TRACKER/utils"
	"github.com/spf13/cobra"
)

// modCmd represents the mod command
var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "Modifiy tasks",
	Long:  `Modifiy tasks, Status, Due Date, Assigned to.`,
	Run: func(cmd *cobra.Command, args []string) {
		modTaskId, _ := strconv.Atoi(args[0])
		priority, _ := cmd.Flags().GetString("priority")
		status, _ := cmd.Flags().GetString("status")
		due, _ := cmd.Flags().GetString("due")
		Assigned, _ := cmd.Flags().GetString("assigned")

		data := utils.CsvRead("tasks.csv")
		for i := 1; i < len(data); i++ {
			dataID, _ := strconv.Atoi(data[i][0])
			if modTaskId == dataID {
				utils.CsvMod("tasks.csv", modTaskId, priority, status, due, Assigned)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(modCmd)

	modCmd.Flags().StringP("priority", "p", " ", "Modify Priority of Task")
	modCmd.Flags().StringP("status", "s", " ", "Status of Task")
	modCmd.Flags().StringP("due", "d", " ", "Modify due date of Task")
	modCmd.Flags().StringP("assigned", "a", " ", "Modify task assignment")
}
