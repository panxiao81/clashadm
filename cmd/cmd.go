package cmd

import (
	"fmt"
	config2 "github.com/panxiao81/clashadm/config"
	install2 "github.com/panxiao81/clashadm/installer"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var Version = "Alpha v0.1"
var Commit = "HEAD"
var BuildTime = "Unknown"
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "clashadm",
	Short: "Clashadm is a deployment and management tool for clash",
	Long: `A tool for install, manage, and config file subscriber for clash.
	with a buildin dashboard
	more information on Github Page: https://github.com/panxiao81/clashadm`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Clashadm",
	Long:  `Print the version number of Clashadm, of course`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Clashadm, version ", Version, " -- ", Commit, " Build at: ", BuildTime)
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start clash",
	Long:  "Start clash through systemd, why not?",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() != 0 {
			log.Fatal("You must run clashadm start as root!")
		}
		err := install2.Start()
		if err != nil {
			log.Fatal("Start Failed, Run \"journalctl -xeu clash.service \" for details")
		}
		log.Printf("Clash has started successfully")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop clash",
	Long:  "Stop clash through systemd, why not?",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Geteuid() != 0 {
			log.Fatal("You must run clashadm stop as root!")
		}
		err := install2.Stop()
		if err != nil {
			log.Fatal("Stop Failed, Run \"journalctl -xeu clash.service \" for details")
		}
		log.Printf("Clash has stopped successfully")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file (default is /etc/clash/clashadm.yaml")
	rootCmd.AddCommand(installCmd, versionCmd, startCmd, stopCmd)
}

func initConfig() {
	config2.InitConfig(cfgFile)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
