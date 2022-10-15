package sub

import (
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list subscription URL",
	Run: func(cmd *cobra.Command, args []string) {
		ls()
	},
}

func ls() {
	c, err := config2.NewConfigManager(viper.GetViper())
	if err != nil {
		log.Fatal(err)
	}
	c.SubscriptionUrl.List()
}
