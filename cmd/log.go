/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log of your activities.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// log --duration="1:30"
	logCmd.Flags().StringP("duration", "d", "", "duration of the activity you are about to log.")
	// log --start="09:30"
	logCmd.Flags().StringP("start", "s", "", "start time of the activity.")
	// log --description="something that i used to do"
	logCmd.Flags().StringP("description", "", "", "description of the activity.")

	// _ = logCmd.MarkFlagRequired("duration")
}
