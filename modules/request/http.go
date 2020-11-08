package request

import (
	"context"
	"crypto/tls"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/certificate"
	"net/http"
	"time"
)

// GetDefaultHTTPClient returns http clients over quic and tls
func GetDefaultHTTPClient() (quicClient, tlsClient *http.Client, err error) {
	if defaultHTTPClientQUIC == nil || defaultHTTPClientTLS == nil {
		err = createDefaultHTTPClient()
	}
	quicClient = defaultHTTPClientQUIC
	tlsClient = defaultHTTPClientTLS
	return
}

var defaultHTTPClientTLS *http.Client
var defaultHTTPClientQUIC *http.Client

func createDefaultHTTPClient() (err error) {
	_, err = certificate.GetCert("client", utils.GetHostname()+" Client")
	if err != nil {
		return
	}
	cert, err := tls.LoadX509KeyPair(certificate.GetCertFilename("client"), certificate.GetCertKeyFilename("client"))
	if err != nil {
		return
	}
	tlsCfg := &tls.Config{
		//todo: Do not skip verifies
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
	defaultHTTPClientQUIC = &http.Client{Transport: &http3.RoundTripper{TLSClientConfig: tlsCfg}}
	defaultHTTPClientTLS = &http.Client{Transport: &http.Transport{TLSClientConfig: tlsCfg}}
	return
}

// HTTPType describe the the transport layer of a http connection
type HTTPType int

const (
	// HTTPTypeQUIC describe a http connection using QUIC (over udp)
	HTTPTypeQUIC HTTPType = 1 << (iota)
	// HTTPTypeTLS describe a http connection using TLS (over tcp)
	HTTPTypeTLS
)

// TestHostAvailable returns which connection type is available or faster
// Notice: 0 indicates this host cannot be connected in any way
func TestHostAvailable(address string) HTTPType {
	//TODO: Should Optimize
	status := make(chan HTTPType)
	udpClient, tcpClient, err := GetDefaultHTTPClient()
	if err != nil {
		return 0
	}
	udpCtx, udpCancel := context.WithTimeout(context.Background(), time.Second*10)
	tcpCtx, tcpCancel := context.WithTimeout(context.Background(), time.Second*10)
	go func() {
		req, err := http.NewRequestWithContext(udpCtx, "HEAD", "https://"+address, nil)
		if err == nil {
			if _, err = udpClient.Do(req); err == nil {
				status <- HTTPTypeQUIC
				return
			}
		}
		status <- 0
	}()
	go func() {
		req, err := http.NewRequestWithContext(tcpCtx, "HEAD", "https://"+address, nil)
		if err == nil {
			if _, err = tcpClient.Do(req); err == nil {
				status <- HTTPTypeTLS
				return
			}
		}
		status <- 0
	}()

	failed := 0
	rt := HTTPType(0)
	select {
	case rt = <-status:
		if rt != 0 {
			break
		}
		failed++
		if failed == 2 {
			break
		}
	}
	tcpCancel()
	udpCancel()
	return rt
}

//todo: Create QUIC request client
