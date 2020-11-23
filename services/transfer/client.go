package transfer

import (
	"errors"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"net/http"
)

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
		err = errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
	}
	return
}

// QueryFileInfoByFileHash returns fileInfo that matches fileInfoHash from peerAddr
func QueryFileInfoByFileHash(peerAddr string, fileHash string) (fileInfo []storage.FileInfo, err error) {
	client, err := services.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Get(peerAddr + "/transfer/fileInfo/fileInfoHash/" + fileHash)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		err = services.ReadJSONBody(resp, fileInfo)
	} else {
		err = errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
	}
	return
}

// QuerySeedInfoBySeedInfoHash returns seedInfoHash that matches seedInfoHash from peerAddr
func QuerySeedInfoBySeedInfoHash(peerAddr string, seedInfoHash string) (seedInfo *storage.SeedInfo, err error) {
	client, err := services.GetHttpClient()
	if err != nil {
		return
	}

	resp, err := client.Get(peerAddr + "/transfer/seedInfo/" + seedInfoHash)
	if err != nil {
		return nil, err
	}
	seedInfo = new(storage.SeedInfo)
	if resp.StatusCode == http.StatusOK {
		err = services.ReadJSONBody(resp, seedInfo)
	} else {
		err = errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
	}
	return
}
