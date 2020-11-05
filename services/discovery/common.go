package discovery

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"net/http"
)

type coreDiscoveryConfig struct {
	LocalDiscoveryPort int
}

type coreDiscoveryPlugin struct {
	config coreDiscoveryConfig
}

type coreDiscoveryContext struct {
	Config coreDiscoveryConfig
}

func (p *coreDiscoveryPlugin) Init() error {
	var err error
	go func() {
		_, err = LocalBroadcast(p.config.LocalDiscoveryPort)
		if err == nil {
			logging.Log().Info("init mDNS service ok")
		}
	}()
	return err
}

func (p *coreDiscoveryPlugin) Config() interface{} {
	return &p.config
}

func (p *coreDiscoveryPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/discovery",
		Version: "0.0.0",
		Prefix:  "/discovery",
		Buckets: []string{"discovery_registry"},
	}
}

func (p *coreDiscoveryPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "statusCtx", coreDiscoveryContext{Config: p.config})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreDiscoveryPlugin{})
}

type NodeInfo struct {
	Name      string
	Version   string
	BuildName string
}
