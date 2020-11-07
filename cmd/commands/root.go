package commands

import (
	"github.com/p2pNG/core"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "p2pNG-core",
	Short:   "An experimental p2pNG Core implement.",
	Version: getVersionStatement(),
}

// Execute returns the root command for main to start the whole application
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(commandRun)
}

func getVersionStatement() string {
	mod := core.GoModule()
	return mod.Version
}
