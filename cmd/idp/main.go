package main

import (
	"github.com/choigonyok/idp/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "home-idp",
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			zap.S().Fatalln("Config file is required")
		}

		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			zap.S().Fatalf("%s: %s\n", "Cannot read config file", zap.Error(err))
		}

		zap.S().Infof("%s: %s", "Config loaded", viper.ConfigFileUsed())

		srv := server.NewServer()
		if err := srv.Run(); err != nil {
			zap.S().Fatalln("Cannot run server", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file path")
}

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
	zap.S().Infoln("Set global zap logger")

	if err := rootCmd.Execute(); err != nil {
		zap.S().Fatalln("Cannot execute cli", err)
	}
}
