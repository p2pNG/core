package status

import (
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"net/http"
)

// getNodeStatus returns the NodeInfo of current node
func getNodeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("statusCtx").(coreStatusContext)
	node := nodeInfo{
		Name:      utils.GetHostname(),
		Version:   core.GoModule().Version,
		BuildName: ctx.Config.BuildName,
	}
	services.WriteRespDataAsJson(w, &node)
}

// getNodePeers returns the peers of current node
func getNodePeers(w http.ResponseWriter, r *http.Request) {
	peers, err := discovery.GetPeerRegistry()
	if err == nil {
		services.WriteRespDataAsJson(w, &peers)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// getNodeSeeds returns the SeedInfo list of current node
func getNodeSeeds(w http.ResponseWriter, r *http.Request) {
	// todo:return seed list
}

// getNodeFileHash returns the FileHash list of current node
func getNodeFileHash(w http.ResponseWriter, r *http.Request) {
	// todo:return file list
}

// getNodeFileInfoHash returns the FileInfoHash list of current node
func getNodeFileInfoHash(w http.ResponseWriter, r *http.Request) {
	// todo:return file list
}
