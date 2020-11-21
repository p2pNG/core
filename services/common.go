package services

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/modules/database"
	bolt "go.etcd.io/bbolt"
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
