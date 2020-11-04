package status

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/utils"
	"net/http"
)

type coreStatusPlugin struct {
}

func (p *coreStatusPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/status",
		Version: "0.0.0",
		Prefix:  "/status",
		Buckets: []string{"test-bucket"},
	}
}

func (p *coreStatusPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/info", getNodeInfo)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreStatusPlugin{})
}

func getNodeInfo(w http.ResponseWriter, r *http.Request) {

	node := NodeInfo{
		Name:    utils.GetHostname(),
		Version: core.GoModule().Version,
	}
	data, err := json.Marshal(&node)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

type NodeInfo struct {
	Name      string
	Version   string
	BuildName string
}
