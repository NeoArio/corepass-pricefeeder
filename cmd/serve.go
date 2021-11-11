package cmd

import (
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/controller"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serve http",
	Long:  `This command serve rest api`,
	Run:   serve,
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	// Read Config from file
	if configFile, err := cmd.Flags().GetString("config"); err == nil && configFile != "" {
		viper.SetConfigFile(configFile)

	} else {
		viper.SetConfigFile("config.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	core.InitConfig()

	priceFeederApiServer := http.NewServeMux()
	priceFeederApiServer.HandleFunc("/", http.NotFound)
	priceFeederApiServer.HandleFunc("/health_check", controller.HealthCheckHandle)
	priceFeederApiServer.HandleFunc("/core/get", controller.GetSimpleStorageCoreCoin)
	exposeAddress := viper.GetString("pricefeeder.expose_address")

	fmt.Println("start priceFeeder API Server on " + exposeAddress)

	go func() {
		httpError := http.ListenAndServe(exposeAddress, priceFeederApiServer)
		if httpError != nil {
			fmt.Println("While serving HTTP: ", httpError)
		}
	}()

	select {
	case <-time.After(time.Second * 10):
		log.Info("Add price process started...")
		controller.AddPrice()
	}
}
