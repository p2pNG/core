package storage

import (
	"errors"
	"io"
	"os"
	"time"
)

// StatLocalFile return the LocalFileInfo for a file in the disk
func StatLocalFile(filepath string, pieceLength int64, wellKnown []string) (lf *LocalFileInfo, err error) {
	// todo: merge StatFile function
	stat, err := os.Stat(filepath)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = errors.New("not a valid file")
		return
	}
	fileInfo, err := statFile(filepath, pieceLength, wellKnown)
	if err != nil {
		return nil, err
	}
	modTime, err := time.Parse(TimeLayout, stat.ModTime().Format(TimeLayout))
	if err != nil {
		return nil, err
	}
	return &LocalFileInfo{
		LastModify: modTime,
		Path:       filepath,
		FileInfo:   *fileInfo,
	}, nil
}

// statFile return the FileInfo for a file in the disk
func statFile(filepath string, pieceLength int64, wellKnown []string) (fileInfo *FileInfo, err error) {

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

	var fileBuf []byte
	var pieceHashList []string
	lastPieceLength := 0
	for {
		pieceBuf := make([]byte, pieceLength)
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
			pieceBuf = pieceBuf[:lastPieceLength]
		}

		fileBuf = append(fileBuf, pieceBuf...)

		if pieceHash, err := HashFilePieceInBytes(pieceBuf); err == nil {
			pieceHashList = append(pieceHashList, pieceHash)
		} else {
			return nil, err
		}
	}

	// check file length
	if stat.Size() != int64(len(pieceHashList)-1)*pieceLength+int64(lastPieceLength) {
		err = errors.New("read file error, length not matched")
		return nil, err
	}
	fileHash, err := HashFileInBytes(fileBuf)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Size:        stat.Size(),
		Hash:        fileHash,
		PieceLength: pieceLength,
		PieceHash:   pieceHashList,
		WellKnown:   wellKnown,
	}, nil
}
