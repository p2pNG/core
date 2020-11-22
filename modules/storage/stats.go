package storage

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

// StatLocalFile return the LocalFileInfo for a file in the disk
func StatLocalFile(filepath string, pieceLength int64) (lf *LocalFileInfo, err error) {
	// todo: merge StatFile function
	stat, err := os.Stat(filepath)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = errors.New("not a valid file")
		return
	}
	fileInfo, err := StatFile(filepath, pieceLength)
	if err != nil {
		return nil, err
	}
	return &LocalFileInfo{
		LastModify: stat.ModTime(),
		Path:       filepath,
		FileInfo:   *fileInfo,
	}, nil
}

// StatFile return the FileInfo for a file in the disk
func StatFile(filepath string, pieceLength int64) (fileInfo *FileInfo, err error) {

	if pieceLength <= MinFilePieceLength {
		pieceLength = DefaultFilePieceLength
	}

	stat, err := os.Stat(filepath)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = errors.New("not a valid file")
		return
	}
	file, err := os.Open(filepath)
	if err != nil {
		return
	}

	pieceBuf := make([]byte, pieceLength)
	fileHash := sha512.New()
	var pieceHashList []string
	lastPieceLength := 0
	for {
		n, err := file.Read(pieceBuf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if int64(n) != pieceLength {
			lastPieceLength = n
		}

		_, err = fileHash.Write(pieceBuf)
		if err != nil {
			return nil, err
		}

		if pieceHash, err := HashFilePieceInBytes(pieceBuf); err != nil {
			pieceHashList = append(pieceHashList, pieceHash)
		} else {
			return nil, err
		}
	}

	// check file length
	if stat.Size() != int64(len(pieceHashList))*pieceLength+int64(lastPieceLength) {
		err = errors.New("read file error, length not matched")
		return
	}

	return &FileInfo{
		Size:        stat.Size(),
		Hash:        base64.URLEncoding.EncodeToString(fileHash.Sum(nil)),
		PieceLength: pieceLength,
		PieceHash:   pieceHashList,
	}, nil
}
