package database

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/internal/utils"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"path"
	"time"
)

var defaultDBEngine *bolt.DB

// OpenDB open or create a DB engine
func OpenDB(filename string) (err error) {
	dbPath := path.Join(utils.AppDataDir(), filename)
	opts := bolt.DefaultOptions
	opts.Timeout = time.Second * 5
	defaultDBEngine, err = bolt.Open(dbPath, 0644, opts)
	return
}

// GetDBEngine return the default DB Engine; if it is not opened, it will try open
func GetDBEngine() (engine *bolt.DB, err error) {
	if defaultDBEngine == nil {
		err = OpenDB("database")
	}
	engine = defaultDBEngine
	return
}

// InitBuckets is a helper func to create some buckets in the opened DB
func InitBuckets(db *bolt.DB, buk []string) error {
	return db.Update(func(tx *bolt.Tx) error {
		for _, buk := range buk {
			_, err := tx.CreateBucketIfNotExists([]byte(buk))
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// CloseDBEngine close the default DB Engine
func CloseDBEngine() {
	if defaultDBEngine != nil {
		if err := defaultDBEngine.Close(); err != nil {
			logging.Log().Error("close database failed", zap.Error(err))
		}
	}
}
