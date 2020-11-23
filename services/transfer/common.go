package transfer

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services"
	"net/http"
)

type coreTransferConfig struct {
	BuildName string
}

type coreTransferPlugin struct {
	config coreTransferConfig
}

type coreTransferContext struct {
	Config coreTransferConfig
}

func (p *coreTransferPlugin) Init() error {
	logging.Log().Info("Core Transfer Plugin Init OK!")
	return nil
}

func (p *coreTransferPlugin) Config() interface{} {
	return &p.config
}

func (p *coreTransferPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/transfer",
		Version: "0.0.0",
		Prefix:  "/transfer",
		Buckets: []string{services.SeedHashToSeedDB, services.FileInfoHashToFileDB, services.FileHashToFileDB, services.FileInfoHashToLocalFileDB},
	}
}

func (p *coreTransferPlugin) GetRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), services.TransferContext, coreTransferContext{Config: p.config})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Get("/seedInfo/{seedInfoHash}", serverGetSeedInfo)
	r.Get("/fileInfo/fileInfoHash/{fileInfoHash}", serverGetFileInfoByFileInfoHash)
	r.Get("/fileInfo/fileHash/{fileHash}", serverGetFileInfoByFileHash)
	r.Get("/file/fileInfoHash/{fileInfoHash}/piece/{pieceIndex}/", serverGetFilePiece)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreTransferPlugin{})
}
