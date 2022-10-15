package cmd

import (
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install clash",
	Run:   install,
}

var path string
var distro string
var config string

func init() {
	installCmd.Flags().StringVarP(&path, "dir", "d", "", "Which directory to install")
	installCmd.Flags().StringVarP(&distro, "release", "r", "", "Which release to install, support: clash, premium, tun")
	viper.BindPFlag("install.path", installCmd.Flag("dir"))
	viper.BindPFlag("install.release", installCmd.Flag("release"))
}

func install(cmd *cobra.Command, args []string) {
	c, err := config2.NewConfigManager(viper.GetViper())
	if err != nil {
		panic(err)
	}
	c.Installer.Install()
}
