package transfer

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services"
	"net/http"
)

// SeedHashToSeedDB 		Key=SeedHash,Value=SeedInfo
// FileInfoHashToFileDB 	Key=FileInfoHash,Value=FileInfo
// FileHashToFileDB 		Key=FileHash,Value=FileInfo
// FileInfoHashToLocalFileDB 	Key=FileInfoHash,Value=LocalFileInfo
const (
	SeedHashToSeedDB          = "SeedInfoHash-SeedInfo"
	FileInfoHashToFileDB      = "FileInfoHash-FileInfo"
	FileHashToFileDB          = "FileHash-FileInfo"
	FileInfoHashToLocalFileDB = "FileInfoHash-LocalFileInfo"
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
		Buckets: []string{SeedHashToSeedDB, FileInfoHashToFileDB, FileHashToFileDB, FileInfoHashToLocalFileDB},
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
	r.Get("/seedInfo/{seedInfoHash}", getSeedInfo)
	r.Get("/fileInfo/fileInfoHash/{fileInfoHash}", getFileInfoByFileInfoHash)
	r.Get("/fileInfo/fileHash/{fileHash}", getFileInfoByFileHash)
	r.Get("/file/fileInfoHash/{fileInfoHash}/piece/{pieceIndex}/", getFilePiece)
	return r
}

func init() {
	core.RegisterRouterPlugin(&coreTransferPlugin{})
}
