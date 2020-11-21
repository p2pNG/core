package commands

import (
	"fmt"
	"github.com/p2pNG/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	httpListen    uint16
	bootstrapPeer string

	rootCmd = &cobra.Command{
		Use:     "p2pNG",
		Short:   "An experimental p2pNG Core implement.",
		Version: getVersionStatement(),
	}
)

// Execute returns the root command for main to start the whole application
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	commandRun.Flags().Uint16VarP(&httpListen, "http-listen", "l", 0, "")
	commandRun.Flags().StringVarP(&bootstrapPeer, "bootstrap-peer", "b", "", "???")

	//rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(commandRun)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigName("p2pNG")
	// viper.AddConfigPath(utils.AppConfigDir())
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error config file: %s \n\n", err)
	}
}

func getVersionStatement() string {
	mod := core.GoModule()
	return mod.Version
}
