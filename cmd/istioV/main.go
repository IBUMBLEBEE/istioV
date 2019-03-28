package main

import (
	"fmt"
	"os"

	"IBUMBLEBEE/istioV/cmd/istioV/router"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"IBUMBLEBEE/istioV/cmd/istioV/bootstrap"
)

var (
	rootCmd = &cobra.Command{
		Use:               "istioV options",
		Short:             "istioV backend server",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Start istioV backend server`,
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	}

	startCmd = &cobra.Command{
		Use:     "start",
		Short:   "start istioV backend server",
		Example: "istioV -c config.json",
		RunE:    start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.ConfigFile, "config", "c", "conf/istiov.json",
		"Start server with peovided configuration file")
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.Host, "host", "H", "0.0.0.0",
		"Start server with provided port")
	startCmd.PersistentFlags().IntVarP(&bootstrap.Args.Port, "port", "p", 3083,
		"Start server with provided port")
	startCmd.PersistentFlags().BoolVarP(&bootstrap.Args.InCluster, "ink8sCluster", "i", false,
		"Start server in kube cluster")
	startCmd.PersistentFlags().StringVarP(&bootstrap.Args.IstioNamespace, "istioNamespace", "I", "istio-system",
		"wiith provided deployed Istio namespace")

	rootCmd.AddCommand(startCmd)
	// rootCmd.AddCommand(version.Command())
}

func start(_ *cobra.Command, _ []string) error {
	parseConfig()

	// wait log init
	// mode := viper.GetString("mode")
	// gin.SetMode("debug")

	// setting quit single
	// quitch := make(chan os.Signal)
	// signla.Notify(quitch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	engine := gin.Default()
	// engine.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// engine.Run()
	// executor.Init()
	router.Executor(engine)
	engine.Run()
	// service.Init()
	// server := &http.Server{
	// 	Addr: fmt.Sprintf("%s:%d", bootstrap.Args.Host, bootstrap.Args.Port),
	// 	Handler: engine,
	// }

	// graceful shutdown server
	// go func() {}()
	return nil
}

func parseConfig() {
	viper.SetConfigFile(bootstrap.Args.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("paser config file failed: %s", err))
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("boostrap args",
		zap.Int("Port", bootstrap.Args.Port),
		zap.Bool("Ink8sCluster", bootstrap.Args.InCluster),
		zap.String("ConfigFile", bootstrap.Args.ConfigFile),
		zap.String("IstioNamespace", bootstrap.Args.IstioNamespace))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
