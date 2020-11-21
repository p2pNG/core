package commands

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/listener"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	x := "{\"BuildName\":\"Hello World\"}"
	for _, plugin := range plugins {
		info := plugin.PluginInfo()
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
		l := "6443"
		b := "bootstrap-peer.p2png.org"

		if fmt.Sprintf("%d", httpListen) != "0" {
			l = fmt.Sprintf("%d", httpListen)
		}
		if viper.GetString("http-listen") != "0" {
			l = viper.GetString("http-listen")
		}

		if fmt.Sprintf("%s", bootstrapPeer) != "" {
			b = fmt.Sprintf("%s", bootstrapPeer)
		}
		if viper.GetString("bootstrap-peer") != "" {
			b = viper.GetString("bootstrap-peer")
		}

		err = listener.ListenBoth(router, ":"+l)
		if err != nil {
			logging.Log().Fatal("start http service failed", zap.Error(err))
		}
		err = listener.ListenBoth(router, b)
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
