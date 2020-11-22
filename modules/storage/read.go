package storage

import (
	"errors"
	"os"
)

// ReadFilePiece returns file piece matches the specified pieceIndex
func ReadFilePiece(localFileInfo LocalFileInfo, pieceIndex int64) (piece []byte, err error) {
	offset := localFileInfo.PieceLength * pieceIndex
	if offset >= localFileInfo.Size || offset < 0 {
		return nil, errors.New("invalid pieceIndex")
	}
	file, err := os.Open(localFileInfo.Path)
	if err != nil {
		return nil, err
	}
	piece = make([]byte, localFileInfo.PieceLength)
	_, err = file.ReadAt(piece, offset)
	if err != nil {
		return nil, err
	}
	return piece, nil
}
