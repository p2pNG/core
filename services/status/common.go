package status

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"net/http"
)

type coreStatusConfig struct {
	BuildName string
}

type coreStatusPlugin struct {
	config coreStatusConfig
}

type coreStatusContext struct {
	Config coreStatusConfig
}

func (p *coreStatusPlugin) Init() error {
	logging.Log().Info("Core Status Plugin Init OK!")
	return nil
}

func (p *coreStatusPlugin) Config() interface{} {
	return &p.config
}

func (p *coreStatusPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/status",
		Version: "0.0.0",
		Prefix:  "/status",
		Buckets: []string{"test-bucket"},
	}
}

type contextType int

const (
	pluginContext contextType = iota
)

func (p *coreStatusPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), pluginContext, coreStatusContext{Config: p.config})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Get("/info", getNodeInfo)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreStatusPlugin{})
}

// NodeInfo described the basic info of a node. Used for peer discovery
type NodeInfo struct {
	Name      string
	Version   string
	BuildName string
}
