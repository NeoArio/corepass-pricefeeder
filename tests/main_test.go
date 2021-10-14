package tests

import (
	"flag"
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/spf13/viper"
	"os"
	"testing"
)

var configFilePath = flag.String("config", "", "")

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	// Read Config from file
	flag.Parse()
	if *configFilePath !="" {
		viper.SetConfigFile(*configFilePath)

	} else {
		viper.SetConfigFile("config.yaml")

	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	core.InitConfig()
	fmt.Println("Config finished")
	os.Exit(m.Run())
}
