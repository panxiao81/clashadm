package cmd

import (
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove clash and cleanup",
	Long:  "Remove clash binary and config file, maybe even clashadm database",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() != 0 {
			log.Fatal("You need to run clashadm as root!")
		}
		c, err := config2.NewConfigManager(viper.GetViper())
		if err != nil {
			panic(err)
		}
		err = c.Installer.Remove()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
