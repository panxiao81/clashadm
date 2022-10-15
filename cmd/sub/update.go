package sub

import (
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var UpdateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "update config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config2.NewConfigManager(viper.GetViper())
		if err != nil {
			log.Fatal(err)
		}

		err = c.SubscriptionUrl.Update(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
