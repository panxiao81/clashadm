package cmd

import (
	"github.com/panxiao81/clashadm/cmd/sub"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage clash config",
	Long:  "Manage clash config file",
}

var subscriptionCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe clash config",
	Long:  "Add, Modify or Remove clash config file subscription",
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(subscriptionCmd)
	subscriptionCmd.AddCommand(sub.AddCmd, sub.LsCmd, sub.UseCmd, sub.UpdateCmd)
}
