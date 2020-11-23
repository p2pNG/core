package transfer

import (
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type PeerSelection int
type ProgressStatus float64

// RandomSelection select peer randomly
// Progress[i] has the following possible values:
// Unassigned means this piece haven't been assigned to download
// Assigned means this piece has already been assigned to download but not start
// 0-1 means some part of this piece has been downloaded
// Downloaded means this piece has been downloaded completely
const (
	RandomSelection PeerSelection = iota
	eoPiece         int           = -1
	noAvlPeer       int           = -1
	peerAvl         byte          = 1
	peerUnAvl       byte          = 0
	// Progress[i]
	Unassigned ProgressStatus = -1
	Assigned   ProgressStatus = 0
	Downloaded ProgressStatus = 1
)

// FileDownloader use to downloading a file to a specified path by fileInfo
// Progress describes download progress of each piece of file
// Progress[i] correspond to fileInfo.PieceHash[i]
// Progress[i] has the following possible values :
// Unassigned, Assigned, 0-1 and Downloaded
type FileDownloader struct {
	Progress     []ProgressStatus
	fileInfo     storage.FileInfo
	fileWriter   *storage.FileWriter
	fileInfoHash string
	pieceMap     [][]byte
	client       *http.Client
}

func NewFileDownloader(fileInfo storage.FileInfo, desFilePath string) (*FileDownloader, error) {

	if fileInfo.WellKnown == nil || len(fileInfo.WellKnown) <= 0 {
		return nil, errors.New("no WellKnown peer to download")
	}
	if fileInfo.PieceHash == nil || len(fileInfo.PieceHash) <= 0 {
		return nil, errors.New("no piece hash data to download")
	}

	// check if file exist
	if utils.IsFilePathExist(desFilePath) {
		return nil, errors.New("file is already exist:" + desFilePath)
	}

	writer, err := storage.NewFileWriter(fileInfo, desFilePath)
	if err != nil {
		return nil, err
	}

	// init http client
	client, err := services.GetHttpClient()
	if err != nil {
		return nil, err
	}

	// init progress
	progress := make([]ProgressStatus, len(fileInfo.PieceHash))
	for i, _ := range progress {
		progress[i] = Unassigned
	}

	return &FileDownloader{
		Progress:     progress,
		fileInfo:     fileInfo,
		fileWriter:   writer,
		fileInfoHash: storage.HashFileInfo(fileInfo),
		pieceMap:     nil,
		client:       client,
	}, nil
}

func (w *FileDownloader) createPieceMap() {
	pieceMap := make([][]byte, len(w.fileInfo.WellKnown))
	for i, _ := range pieceMap {
		pieceMap[i] = make([]byte, len(w.fileInfo.PieceHash))
		// todo: request piece map
	}
	w.pieceMap = pieceMap
}

// DownloadFile download file with random peer selection algorithm
func (w *FileDownloader) DownloadFile() error {
	err := w.downloadFile()
	if err != nil {
		w.fileWriter.Clean()
		logging.Log().Error("fail to download file：", zap.Error(err))
	}
	return nil
}

func (w *FileDownloader) downloadFile() error {
	for {
		w.createPieceMap()
		peerIndex, pieceIndex := w.selectPeerRandomly()
		if pieceIndex == eoPiece {
			break
		}
		if peerIndex == noAvlPeer {
			w.Progress[pieceIndex] = Unassigned
			logging.Log().Warn("no peer to download piece:" + strconv.Itoa(pieceIndex))
			break
		}
		err := w.DownloadPiece(peerIndex, pieceIndex)
		if err != nil {
			// retry
			w.Progress[pieceIndex] = Unassigned
		}
	}
	// get LocalFileInfo and save
	_, err := w.fileWriter.Complete()
	if err != nil {
		return err
	}
	// todo: save LocalFileInfo and provide file
	//return SaveLocalFileInfo(w.fileInfoHash, *localFileInfo)
	return nil
}

// selectPeerRandomly returns selected peerIndex and pieceIndex
// if pieceIndex == eoPiece , it means all of pieces haven‘t been Assigned
// if peerIndex == noAvlPeer , it means there is no peer provides piece[pieceIndex]
func (w *FileDownloader) selectPeerRandomly() (peerIndex int, pieceIndex int) {
	// select piece randomly
	var unassignedPiece []int
	for i := 0; i < len(w.Progress); i++ {
		if w.Progress[i] == Unassigned {
			unassignedPiece = append(unassignedPiece, i)
		}
	}
	if unassignedPiece == nil {
		return 0, eoPiece
	}
	pieceIndex = unassignedPiece[rand.Intn(len(unassignedPiece))]
	// select peer randomly
	var avlPeer []int
	for i := 0; i < len(w.pieceMap); i++ {
		if w.pieceMap[i][pieceIndex] == peerAvl {
			avlPeer = append(avlPeer, i)
		}
	}
	if avlPeer == nil {
		return noAvlPeer, pieceIndex
	}
	// assign piece
	w.Progress[pieceIndex] = Assigned
	return peerIndex, pieceIndex
}

// DownloadPiece download file piece that matches pieceIndex
// from peer in w.fileInfo.WellKnown which index is peerIndex
func (w *FileDownloader) DownloadPiece(peerIndex int, pieceIndex int) error {
	resp, err := w.client.Get(w.fileInfo.WellKnown[peerIndex] + "/transfer/file/fileInfoHash/" + w.fileInfoHash + "/piece/" + strconv.Itoa(int(pieceIndex)) + "/")
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

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

// DownloadSeed download all files in seed
func DownloadSeed(seedInfo storage.SeedInfo, seedPath string) error {

	if utils.IsFilePathExist(seedPath) {
		return errors.New("seed file is already exist:" + seedPath)
	} else {
		err := os.MkdirAll(seedPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := downloadSeed(seedInfo, seedPath, sequencePeerSelector); err != nil {
		utils.RemoveFilePath(seedPath)
	}

	// todo: provide seed
	return nil
}

// QueryFileInfoByFileInfoHash returns fileInfo that matches fileInfoHash from peerAddr
func QueryFileInfoByFileInfoHash(peerAddr string, fileInfoHash string) (fileInfo *storage.FileInfo, err error) {
	client, err := services.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Get(peerAddr + "/transfer/fileInfo/fileInfoHash/" + fileInfoHash)
	if err != nil {
		return nil, err
	}
	fileInfo = new(storage.FileInfo)
	if resp.StatusCode == http.StatusOK {
		err = services.ReadJSONBody(resp, fileInfo)
	} else {
		err = errors.New(resp.Header.Get(services.P2PNGMsg))
	}
	return
}

// randomPeerSelector select peer from peers randomly
func randomPeerSelector(peers []string) func() (peer string) {
	return func() (peer string) {
		i := rand.Intn(len(peers))
		return peers[i]
	}
}

// sequencePeerSelector select peer from peers in sequence
func sequencePeerSelector(peers []string) func() (peer string) {
	i := -1
	return func() (peer string) {
		i = (i + 1) % len(peers)
		return peers[i]
	}
}

func downloadSeed(seedInfo storage.SeedInfo, seedPath string, selector func(peers []string) func() (peer string)) error {
	selectPeer := selector(seedInfo.WellKnown)
	for _, item := range seedInfo.Files {
		fileInfo, err := QueryFileInfoByFileInfoHash(selectPeer(), item.RecFileInfoHash)
		if err != nil {
			return err
		}
		downloader, err := NewFileDownloader(*fileInfo, seedPath+item.Path)
		if err != nil {
			return err
		}
		err = downloader.DownloadFile()
		if err != nil {
			return err
		}
	}
	return nil
}
