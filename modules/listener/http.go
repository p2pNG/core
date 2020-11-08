package listener

import (
	"github.com/go-chi/chi"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/certificate"
	"net/http"
)

func getServerCertAndKey() (cert, key string) {
	_, err := certificate.GetCert("server", utils.GetHostname()+" Server")
	if err != nil {
		return
	}
	cert = certificate.GetCertFilename("server")
	key = certificate.GetCertKeyFilename("server")
	return
}

// ListenBoth bootstrap the router and serve on both udp and tcp
func ListenBoth(r chi.Router, addr string) error {
	cert, key := getServerCertAndKey()
	//todo: dealing these errors
	return http3.ListenAndServe(addr, cert, key, r)
}

// ListenQUIC bootstrap the router and serve on udp
func ListenQUIC(r chi.Router, addr string) error {
	cert, key := getServerCertAndKey()
	//todo: dealing these errors
	return http3.ListenAndServeQUIC(addr, cert, key, r)
}

// ListenTLS bootstrap the router and serve on tcp
func ListenTLS(r chi.Router, addr string) error {
	cert, key := getServerCertAndKey()
	//todo: dealing these errors
	return http.ListenAndServeTLS(addr, cert, key, r)
}
