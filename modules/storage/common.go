package storage

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	separator = ":"
)

// FileInfo describes how to download a file
type FileInfo struct {
	Size        int64
	Hash        string
	PieceLength int64
	PieceHash   []string
	WellKnown   []string
}

// SeedFileItem describes base info of a file
type SeedFileItem struct {
	Path            string
	Size            int64
	Hash            string
	RecPieceLength  int64
	RecFileInfoHash string
}

// SeedInfo describes a p2pNG seed
type SeedInfo struct {
	Title     string
	Files     []SeedFileItem
	ExtraInfo map[string]string
	WellKnown []string
}

// LocalFileInfo describes a local file’ path and last modified time
type LocalFileInfo struct {
	Path       string
	LastModify time.Time
}

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
	return base64.StdEncoding.EncodeToString(sha512.New().Sum(seedInfoContent))
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
	return base64.StdEncoding.EncodeToString(sha512.New().Sum(fileInfoContent))
}

// HashFile returns hash of File content
func HashFile(file *os.File) (fileHash string, err error) {
	if file == nil {
		return "", errors.New("could not hash a nil file")
	}
	hash := sha512.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum), nil
}
