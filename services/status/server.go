package status

import (
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"github.com/p2pNG/core/services/transfer"
	"net/http"
)

// getNodeStatus returns the NodeInfo of current node
func getNodeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(services.StatusContext).(coreStatusContext)
	node := nodeInfo{
		Name:      utils.GetHostname(),
		Version:   core.GoModule().Version,
		BuildName: ctx.Config.BuildName,
	}
	services.WriteRespDataAsJSON(w, &node)
}

// getNodePeers returns the peers of current node
func getNodePeers(w http.ResponseWriter, r *http.Request) {
	peers, err := discovery.GetPeerRegistry()
	if err == nil {
		services.WriteRespDataAsJSON(w, &peers)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// getNodeSeeds returns the SeedInfo list of current node
func getNodeSeeds(w http.ResponseWriter, r *http.Request) {
	seedHashList, err := transfer.GetSeedInfoHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, seedHashList)
	}
}

// getNodeFileHash returns the FileHash list of current node
func getNodeFileHash(w http.ResponseWriter, r *http.Request) {
	fileHashList, err := transfer.GetFileHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, fileHashList)
	}
}

// getNodeFileInfoHash returns the FileInfoHash list of current node
func getNodeFileInfoHash(w http.ResponseWriter, r *http.Request) {
	fileInfoHashList, err := transfer.GetFileInfoHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, fileInfoHashList)
	}
}
