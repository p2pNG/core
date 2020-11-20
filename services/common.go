package services

import (
	"encoding/json"
	"net/http"
)

// FileInfo describes how to download a file
type FileInfo struct {
	Size        int64
	Hash        string
	PieceLength int64
	PieceHash   []string
	WellKnown   []string
}

// SeedFileItem describes base info of a file
type SeedFileItem struct {
	Path            string
	Size            int64
	Hash            string
	RecPieceLength  int64
	RecFileInfoHash string
}

// SeedInfo describes a p2pNG seed
type SeedInfo struct {
	Title     string
	Files     []SeedFileItem
	ExtraInfo map[string]string
	WellKnown []string
}

// WriteRespDataAsJson convert data into json format and response to client
func WriteRespDataAsJson(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(&data)
	if err == nil {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
