package commands

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/listener"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// todo: Use core-builder to load
import (
	// Services should import by core-builder, this is temporary solution
	_ "github.com/p2pNG/core/services/status"
	_ "github.com/p2pNG/core/services/transfer"
)

var commandRun = &cobra.Command{
	Use:   "start",
	Short: "Encrypt your password, so that put in config file",
	Run:   commandRunExec,
}

func commandRunExec(_ *cobra.Command, _ []string) {

	db, err := database.GetDBEngine()
	if err != nil {
		logging.Log().Fatal("init database error", zap.Error(err))
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	plugins := core.GetRouterPluginRegistry()
	//todo: Replace with real config data
	x := "{\"BuildName\":\"Hello World\",\"LocalDiscoveryPort\":6553}"
	for _, plugin := range plugins {
		info := plugin.PluginInfo()
		// todo: delete
		logging.Log().Info("loading router plugin",
			zap.String("plugin", info.Name), zap.String("version", info.Version))
		//todo: Use Real config
		err := json.Unmarshal([]byte(x), plugin.Config())
		if err != nil {
			logging.Log().Fatal("load config for plugin failed", zap.Error(err), zap.String("plugin", info.Name))
		}

		err = plugin.Init()
		if err != nil {
			logging.Log().Fatal("init for plugin failed", zap.Error(err), zap.String("plugin", info.Name))
		}

		router.Mount(info.Prefix, plugin.GetRouter())

		err = db.Update(func(tx *bolt.Tx) error {
			for _, buk := range info.Buckets {
				_, err := tx.CreateBucketIfNotExists([]byte(buk))
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			logging.Log().Fatal("init database bucket failed", zap.Error(err), zap.String("plugin", info.Name))
		}
	}
	go func() {
		err = listener.ListenBoth(router, ":6480")
		if err != nil {
			logging.Log().Fatal("start http service failed", zap.Error(err))
		}
	}()
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}
