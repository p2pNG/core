package status

import (
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"net/http"
)

// visitPeers visit each peer and run fn function
func visitPeers(fn func(peer discovery.PeerInfo) error) error {
	peers, err := discovery.GetPeerRegistry()
	if err != nil {
		return err
	}

	for _, peer := range peers {
		err = fn(peer)
		if err != nil {
			logging.Log().Warn(err.Error())
		}
	}
	return nil
}

// exchangePeers query peers from discovered nodes and save
func exchangePeers() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHTTPClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(services.PeerInfoToStringAddr(peer) + "/status/peer")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var peers []discovery.PeerInfo
			err = services.ReadJSONBody(resp, &peers)
			if err != nil {
				return err
			}
			err = discovery.SavePeers(peers)
			if err != nil {
				return err
			}
		} else {
			return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
		}
		return nil
	})
}

// exchangeSeeds query SeedInfoHash list from discovered nodes and save
func exchangeSeeds() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHTTPClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(services.PeerInfoToStringAddr(peer) + "/status/seed")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var seedHashList []string
			err = services.ReadJSONBody(resp, &seedHashList)
			if err != nil {
				return err
			}
			err = saveSeedInfoHash(seedHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
		}
		return nil
	})
}

// exchangeFileInfoHash query FileInfoHash list from discovered nodes and save
func exchangeFileInfoHash() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHTTPClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(services.PeerInfoToStringAddr(peer) + "/status/fileInfoHash")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {

			var fileInfoHashList []string
			err = services.ReadJSONBody(resp, &fileInfoHashList)
			if err != nil {
				return err
			}
			err = saveFileInfoHash(fileInfoHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
		}
		return nil
	})
}

// exchangeFileHash query FileHash list from discovered nodes and save
func exchangeFileHash() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHTTPClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(services.PeerInfoToStringAddr(peer) + "/status/fileHash")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var fileHashList []string
			err := services.ReadJSONBody(resp, &fileHashList)
			if err != nil {
				return err
			}
			err = saveFileInfoHash(fileHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
		}
		return nil
	})
}

// exchangePeerPieceInfo query PeerPieceInfo list from discovered nodes and save
func exchangePeerPieceInfo() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHTTPClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(services.PeerInfoToStringAddr(peer) + "/status/peerPieceInfo")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var ppInfoList map[string]storage.PeerPieceInfo
			err = services.ReadJSONBody(resp, &ppInfoList)
			if err != nil {
				return err
			}
			err = savePeerPieceInfoList(ppInfoList)
			if err != nil {
				return err
			}
		} else {
			return errors.New("peer response error:" + resp.Header.Get(services.P2PNGMsg))
		}
		return nil
	})
}
