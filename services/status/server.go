package status

import (
	"encoding/json"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/utils"
	"net/http"
)

func getNodeInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("statusCtx").(coreStatusContext)
	node := NodeInfo{
		Name:      utils.GetHostname(),
		Version:   core.GoModule().Version,
		BuildName: ctx.Config.BuildName,
	}
	data, err := json.Marshal(&node)
	if err == nil {
		w.Header().Set("content-type", "")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
