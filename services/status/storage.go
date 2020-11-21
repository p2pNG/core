package status

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/services/discovery"
	bolt "go.etcd.io/bbolt"
)

// SaveSeedInfoHash to save SeedInfoHash list
// input the SeedInfo that the peer have
func SaveSeedInfoHash(seedHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(SeedHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + SeedHashToPeerDB + "] does not exist")
		}
		for _, seedHash := range seedHashList {
			// add to peer list if not exist
			var peers = make(map[string]discovery.PeerInfo)
			resp := buk.Get([]byte(seedHash))
			if resp != nil {
				err = json.Unmarshal(resp, &peers)
				if err != nil {
					return err
				}
			}

			peerKey := getPeerKey(peer)
			if _, ok := peers[peerKey]; !ok {
				peers[peerKey] = peer
				jsonData, err := json.Marshal(peers)
				if err != nil {
					return err
				}
				return buk.Put([]byte(seedHash), jsonData)
			}
		}
		return nil
	})
}

// GetPeerBySeedHash returns peers that has this SeedInfo match the seedHash
func GetPeerBySeedHash(seedHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(SeedHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + SeedHashToPeerDB + "] does not exist")
		}
		peerMapJSON := buk.Get([]byte(seedHash))
		if peerMapJSON == nil {
			err = errors.New("peerInfo not fond")
		}
		var peerMap map[string]discovery.PeerInfo
		err := json.Unmarshal(peerMapJSON, &peerMap)
		if err != nil {
			return err
		}
		for _, peerInfo := range peerMap {
			peers = append(peers, peerInfo)
		}
		return nil
	})
	return
}

// SaveFileInfoHash to save FileInfoHash list
// input the FileInfo that the peer have
func SaveFileInfoHash(fileInfoHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileInfoHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileInfoHashToPeerDB + "] does not exist")
		}
		for _, fileInfoHash := range fileInfoHashList {
			// add to peer list if not exist
			var peers = make(map[string]discovery.PeerInfo)
			resp := buk.Get([]byte(fileInfoHash))
			if resp != nil {
				err = json.Unmarshal(resp, &peers)
				if err != nil {
					return err
				}
			}

			peerKey := getPeerKey(peer)
			if _, ok := peers[peerKey]; !ok {
				peers[peerKey] = peer
				jsonData, err := json.Marshal(peers)
				if err != nil {
					return err
				}
				return buk.Put([]byte(fileInfoHash), jsonData)
			}
		}
		return nil
	})
}

// GetPeerByFileInfoHash returns peers that has this FileInfo match the fileInfoHash
func GetPeerByFileInfoHash(fileInfoHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileInfoHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileInfoHashToPeerDB + "] does not exist")
		}
		peerMapJSON := buk.Get([]byte(fileInfoHash))
		if peerMapJSON == nil {
			err = errors.New("peerInfo not fond")
		}
		var peerMap map[string]discovery.PeerInfo
		err := json.Unmarshal(peerMapJSON, &peerMap)
		if err != nil {
			return err
		}
		for _, peerInfo := range peerMap {
			peers = append(peers, peerInfo)
		}
		return nil
	})
	return
}

// SaveFileHash to save FileHash list
// input the FileHash that the peer have
func SaveFileHash(fileHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileHashToPeerDB + "] does not exist")
		}
		for _, fileHash := range fileHashList {
			// add to peer list if not exist
			var peers = make(map[string]discovery.PeerInfo)
			resp := buk.Get([]byte(fileHash))
			if resp != nil {
				err = json.Unmarshal(resp, &peers)
				if err != nil {
					return err
				}
			}

			peerKey := getPeerKey(peer)
			if _, ok := peers[peerKey]; !ok {
				peers[peerKey] = peer
				jsonData, err := json.Marshal(peers)
				if err != nil {
					return err
				}
				return buk.Put([]byte(fileHash), jsonData)
			}
		}
		return nil
	})
}

// GetPeerByFileHash returns peers that has this FileInfo match the fileHash
func GetPeerByFileHash(fileHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileHashToPeerDB + "] does not exist")
		}
		peerMapJSON := buk.Get([]byte(fileHash))
		if peerMapJSON == nil {
			err = errors.New("peerInfo not fond")
		}
		var peerMap map[string]discovery.PeerInfo
		err := json.Unmarshal(peerMapJSON, &peerMap)
		if err != nil {
			return err
		}
		for _, peerInfo := range peerMap {
			peers = append(peers, peerInfo)
		}
		return nil
	})
	return
}
