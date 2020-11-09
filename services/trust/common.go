package trust

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"net/http"
)

type coreTrustConfig struct {
	BuildName string
}

type coreTrustPlugin struct {
	config coreTrustConfig
}

type coreTrustContext struct {
	Config coreTrustConfig
}

func (p *coreTrustPlugin) Init() error {
	logging.Log().Info("Core Trust Plugin Init OK!")
	return nil
}

func (p *coreTrustPlugin) Config() interface{} {
	return &p.config
}

func (p *coreTrustPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/trust",
		Version: "0.0.0",
		Prefix:  "/trust",
		Buckets: []string{"test-bucket"},
	}
}

type contextType int

const (
	pluginContext contextType = iota
)

func (p *coreTrustPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), pluginContext, coreTrustContext{Config: p.config})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Get("/ocsp", ocspResponder)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreTrustPlugin{})
}

// NodeInfo described the basic info of a node. Used for peer discovery
type NodeInfo struct {
	Name      string
	Version   string
	BuildName string
}
