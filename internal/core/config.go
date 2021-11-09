package core

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type PriceFeederConfig struct {
	Debug bool
	port  int
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.AutomaticEnv() // read in environment variables that match

	// default values for config

	viper.SetDefault("price_feeder.service", "http://127.0.0.1")
	viper.SetDefault("price_feeder.expose_address", ":80")
	viper.SetDefault("prometheus.port", ":2112")

	viper.SetDefault("network_id", "3")
	viper.SetDefault("xcb_node_url", "http://127.0.0.1:8545")

	// end defaults

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
