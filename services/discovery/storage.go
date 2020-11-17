package discovery

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

// GetPeerRegistry returns the whole peer registry in the local DB
func GetPeerRegistry() (p []PeerInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return nil, err
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte("discovery_registry"))
		p = make([]PeerInfo, buk.Stats().KeyN)
		i := 0
		return buk.ForEach(func(k, v []byte) error {
			err := json.Unmarshal(v, &p[i])
			if err != nil {
				logging.Log().Warn("broken peer data:", zap.ByteString("content", v))
			}
			i++
			return nil
		})
	})
	return
}

// QueryPeer returns the PeerInfo if name matches
func QueryPeer(name string) (p *PeerInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte("discovery_registry"))
		resp := buk.Get([]byte(name))
		if resp == nil {
			err = errors.New("peer not fond")
		}
		p = new(PeerInfo)
		return json.Unmarshal(resp, p)
	})
	return
}

// SavePeers is used to Add some PeerInfo into the local registry
func SavePeers(peers []PeerInfo) (err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte("discovery_registry"))
		for _, item := range peers {
			data, err := json.Marshal(item)
			if err != nil {
				return err
			}
			if err = buk.Put([]byte(item.DNS[0]), data); err != nil {
				return err
			}
		}
		return nil
	})
}
