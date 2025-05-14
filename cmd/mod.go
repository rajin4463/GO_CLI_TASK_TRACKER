/*
Copyright © 2025 NAME HERE rajin4463
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// modCmd represents the mod command
var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "Modifiy tasks",
	Long:  `Modifiy tasks, Status, Due Date, Assigned to.`,
	Run: func(cmd *cobra.Command, args []string) {
		modTaskId, _ := strconv.Atoi(args[0])
		status, _ := cmd.Flags().GetString("status")
		due, _ := cmd.Flags().GetString("due")
		Assigned, _ := cmd.Flags().GetString("assigned")
	},
}

func init() {
	rootCmd.AddCommand(modCmd)

	modCmd.Flags().StringP("status", "s", "Not Started", "Modifiy Status of Task")
}
