package status

import (
	"encoding/json"
	"errors"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"net/http"
)

// exchangePeers query peers from discovered nodes and save
func exchangePeers() error {
	return visitPeers(func(node discovery.PeerInfo) error {
		client, err := services.GetHttpClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(node.Address.String() + "/status/peer")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var body []byte
			_, err = resp.Body.Read(body)
			if err != nil {
				return err
			}
			var peers []discovery.PeerInfo
			err = json.Unmarshal(body, &peers)
			if err != nil {
				return err
			}
			err = discovery.SavePeers(peers)
			if err != nil {
				return err
			}
		} else {
			return errors.New(resp.Header.Get("p2pNG-msg"))
		}
		return nil
	})
}

// exchangeSeeds query SeedInfoHash list from discovered nodes and save
func exchangeSeeds() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHttpClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(peer.Address.String() + "/status/seed")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var body []byte
			_, err = resp.Body.Read(body)
			if err != nil {
				return err
			}
			var seedHashList []string
			err = json.Unmarshal(body, &seedHashList)
			if err != nil {
				return err
			}
			err = SaveSeedInfoHash(seedHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New(resp.Header.Get("p2pNG-msg"))
		}
		return nil
	})
}

// exchangeFileInfoHash query FileInfoHash list from discovered nodes and save
func exchangeFileInfoHash() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHttpClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(peer.Address.String() + "/status/fileInfoHash")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var body []byte
			_, err = resp.Body.Read(body)
			if err != nil {
				return err
			}
			var fileInfoHashList []string
			err = json.Unmarshal(body, &fileInfoHashList)
			if err != nil {
				return err
			}
			err = SaveFileInfoHash(fileInfoHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New(resp.Header.Get("p2pNG-msg"))
		}
		return nil
	})
}

// exchangeFileHash query FileHash list from discovered nodes and save
func exchangeFileHash() error {
	return visitPeers(func(peer discovery.PeerInfo) error {
		client, err := services.GetHttpClient()
		if err != nil {
			return err
		}
		resp, err := client.Get(peer.Address.String() + "/status/fileHash")
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			var body []byte
			_, err = resp.Body.Read(body)
			if err != nil {
				return err
			}
			var fileHashList []string
			err = json.Unmarshal(body, &fileHashList)
			if err != nil {
				return err
			}
			err = SaveFileInfoHash(fileHashList, peer)
			if err != nil {
				return err
			}
		} else {
			return errors.New(resp.Header.Get("p2pNG-msg"))
		}
		return nil
	})
}

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
