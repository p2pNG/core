package services

import (
	"encoding/json"
	"net/http"
)

type contextType int

const (
	StatusContext   contextType = iota
	TransferContext contextType = iota
)

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
