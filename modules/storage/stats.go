package storage

import (
	"crypto/sha512"
	"errors"
	"io"
	"os"
	"time"
)

// DefaultFileBlockSize is used for the default parameter to split a file to several blocks
//todo: May Declare in elsewhere
const DefaultFileBlockSize = 4 * 1024 * 1024

// StatLocalFile return the LocalFileInfo for a file in the disk
func StatLocalFile(filepath string, blockSize int64) (lf *LocalFileInfo, err error) {
	stat, err := os.Stat(filepath)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = errors.New("not a valid file")
		return
	}
	fi, err := StatFile(filepath, blockSize)
	if err != nil {
		return
	}

	lf = new(LocalFileInfo)
	lf.FileInfo = *fi
	lf.LastModify = stat.ModTime()
	lf.Path = filepath
	return
}

// StatFile return the FileInfo for a file in the disk
func StatFile(filepath string, blockSize int64) (fi *FileInfo, err error) {
	//todo: calling should be revered, while LocalFileInfo including full os.FileInfo
	if blockSize <= 1024*1024 {
		blockSize = DefaultFileBlockSize
	}

	stat, err := os.Stat(filepath)
	if err != nil {
		return
	}
	if stat.IsDir() {
		err = errors.New("not a valid file")
		return
	}
	fi = new(FileInfo)
	fi.BlockSize = blockSize
	fi.Name, fi.Size = stat.Name(), stat.Size()

	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	buf := make([]byte, blockSize)
	fileSum := sha512.New()
	blockHash := sha512.New512_256()
	flagTail := false
	var n int

	for {
		n, err = f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		}
		if int64(n) != blockSize {
			flagTail = true
			break
		}

		_, _ = fileSum.Write(buf)
		blockHash.Reset()
		_, _ = blockHash.Write(buf)
		fi.BlockHash = append(fi.BlockHash, blockHash.Sum(nil))
	}
	if flagTail {
		if int64(len(fi.BlockHash))*blockSize+int64(n) != fi.Size {
			err = errors.New("read file error, length not matched")
			return
		}
		_, _ = fileSum.Write(buf)
		blockHash.Reset()
		_, _ = blockHash.Write(buf)
		fi.BlockHash = append(fi.BlockHash, blockHash.Sum(nil))
	}
	fi.Hash = fileSum.Sum(nil)
	return fi, nil
}

// FileInfo describe a file by the SHA512 checksum for itself and SHA512-256 of every block of it
type FileInfo struct {
	Name string
	Size int64
	Hash []byte

	BlockSize int64
	BlockHash [][]byte
}

// LocalFileInfo extends FileInfo with the filepath and modify time
type LocalFileInfo struct {
	FileInfo
	Path       string
	LastModify time.Time
}

// SeedInfo contains a list of Files and WellKnown Peers
type SeedInfo struct {
	Title     string     //展示给用户的Title
	Files     []FileInfo //包含的文件列表
	WellKnown []string   //已知提供下载的节点地址
	ExtraInfo []string   //附加展示给用户的信息
}
