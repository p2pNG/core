package transfer

import (
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ProgressStatus float64 data to describes download progress
type ProgressStatus float64

// Progress[i] has the following possible values:
// Unassigned means this piece haven't been assigned to download
// Assigned means this piece has already been assigned to download but not start
// 0-1 means some part of this piece has been downloaded
// Downloaded means this piece has been downloaded completely
const (
	eoPiece   int    = -1
	noAvlPeer string = "noAvlPeer"
	peerAvl   byte   = 1
	//	peerUnAvl byte   = 0
	// Progress[i]
	Unassigned ProgressStatus = -1
	Assigned   ProgressStatus = 0
	Downloaded ProgressStatus = 1
)

// Progress describes downloading progress of each piece of file
type Progress []ProgressStatus

// FileDownloader use to downloading a file to a specified path by fileInfo
// Progress describes DownloadFile progress of each piece of file
// Progress[i] correspond to fileInfo.PieceHash[i]
// Progress[i] has the following possible values :
// Unassigned, Assigned, 0-1 and Downloaded
type FileDownloader struct {
	Progress     Progress
	fileInfo     storage.FileInfo
	fileWriter   *storage.FileWriter
	fileInfoHash string
	ppInfo       storage.PeerPieceInfo
	client       *http.Client
}

// NewFileDownloaderByFileInfoHash query storage.FileInfo by fileInfoHash and invoke NewFileDownloader
func NewFileDownloaderByFileInfoHash(fileInfoHash string, desFilePath string) (*FileDownloader, error) {
	peers, err := getPeerByFileInfoHash(fileInfoHash)
	if err != nil {
		return nil, err
	}
	ch := make(chan *storage.FileInfo)
	go queryFileInfo(fileInfoHash, peers, ch)
	t := time.NewTicker(time.Second * 15)
	select {
	case fileInfo := <-ch:
		return NewFileDownloader(*fileInfo, desFilePath)
	case <-t.C:
		return nil, errors.New("timeout")
	}
}

// NewFileDownloader download file desFilePath to according to fileInfo
// desFilePath is a complete file path like /download/filename.txt
// desFilePath must be not exist or else error will be "file is already exist:"
func NewFileDownloader(fileInfo storage.FileInfo, desFilePath string) (*FileDownloader, error) {

	if fileInfo.WellKnown == nil || len(fileInfo.WellKnown) <= 0 {
		return nil, errors.New("no WellKnown peer to DownloadFile")
	}
	if fileInfo.PieceHash == nil || len(fileInfo.PieceHash) <= 0 {
		return nil, errors.New("no piece hash data to DownloadFile")
	}

	// check if file exist
	if utils.IsFilePathExist(desFilePath) {
		return nil, errors.New("file is already exist:" + desFilePath)
	}

	// init http client
	client, err := services.GetHTTPClient()
	if err != nil {
		return nil, err
	}

	// init progress
	progress := make([]ProgressStatus, len(fileInfo.PieceHash))
	for i := range progress {
		progress[i] = Unassigned
	}

	fileInfoHash := storage.HashFileInfo(fileInfo)

	ppInfo, err := getPeerPieceInfoByFileInfoHash(fileInfoHash)
	if err != nil {
		return nil, err
	}

	writer, err := storage.NewFileWriter(fileInfo, desFilePath)
	if err != nil {
		return nil, err
	}

	return &FileDownloader{
		Progress:     progress,
		fileInfo:     fileInfo,
		fileWriter:   writer,
		fileInfoHash: fileInfoHash,
		ppInfo:       ppInfo,
		client:       client,
	}, nil
}

// DownloadFile downloads file with random peer selection algorithm
func (w *FileDownloader) DownloadFile() (err error) {
	ch := make(chan error)
	t := time.NewTicker(time.Second * 3)
	go w.downloadFile(ch)
	select {
	case err = <-ch:
		break
	case <-t.C:
		err = errors.New("timeout")
		break
	}
	if err != nil {
		w.fileWriter.Clean()
		logging.Log().Error("fail to DownloadFile file：", zap.Error(err))
	}
	return err
}

func (w *FileDownloader) downloadFile(ch chan error) {
	for {
		// refresh ppInfo
		ppInfo, err := getPeerPieceInfoByFileInfoHash(w.fileInfoHash)
		if err != nil {
			break
		}
		w.ppInfo = ppInfo

		// piece select
		peerAddr, pieceIndex := w.selectPeerRandomly()
		if pieceIndex == eoPiece {
			break
		}
		if peerAddr == noAvlPeer {
			w.Progress[pieceIndex] = Unassigned
			logging.Log().Warn("no peer to DownloadFile piece:" + strconv.Itoa(pieceIndex))
			break
		}
		err = w.downloadPiece(peerAddr, pieceIndex)
		if err != nil {
			// retry
			w.Progress[pieceIndex] = Unassigned
		}
	}
	// get LocalFileInfo and save
	_, err := w.fileWriter.Complete()
	if err != nil {
		ch <- nil
	}
	// todo: save LocalFileInfo and provide file
	//return saveLocalFileInfo(w.fileInfoHash, *localFileInfo)
	ch <- nil
}

// selectPeerRandomly returns selected peerIndex and pieceIndex
// if pieceIndex == eoPiece , it means all of pieces haven‘t been Assigned
// if peerIndex == noAvlPeer , it means there is no peer provides piece[pieceIndex]
func (w *FileDownloader) selectPeerRandomly() (peerAddr string, pieceIndex int) {
	// select piece randomly
	var unassignedPiece []int
	for i := 0; i < len(w.Progress); i++ {
		if w.Progress[i] == Unassigned {
			unassignedPiece = append(unassignedPiece, i)
		}
	}
	if unassignedPiece == nil {
		return "", eoPiece
	}
	pieceIndex = unassignedPiece[rand.Intn(len(unassignedPiece))]
	// select peer randomly
	var avlPeer []string
	for k, v := range w.ppInfo {
		if v[pieceIndex] == peerAvl {
			avlPeer = append(avlPeer, k)
		}
	}
	if avlPeer == nil {
		return noAvlPeer, pieceIndex
	}
	peerAddr = avlPeer[rand.Intn(len(avlPeer))]
	// assign piece
	w.Progress[pieceIndex] = Assigned
	return peerAddr, pieceIndex
}

// downloadPiece DownloadFile file piece that matches pieceIndex
// from peer in w.fileInfo.WellKnown which index is peerIndex
func (w *FileDownloader) downloadPiece(peerAddr string, pieceIndex int) error {
	resp, err := w.client.Get(peerAddr + "/transfer/file/fileInfoHash/" + w.fileInfoHash + "/piece/" + strconv.Itoa(int(pieceIndex)) + "/")
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// check piece length
		if int64(len(body)) > w.fileInfo.PieceLength {
			return errors.New("receive content length is over piece length")
		}

		// check piece hash
		hash, err := storage.HashFilePieceInBytes(body)
		if err != nil {
			return err
		}
		if strings.Compare(hash, w.fileInfo.PieceHash[pieceIndex]) != 0 {
			return errors.New("piece hash does not matches")
		}

		// write to file
		err = w.fileWriter.WritePiece(body, pieceIndex)
		if err != nil {
			return err
		}

		// update progress
		w.Progress[pieceIndex] = Downloaded
	} else {
		return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
	}
	return nil
}

func queryFileInfo(fileInfoHash string, peers []discovery.PeerInfo, ch chan *storage.FileInfo) {
	i := 0
	for {
		i = (i + 1) % len(peers)
		peerAddr := services.PeerInfoToStringAddr(peers[i])
		fileInfo, err := QueryFileInfoByFileInfoHash(peerAddr, fileInfoHash)
		if err == nil {
			ch <- fileInfo
		} else {
			logging.Log().Warn("fail to query file info from "+peerAddr, zap.Error(err))
		}
	}
}
