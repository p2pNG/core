package transfer

import (
	"errors"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/storage"
	"math/rand"
	"os"
)

// todo: use SeedDownloader to provide progress info

// DownloadSeed DownloadFile all files in seed
func DownloadSeed(seedInfoHash string, seedPath string) error {

	// todo : peer selection
	//peers, err := getPeerBySeedHash(seedInfoHash)
	//if err != nil {
	//	return err
	//}

	peer := storage.TestPeerAddr

	seedInfo, err := QuerySeedInfoBySeedInfoHash(peer, seedInfoHash)
	if err != nil {
		return err
	}

	if utils.IsFilePathExist(seedPath) {
		return errors.New("seed file is already exist:" + seedPath)
	} else {
		err := os.MkdirAll(seedPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := downloadSeed(seedInfo, seedPath, sequencePeerSelector); err != nil {
		// delete files and remove path
		utils.RemoveFilePath(seedPath)
	}

	// todo: provide seed
	return nil
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

func downloadSeed(seedInfo *storage.SeedInfo, seedPath string, selector func(peers []string) func() (peer string)) error {
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
