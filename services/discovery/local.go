package discovery

import (
	"github.com/p2pNG/mdns"
	"os"
)

// LocalBroadcast listen and provide mDNS service for local devices
// Attention: Should not Listen in the same goroutine with http service
// todo: Use Structured Info, convey the https service port
func LocalBroadcast(port int) (*mdns.Server, error) {
	host, _ := os.Hostname()
	info := []string{"p2pNG Server Core"}
	service, _ := mdns.NewMDNSService(host, "_p2pNG._https", "", "", port, nil, info)
	return mdns.NewServer(&mdns.Config{Zone: service})
}

// LocalScan send multicast udp packet and wait for response to discovery local peers
func LocalScan() ([]mdns.ServiceEntry, error) {
	var rt []mdns.ServiceEntry
	//todo: Optimize
	entriesCh := make(chan *mdns.ServiceEntry, 64)
	go func() {
		for x := range entriesCh {
			rt = append(rt, *x)
		}
	}()
	err := mdns.Lookup("_p2pNG._https", entriesCh)
	close(entriesCh)
	return rt, err
}
