package storage

import (
	"errors"
	"github.com/p2pNG/core/internal/utils"
	"os"
	"path/filepath"
	"time"
)

// FileWriter a struct to provide local file writing operation
type FileWriter struct {
	info            LocalFileInfo
	file            *os.File
	paddingFilePath string
	isCompleted     bool
}

// NewFileWriter returns a FileWriter that provides writing function
// fileInfo is used to generate LocalFileInfo
// filePath is the path you want to write file at
// You should use `defer FileWriter.Clean()` to ensure that resources are released
func NewFileWriter(fileInfo FileInfo, filePath string) (writer *FileWriter, err error) {
	if utils.IsFilePathExist(filePath) {
		return nil, errors.New("file is already exist:" + filePath)
	}

	// specified suffix to a origin filename to identify that it is downloading
	paddingFilePath := filePath + ".downloading"
	file, err := createPaddingFile(paddingFilePath, fileInfo.Size)
	if err != nil {
		return nil, err
	}

	return &FileWriter{
		info: LocalFileInfo{
			FileInfo:   fileInfo,
			Path:       filePath,
			LastModify: time.Now(),
		},
		file:            file,
		paddingFilePath: paddingFilePath,
	}, nil
}

// WritePiece write data to the corresponding position of piece in w.file
// After writing all piece info file,you should use Complete() to sync,close,
// move padding file to destination and get LocalFileInfo.
func (w *FileWriter) WritePiece(data []byte, pieceIndex int) error {
	offset := int64(pieceIndex) * w.info.PieceLength
	offset, err := w.file.Seek(offset, 0)
	if err != nil {
		return err
	}
	n, err := w.file.WriteAt(data, offset)
	if n != len(data) {
		return errors.New("piece writing is not completed")
	}
	return err
}

// Complete returns LocalFileInfo sum from downloaded file
// This function will do the following before return:
// - sync and close w.file
// - move padding file to w.info.Path
// - update file modify time in w.info
func (w *FileWriter) Complete() (localFileInfo *LocalFileInfo, err error) {
	// check if desFile is exist
	if utils.IsFilePathExist(w.info.Path) {
		return nil, errors.New("file is already exist:" + w.info.Path)
	}
	// sync and close file
	err = w.file.Sync()
	if err != nil {
		return nil, err
	}
	utils.CloseFile(w.file)

	// move to desFilePath
	err = os.Rename(w.paddingFilePath, w.info.Path)
	if err != nil {
		return nil, err
	}

	// update modify time
	stat, err := os.Stat(w.info.Path)
	if err != nil {
		return nil, err
	}
	w.info.LastModify = stat.ModTime()
	w.isCompleted = true
	return &w.info, nil
}

// Clean close w.file and remove padding file
// It will do nothing if Clean() or Complete() has been called
func (w *FileWriter) Clean() {
	if !w.isCompleted {
		utils.CloseFile(w.file)
		utils.RemoveFilePath(w.paddingFilePath)
		w.isCompleted = true
	}
}

// createPaddingFile returns a os.O_WRONLY padding file that length = fileSize
// fileWriter should write file data at downloading file path (like a.txt.downloading)
// when download completed,use fileWriter.Complete(desFilePath) to move
// downloading file to original filepath (like a.txt)
func createPaddingFile(paddingFilePath string, fileSize int64) (*os.File, error) {

	err := os.MkdirAll(filepath.Dir(paddingFilePath), os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(paddingFilePath)
	defer utils.CloseFile(file)
	if err != nil {
		return nil, err
	}
	padding := make([]byte, fileSize)
	n, err := file.Write(padding)
	if err != nil || int64(n) != fileSize {
		utils.CloseFile(file)
		utils.RemoveFilePath(paddingFilePath)
		return nil, errors.New("fail to create padding file")
	}
	return os.OpenFile(paddingFilePath, os.O_WRONLY, os.ModePerm)
}
