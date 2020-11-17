package discovery

import (
	"net"
	"testing"
	"time"
)

func TestPeers(t *testing.T) {
	peer := PeerInfo{
		Address:  net.IPv6loopback,
		Port:     18000,
		DNS:      []string{"localhost"},
		LastSeen: time.Now(),
	}
	if err := SavePeers([]PeerInfo{peer}); err != nil {
		t.Error(err)
	}
	if reg, err := GetPeerRegistry(); err != nil || len(reg) == 0 {
		t.Error(err)
	}
	if resp, err := QueryPeer(peer.DNS[0]); err != nil || resp.DNS[0] != peer.DNS[0] {
		t.Error(err)
	}
}
