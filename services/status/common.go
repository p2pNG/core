package status

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"net/http"
	"strconv"
)

// SeedHashToPeerDB 	Key=SeedHash,Value=PeerInfo
// FileInfoHashToPeerDB Key=FileInfoHash,Value=PeerInfo
// FileHashToPeerDB 	Key=FileHash,Value=PeerInfo
const (
	SeedHashToPeerDB     = "SeedInfoHash-PeerInfo"
	FileInfoHashToPeerDB = "FileInfoHash-PeerInfo"
	FileHashToPeerDB     = "FileHash-PeerInfo"
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
		Buckets: []string{SeedHashToPeerDB, FileInfoHashToPeerDB, FileHashToPeerDB},
	}
}

func (p *coreStatusPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), services.StatusContext, coreStatusContext{Config: p.config})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Get("/node", getNodeStatus)
	r.Get("/peer", getNodePeers)
	r.Get("/seed", getNodeSeeds)
	r.Get("/fileHash", getNodeFileHash)
	r.Get("/fileInfoHash", getNodeFileInfoHash)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreStatusPlugin{})
}

type nodeInfo struct {
	Name      string
	Version   string
	BuildName string
}

// getPeerKey returns PeerInfo's key used in database
func getPeerKey(peer discovery.PeerInfo) (peerKey string) {
	return peer.Address.String() + ":" + strconv.Itoa(peer.Port)
}
