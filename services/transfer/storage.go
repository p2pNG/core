package transfer

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	bolt "go.etcd.io/bbolt"
)

// SaveTestData save test data
func SaveTestData() {
	e := saveFileInfo(storage.TestFileInfo)
	if e != nil {
		logging.Log().Warn(e.Error())
	}
	e = saveSeedInfo(storage.TestSeedInfo)
	if e != nil {
		logging.Log().Warn(e.Error())
	}
	e = saveLocalFileInfo(storage.TestFileInfoHash, storage.TestLocalFileInfo)
	if e != nil {
		logging.Log().Warn(e.Error())
	}
	e = savePeerPieceInfo(storage.TestFileInfoHash, storage.TestPeerPieceInfo)
	if e != nil {
		logging.Log().Warn(e.Error())
	}
}

// saveFileInfo save FileInfo to FileInfoHashToFileDB and FileHashToFileDB
func saveFileInfo(fileInfo storage.FileInfo) error {
	db, err := database.GetDBEngine()

	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		// save FileInfoHash
		buk := tx.Bucket([]byte(services.FileInfoHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToFileDB + "] does not exist")
		}
		fileInfoHash := storage.HashFileInfo(fileInfo)
		jsonData, err := json.Marshal(fileInfo)
		if err != nil {
			return err
		}
		err = buk.Put([]byte(fileInfoHash), jsonData)
		if err != nil {
			return err
		}
		// save FileHash
		buk = tx.Bucket([]byte(services.FileHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileHashToFileDB + "] does not exist")
		}
		// get fileInfoHashMap from bucket
		fileInfoHashMapJSON := buk.Get([]byte(fileInfo.Hash))
		fileInfoHashMap := make(map[string]storage.FileInfo)
		if fileInfoHashMapJSON != nil {
			err = json.Unmarshal(fileInfoHashMapJSON, &fileInfoHashMap)
			if err != nil {
				return err
			}
		}
		// overwrite FileInfo if exist,because fileInfoHash won‘t change when well known change
		fileInfoHashMap[fileInfoHash] = fileInfo
		fileInfoHashMapJSON, err = json.Marshal(fileInfoHashMap)
		if err != nil {
			return err
		}
		err = buk.Put([]byte(fileInfo.Hash), fileInfoHashMapJSON)
		if err != nil {
			return err
		}
		return nil
	})
}

// saveSeedInfo save SeedInfo to SeedHashToSeedDB
func saveSeedInfo(seedInfo storage.SeedInfo) error {
	db, err := database.GetDBEngine()

	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.SeedHashToSeedDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.SeedHashToSeedDB + "] does not exist")
		}
		seedInfoHash := storage.HashSeedInfo(seedInfo)
		jsonData, err := json.Marshal(seedInfo)
		if err != nil {
			return err
		}
		return buk.Put([]byte(seedInfoHash), jsonData)
	})
}

// saveLocalFileInfo save LocalFileInfo to FileInfoHashToLocalFileDB
func saveLocalFileInfo(fileInfoHash string, localFileInfo storage.LocalFileInfo) error {
	db, err := database.GetDBEngine()

	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToLocalFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToLocalFileDB + "] does not exist")
		}
		jsonData, err := json.Marshal(localFileInfo)
		if err != nil {
			return err
		}
		return buk.Put([]byte(fileInfoHash), jsonData)
	})
}

func savePeerPieceInfo(fileInfoHash string, ppInfo storage.PeerPieceInfo) (err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToPeerPieceDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToPeerPieceDB + "] does not exist")
		}
		var savedPPInfo = make(storage.PeerPieceInfo)
		resp := buk.Get([]byte(fileInfoHash))
		if resp != nil {
			err = json.Unmarshal(resp, &savedPPInfo)
			if err != nil {
				return err
			}
		}
		for peerAddr, pieceInfo := range ppInfo {
			savedPPInfo[peerAddr] = pieceInfo
		}

		jsonData, err := json.Marshal(savedPPInfo)
		if err != nil {
			return err
		}
		return buk.Put([]byte(fileInfoHash), jsonData)
	})
}

// getSeedInfoBySeedInfoHash returns SeedInfo that matches seedInfoHash
func getSeedInfoBySeedInfoHash(seedInfoHash string) (seed storage.SeedInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.SeedHashToSeedDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.SeedHashToSeedDB + "] does not exist")
		}
		seedJSON := buk.Get([]byte(seedInfoHash))
		if seedJSON == nil {
			return nil
		}
		return json.Unmarshal(seedJSON, &seed)
	})
	return
}

// getFileInfoByFileInfoHash returns FileInfo that matches fileInfoHash
func getFileInfoByFileInfoHash(fileInfoHash string) (file storage.FileInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToFileDB + "] does not exist")
		}
		fileJSON := buk.Get([]byte(fileInfoHash))
		if fileJSON == nil {
			return errors.New("file info not found")
		}
		return json.Unmarshal(fileJSON, &file)
	})
	return
}

// getFileInfoByFileHash returns FileInfo that matches fileHash
func getFileInfoByFileHash(fileHash string) (files []storage.FileInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileHashToFileDB + "] does not exist")
		}
		// get fileInfoHashMap from bucket
		fileInfoHashMapJSON := buk.Get([]byte(fileHash))
		fileInfoHashMap := make(map[string]storage.FileInfo)
		if fileInfoHashMapJSON != nil {
			err = json.Unmarshal(fileInfoHashMapJSON, &fileInfoHashMap)
			if err != nil {
				return err
			}
			for _, v := range fileInfoHashMap {
				files = append(files, v)
			}
		}
		return nil
	})
	return
}

// getLocalFileInfoByFileInfoHash returns LocalFileInfo that matches fileInfoHash
func getLocalFileInfoByFileInfoHash(fileInfoHash string) (localFileInfo storage.LocalFileInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToLocalFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToLocalFileDB + "] does not exist")
		}

		if localFileInfoJSON := buk.Get([]byte(fileInfoHash)); localFileInfoJSON != nil {
			err = json.Unmarshal(localFileInfoJSON, &localFileInfo)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

// getPeerBySeedHash returns peers that has this SeedInfo match the seedHash
func getPeerBySeedHash(seedHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.SeedHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.SeedHashToPeerDB + "] does not exist")
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

// getPeerByFileInfoHash returns peers that has this FileInfo match the fileInfoHash
func getPeerByFileInfoHash(fileInfoHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToPeerDB + "] does not exist")
		}
		peerMapJSON := buk.Get([]byte(fileInfoHash))
		if peerMapJSON == nil {
			return errors.New("peerInfo not fond")
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

// getPeerByFileHash returns peers that has this FileInfo match the fileHash
func getPeerByFileHash(fileHash string) (peers []discovery.PeerInfo, err error) {
	db, err := database.GetDBEngine()

	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileHashToPeerDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileHashToPeerDB + "] does not exist")
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

func getPeerPieceInfoByFileInfoHash(fileInfoHash string) (ppInfo storage.PeerPieceInfo, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return nil, err
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(services.FileInfoHashToPeerPieceDB))
		if buk == nil {
			return errors.New("database error : bucket [" + services.FileInfoHashToPeerPieceDB + "] does not exist")
		}
		ppInfo = make(storage.PeerPieceInfo)
		jsonData := buk.Get([]byte(fileInfoHash))
		return json.Unmarshal(jsonData, &ppInfo)
	})
	return
}
