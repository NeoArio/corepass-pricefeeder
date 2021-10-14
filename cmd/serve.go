package cmd

import (
	"fmt"
	"github.com/NeoArio/corepass-pricefeeder/internal/controller"
	"github.com/NeoArio/corepass-pricefeeder/internal/core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
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

	prometheusServer := http.NewServeMux()
	prometheusServer.Handle("/metrics", promhttp.Handler())
	prometheusPort := viper.GetString("prometheus.port")

	priceFeederApiServer := http.NewServeMux()
	priceFeederApiServer.HandleFunc("/", http.NotFound)
	priceFeederApiServer.HandleFunc("/health_check", controller.HealthCheckHandle)
	exposeAddress := viper.GetString("pricefeeder.expose_address")

	fmt.Println("start priceFeeder API Server on " + exposeAddress)
	fmt.Println("start Prometheus Server on " + prometheusPort)



	go func() {
		httpError := http.ListenAndServe(exposeAddress, priceFeederApiServer)
		if httpError != nil {
			fmt.Println("While serving HTTP: ", httpError)
		}
	}()

	go func() {
		httpError := http.ListenAndServe(prometheusPort, prometheusServer)
		if httpError != nil {
			fmt.Println("While serving HTTP: ", httpError)
		}
	}()
	select {}
}
