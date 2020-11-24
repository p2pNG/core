package discovery

import (
	"github.com/p2pNG/mdns"
	"os"
	"time"
)

// LocalBroadcast listen and provide mDNS service for local devices
// Attention: Should not Listen in the same goroutine with http service
// todo: Use Structured Info
func LocalBroadcast(port int) (*mdns.Server, error) {
	host, _ := os.Hostname()
	// todo: change to dns list
	info := []string{"test.local"}
	service, _ := mdns.NewMDNSService(host, "_p2pNG._https", "", "", port, nil, info)
	return mdns.NewServer(&mdns.Config{Zone: service})
}

// LocalScan send multicast udp packet and wait for response to discovery local peers
func LocalScan() (rt []PeerInfo, err error) {
	entriesCh := make(chan *mdns.ServiceEntry)
	go func() {
		for x := range entriesCh {
			rt = append(rt, PeerInfo{
				Address:  x.Addr,
				Port:     x.Port,
				LastSeen: time.Now(),
				DNS:      x.InfoFields,
			})
		}
	}()
	err = mdns.Lookup("_p2pNG._https", entriesCh)
	close(entriesCh)
	return
}
