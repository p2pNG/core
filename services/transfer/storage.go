package transfer

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	bolt "go.etcd.io/bbolt"
)

// GetSeedInfoHashList returns all keys from SeedHashToSeedDB
func GetSeedInfoHashList() (seedInfoHashList []string, err error) {
	return services.GetAllKeyFromBucket(SeedHashToSeedDB)
}

// GetFileInfoHashList returns all keys from FileInfoHashToFileDB
func GetFileInfoHashList() (fileInfoHashList []string, err error) {
	return services.GetAllKeyFromBucket(FileInfoHashToFileDB)
}

// GetFileHashList returns all keys from FileHashToFileDB
func GetFileHashList() (fileHashList []string, err error) {
	return services.GetAllKeyFromBucket(FileHashToFileDB)
}

// GetSeedInfo returns SeedInfo that matches seedInfoHash
func GetSeedInfo(seedInfoHash string) (seed storage.SeedInfo, err error) {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(SeedHashToSeedDB))
		if buk == nil {
			return errors.New("database error : bucket [" + SeedHashToSeedDB + "] does not exist")
		}
		seedJSON := buk.Get([]byte(seedInfoHash))
		if seedJSON == nil {
			return nil
		}
		return json.Unmarshal(seedJSON, &seed)
	})
	return
}

// GetFileInfoByFileInfoHash returns FileInfo that matches fileInfoHash
func GetFileInfoByFileInfoHash(fileInfoHash string) (file storage.FileInfo, err error) {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileInfoHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileInfoHashToFileDB + "] does not exist")
		}
		fileJSON := buk.Get([]byte(fileInfoHash))
		if fileJSON == nil {
			return nil
		}
		return json.Unmarshal(fileJSON, &file)
	})
	return
}

// GetFileInfoByFileHash returns FileInfo that matches fileHash
func GetFileInfoByFileHash(fileHash string) (files []storage.FileInfo, err error) {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(FileHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileHashToFileDB + "] does not exist")
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

// SaveSeedInfo save SeedInfo to SeedHashToSeedDB
func SaveSeedInfo(seed storage.SeedInfo) error {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(SeedHashToSeedDB))
		if buk == nil {
			return errors.New("database error : bucket [" + SeedHashToSeedDB + "] does not exist")
		}
		seedInfoHash := storage.HashSeedInfo(seed)
		jsonData, err := json.Marshal(seed)
		if err != nil {
			return err
		}
		return buk.Put([]byte(seedInfoHash), jsonData)
	})
}

// SaveFileInfo save FileInfo to FileInfoHashToFileDB and FileHashToFileDB
func SaveFileInfo(file storage.FileInfo) error {
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		// save FileInfoHash
		buk := tx.Bucket([]byte(FileInfoHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileInfoHashToFileDB + "] does not exist")
		}
		fileInfoHash := storage.HashFileInfo(file)
		jsonData, err := json.Marshal(file)
		if err != nil {
			return err
		}
		err = buk.Put([]byte(fileInfoHash), jsonData)
		if err != nil {
			return err
		}
		// save FileHash
		buk = tx.Bucket([]byte(FileHashToFileDB))
		if buk == nil {
			return errors.New("database error : bucket [" + FileHashToFileDB + "] does not exist")
		}
		// get fileInfoHashMap from bucket
		fileInfoHashMapJSON := buk.Get([]byte(file.Hash))
		fileInfoHashMap := make(map[string]storage.FileInfo)
		if fileInfoHashMapJSON != nil {
			err = json.Unmarshal(fileInfoHashMapJSON, &fileInfoHashMap)
			if err != nil {
				return err
			}
		}
		// add current FileInfo if not exist
		if _, ok := fileInfoHashMap[fileInfoHash]; !ok {
			fileInfoHashMap[fileInfoHash] = file
		}
		fileInfoHashMapJSON, err = json.Marshal(fileInfoHashMap)
		if err != nil {
			return err
		}
		err = buk.Put([]byte(file.Hash), fileInfoHashMapJSON)
		if err != nil {
			return err
		}
		return nil
	})
}