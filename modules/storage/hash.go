package storage

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"os"
	"strconv"
)

// HashSeedInfo returns hash of SeedInfo content
func HashSeedInfo(seedInfo SeedInfo) (seedInfoHash string) {
	var seedInfoContent []byte
	seedInfoContent = append([]byte(seedInfo.Title + separator))
	if seedInfo.ExtraInfo != nil {
		for k, v := range seedInfo.ExtraInfo {
			seedInfoContent = append(seedInfoContent, []byte(k+separator+v+separator)...)
		}
	}
	if seedInfo.Files != nil {
		// append each SeedFileItem
		for _, v := range seedInfo.Files {
			var seedFileItemContent []byte
			seedFileItemContent = append([]byte(v.Hash + separator))
			seedFileItemContent = append(seedFileItemContent, []byte(strconv.Itoa(int(v.Size))+separator)...)
			seedFileItemContent = append(seedFileItemContent, []byte(v.Path+separator)...)
			seedFileItemContent = append(seedFileItemContent, []byte(strconv.Itoa(int(v.RecPieceLength))+separator)...)
			seedFileItemContent = append(seedFileItemContent, []byte(v.RecFileInfoHash+separator)...)
			seedInfoContent = append(seedInfoContent, []byte(string(seedFileItemContent)+separator)...)
		}
	}
	sum := sha512.New().Sum(seedInfoContent)
	return base64.URLEncoding.EncodeToString(sum)
}

// HashFileInfo returns hash of FileInfo content
func HashFileInfo(fileInfo FileInfo) (fileInfoHash string) {
	var fileInfoContent []byte
	fileInfoContent = append(fileInfoContent, []byte(strconv.Itoa(int(fileInfo.Size))+separator)...)
	fileInfoContent = append(fileInfoContent, []byte(fileInfo.Hash+separator)...)
	if fileInfo.PieceHash != nil {
		for _, pieceHash := range fileInfo.PieceHash {
			fileInfoContent = append(fileInfoContent, []byte(pieceHash+separator)...)
		}
	}
	fileInfoContent = append(fileInfoContent, []byte(strconv.Itoa(int(fileInfo.PieceLength))+separator)...)
	sum := sha512.New().Sum(fileInfoContent)
	return base64.URLEncoding.EncodeToString(sum)
}

// HashFileInBytes returns hash of File content
func HashFileInBytes(file []byte) (fileHash string, err error) {
	if file == nil {
		err = errors.New("could not hash a nil file")
		return
	}
	sum := sha512.New().Sum(file)
	return base64.URLEncoding.EncodeToString(sum), nil
}

// HashFile returns file‘s hash by filePath
func HashFile(filePath string) (fileHash string, err error) {
	file, err := os.Open(filePath)
	defer func() {
		err := file.Close()
		if err != nil {
			logging.Log().Error(err.Error())
		}
	}()
	if err != nil {
		return
	}
	if stat, err := os.Stat(filePath); err != nil {
		content := make([]byte, stat.Size())
		_, err = file.Read(content)
		if err != nil {
			return "", nil
		}
		return HashFileInBytes(content)
	}
	return
}

// HashFilePieceInBytes returns hash of file piece content
func HashFilePieceInBytes(piece []byte) (pieceHash string, err error) {
	if piece == nil {
		err = errors.New("could not hash a nil file piece")
		return
	}
	sum := sha512.New512_256().Sum(piece)
	return base64.URLEncoding.EncodeToString(sum), nil
}
