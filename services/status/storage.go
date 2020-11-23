package status

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	bolt "go.etcd.io/bbolt"
)

// saveSeedInfoHash to save SeedInfoHash list
// input the SeedInfo that the peer have
func saveSeedInfoHash(seedHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.SeedHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.SeedHashToPeerDB + "] does not exist")
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

// saveFileInfoHash to save FileInfoHash list
// input the FileInfo that the peer have
func saveFileInfoHash(fileInfoHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToPeerDB + "] does not exist")
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

// saveFileHash to save FileHash list
// input the FileHash that the peer have
func saveFileHash(fileHashList []string, peer discovery.PeerInfo) (err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileHashToPeerDB + "] does not exist")
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

// getSeedInfoHashList returns all keys from SeedHashToSeedDB
func getSeedInfoHashList() (seedInfoHashList []string, err error) {
	return services.GetAllKeyFromBucket(services.SeedHashToSeedDB)
}

// getFileInfoHashList returns all keys from FileInfoHashToFileDB
func getFileInfoHashList() (fileInfoHashList []string, err error) {
	return services.GetAllKeyFromBucket(services.FileInfoHashToFileDB)
}

// getFileHashList returns all keys from FileHashToFileDB
func getFileHashList() (fileHashList []string, err error) {
	return services.GetAllKeyFromBucket(services.FileHashToFileDB)
}

func getPeerPieceInfoByFileHash() {

}

func getPeerPieceInfoList() {

}
