package status

import (
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"net/http"
)

// serverGetNodeStatus returns the NodeInfo of current node
func serverGetNodeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value(services.StatusContext).(coreStatusContext)
	node := nodeInfo{
		Name:      utils.GetHostname(),
		Version:   core.GoModule().Version,
		BuildName: ctx.Config.BuildName,
	}
	services.WriteRespDataAsJSON(w, &node)
}

// serverGetNodePeers returns the peers of current node
func serverGetNodePeers(w http.ResponseWriter, r *http.Request) {
	peers, err := discovery.GetPeerRegistry()
	if err == nil {
		services.WriteRespDataAsJSON(w, &peers)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// serverGetNodeSeeds returns the SeedInfo list of current node
func serverGetNodeSeeds(w http.ResponseWriter, r *http.Request) {
	seedHashList, err := getSeedInfoHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, seedHashList)
	}
}

// serverGetNodeFileHash returns the FileHash list of current node
func serverGetNodeFileHash(w http.ResponseWriter, r *http.Request) {
	fileHashList, err := getFileHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, fileHashList)
	}
}

// serverGetNodeFileInfoHash returns the FileInfoHash list of current node
func serverGetNodeFileInfoHash(w http.ResponseWriter, r *http.Request) {
	fileInfoHashList, err := getFileInfoHashList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, fileInfoHashList)
	}
}

// serverGetNodePPList returns the PeerPieceInfo list of current node
func serverGetNodePPList(w http.ResponseWriter, r *http.Request) {
	peerPieceInfoList, err := getPeerPieceInfoList()
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
	} else {
		services.WriteRespDataAsJSON(w, peerPieceInfoList)
	}
}
