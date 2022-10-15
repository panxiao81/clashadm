package cmd

import (
	"github.com/panxiao81/clashadm/dashboard"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Serve a yacd Dashboard",
	Long:  "Open a yacd Dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		serveDashboard()
	},
}

func serveDashboard() {
	dashboard.Serve()
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
