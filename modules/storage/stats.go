package storage

// DefaultFileBlockSize is used for the default parameter to split a file to several blocks
//todo: May Declare in elsewhere
const DefaultFileBlockSize = 4 * 1024 * 1024

// StatLocalFile return the LocalFileInfo for a file in the disk
//func StatLocalFile(filepath string, blockSize int64) (lf *LocalFileInfo, err error) {
//	stat, err := os.Stat(filepath)
//	if err != nil {
//		return
//	}
//	if stat.IsDir() {
//		err = errors.New("not a valid file")
//		return
//	}
//	fi, err := StatFile(filepath, blockSize)
//	if err != nil {
//		return
//	}
//
//	lf = new(LocalFileInfo)
//	lf.FileInfo = *fi
//	lf.LastModify = stat.ModTime()
//	lf.Path = filepath
//	return
//}

// StatFile return the FileInfo for a file in the disk
//func StatFile(filepath string, pieceLength int64) (fi *FileInfo, err error) {
//	//todo: calling should be revered, while LocalFileInfo including full os.FileInfo
//	if pieceLength <= 1024*1024 {
//		pieceLength = DefaultFileBlockSize
//	}
//
//	stat, err := os.Stat(filepath)
//	if err != nil {
//		return
//	}
//	if stat.IsDir() {
//		err = errors.New("not a valid file")
//		return
//	}
//	fi = new(FileInfo)
//	fi.PieceLength = pieceLength
//	fi.Hash, fi.Size = stat.Name(), stat.Size()
//
//	f, err := os.Open(filepath)
//	if err != nil {
//		return
//	}
//	buf := make([]byte, pieceLength)
//	fileSum := sha512.New()
//	blockHash := sha512.New512_256()
//	flagTail := false
//	var n int
//
//	for {
//		n, err = f.Read(buf)
//		if err != nil {
//			if err == io.EOF {
//				break
//			} else {
//				return
//			}
//		}
//		if int64(n) != pieceLength {
//			flagTail = true
//			break
//		}
//
//		_, _ = fileSum.Write(buf)
//		blockHash.Reset()
//		_, _ = blockHash.Write(buf)
//		fi.BlockHash = append(fi.BlockHash, blockHash.Sum(nil))
//	}
//	if flagTail {
//		if int64(len(fi.BlockHash))*pieceLength+int64(n) != fi.Size {
//			err = errors.New("read file error, length not matched")
//			return
//		}
//		_, _ = fileSum.Write(buf)
//		blockHash.Reset()
//		_, _ = blockHash.Write(buf)
//		fi.BlockHash = append(fi.BlockHash, blockHash.Sum(nil))
//	}
//	fi.Hash = fileSum.Sum(nil)
//	return fi, nil
//}
