package sub

import (
	config2 "github.com/panxiao81/clashadm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var AddCmd = &cobra.Command{
	Use:   "add [name] [url]",
	Short: "Add clash subscription url",
	Args:  cobra.ExactArgs(2),
	Run:   Add,
}

func Add(cmd *cobra.Command, args []string) {
	c, err := config2.NewConfigManager(viper.GetViper())
	if err != nil {
		log.Fatal(err)
	}
	err = c.SubscriptionUrl.Add(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}
}
