package database

import (
	"github.com/p2pNG/core/internal/utils"
	bolt "go.etcd.io/bbolt"
	"os"
	"path"
	"time"
)

var defaultDBEngine *bolt.DB

func openDB() (err error) {
	dbPath := path.Join(utils.AppDataDir(), "database")
	err = os.MkdirAll(utils.AppDataDir(), 0755)
	if err != nil {
		return
	}
	opts := bolt.DefaultOptions
	opts.Timeout = time.Second * 5
	defaultDBEngine, err = bolt.Open(dbPath, 0644, opts)
	return
}

// GetDBEngine return the default DB Engine; if it is not opened, it will try open
func GetDBEngine() (engine *bolt.DB, err error) {
	if defaultDBEngine == nil {
		err = openDB()
	}
	engine = defaultDBEngine
	return
}

// CloseDBEngine close the default DB Engine
func CloseDBEngine() {
	if defaultDBEngine != nil {
		_ = defaultDBEngine.Close()
	}
}
