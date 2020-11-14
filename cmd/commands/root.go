package commands

import (
	"fmt"
	"github.com/p2pNG/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile        string
	httpListenPort string
	bootstrapPeer  string

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
	clientCmd.Flags().StringVarP(&httpListenPort, "httpListenPort", "l", "", "")
	clientCmd.Flags().StringVarP(&bootstrapPeer, "bootstrapPeer", "b", "", "???")

	err := viper.BindPFlag("httpListenPort", clientCmd.PersistentFlags().Lookup("httpListenPort"))
	if err != nil {

	}
	err = viper.BindPFlag("bootstrapPeer", clientCmd.PersistentFlags().Lookup("bootstrapPeer"))
	if err != nil {

	}
	//rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(commandRun)
	//rootCmd.AddCommand(clientCmd)
}

func initConfig() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigName("p2pNG")
	// viper.AddConfigPath(utils.AppConfigDir())
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("error config file: %s \n", err)
	}
}

func getVersionStatement() string {
	mod := core.GoModule()
	return mod.Version
}
