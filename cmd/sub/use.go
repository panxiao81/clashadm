package sub

import (
	"fmt"
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var UseCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Enable a config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		use(args[0])
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return getNameList(), cobra.ShellCompDirectiveNoFileComp
	},
}

func getNameList() []string {
	c, err := config2.NewConfigManager(viper.GetViper())
	if err != nil {
		panic(err)
	}
	var s []string
	for _, v := range c.SubscriptionUrl {
		s = append(s, v.Name)
	}
	return s
}

func use(n string) {
	c, err := config2.NewConfigManager(viper.GetViper())
	if err != nil {
		log.Fatal(err)
	}

	err = c.SubscriptionUrl.SetDefault(n)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat(filepath.Join(c.Installer.ConfigPath, n+".yaml")); os.IsNotExist(err) {
		log.Printf("config %s not exists, Download it", n)
		err := c.SubscriptionUrl.Update(n)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err = os.Lstat(filepath.Join(c.Installer.ConfigPath, "config.yaml")); err == nil {
		if err := os.Remove(filepath.Join(c.Installer.ConfigPath, "config.yaml")); err != nil {
			log.Fatal(err)
		}
	}
	os.Chdir(c.Installer.ConfigPath)
	os.Symlink(filepath.Join(c.Installer.ConfigPath, n+".yaml"), "config.yaml")

	fmt.Printf("Use config file '%s'\n", n)
}
