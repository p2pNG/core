package services

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"net/http"
)

type contextType int

// StatusContext request context for status service
// TransferContext request context for transfer service
// P2PNGCode a custom header to describe error code
// P2PNGMsg a custom header to describe error msg
// NoPermissions lack of permissions
const (
	// request context
	StatusContext   contextType = iota
	TransferContext contextType = iota
	// custom header
	P2PNGCode string = "p2pNG-code"
	P2PNGMsg  string = "p2pNG-msg"
	// error code
	NoPermissions string = "3001"
)

// WriteRespDataAsJSON convert data into json format and response to client
func WriteRespDataAsJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(&data)
	if err != nil {
		WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		logging.Log().Error("fail response", zap.Error(err))
	}
}

// WriteRespDataAsOctetStream write response data with Content-Type = application/octet-stream header
func WriteRespDataAsOctetStream(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		logging.Log().Error("fail response", zap.Error(err))
	}
}

// WriteErrorToResp response error msg to client
func WriteErrorToResp(w http.ResponseWriter, err error, statusCode int) {
	logging.Log().Warn("fail request:", zap.Error(err))
	w.Header().Set(P2PNGMsg, err.Error())
	w.WriteHeader(statusCode)
}

// GetAllKeyFromBucket get all keys from specified bucket
func GetAllKeyFromBucket(bucket string) (keys []string, err error) {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return nil, err
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(bucket))
		if buk == nil {
			return errors.New("database error : bucket [" + bucket + "] does not exist")
		}
		err = buk.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
		return err
	})
	return
}
