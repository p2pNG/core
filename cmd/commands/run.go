package commands

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/p2pNG/core/internal/logging"
	"github.com/spf13/cobra"
)

var commandRun = &cobra.Command{
	Use:   "start",
	Short: "Encrypt your password, so that put in config file",
	Run:   commandRunExec,
}

func commandRunExec(c *cobra.Command, _ []string) {
	logging.Log().Info("Hello")
	logging.Log().Info(spew.Sdump(c))
}
