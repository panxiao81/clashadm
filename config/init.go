package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const defaultConfig = `install: 
  path: /usr/local/bin
  release: premium
  config_path: /etc/clash
subscription_url: {}
mode: tun
mixed-in:
  mixed-port: 7890
  tun:
    enable: true
    stack: system
    auto-route: true
`

func setDefault() {
	viper.SetDefault("install", map[string]string{"path": "/usr/local/bin", "release": "premium", "config_path": "/etc/clash"})
	viper.SetDefault("subscription_url", []string{})
	viper.SetDefault("mode", "tun")
	viper.SetDefault("mixed-in.mixed-port", 7890)
}

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/clash")
		viper.SetConfigType("yaml")
		viper.SetConfigName("clashadm")
	}

	setDefault()
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Create Default Config file
			log.Printf("Config file not exists, Create one.")
			if _, err := os.Stat("/etc/clash"); os.IsNotExist(err) {
				err = os.Mkdir("/etc/clash", 0755)
				if err != nil {
					log.Fatal(err)
				}
			}
			f, err := os.OpenFile("/etc/clash/clashadm.yaml", os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err = f.WriteString(defaultConfig)
			if err != nil {
				log.Fatal(err)
			}
			err = viper.ReadInConfig()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}
