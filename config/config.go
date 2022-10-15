package config

import (
	"errors"
	"github.com/panxiao81/clashadm/installer"
	"github.com/spf13/viper"
)

type Config struct {
	Installer       installer.InstallArgs `mapstructure:"install"`
	SubscriptionUrl Subscription          `mapstructure:"subscription_url"`
	Mode            string                `mapstructure:"mode"`
	Clash           ClashConfig           `mapstructure:"mixed-in"`
}

func (s *Config) check() error {
	if viper.Get("install.release") == "clash" && viper.Get("mode") != "redir" {
		return errors.New("Config File is invilded")
	}
	if _, ok := installer.ReleaseUrl[viper.GetString("install.release")]; !ok {
		return errors.New("Not support this release")
	}
	return nil
}

func NewConfigManager(v *viper.Viper) (*Config, error) {
	var C Config
	err := v.Unmarshal(&C)
	if err != nil {
		return nil, err
	}
	err = C.check()
	if err != nil {
		return nil, err
	}
	return &C, nil
}
