package listener

import (
	"github.com/go-chi/chi"
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
)

func ListenHttp(r chi.Router, addr string) error {
	return http.ListenAndServe(addr, r)
}
func ListenTLS(r chi.Router, addr string) error {
	// todo: Change this
	return http.ListenAndServeTLS(addr, "", "", r)
}
func ListenQUIC(r chi.Router, addr string) error {
	// todo: Change this
	return http3.ListenAndServeQUIC(addr, "", "", r)
}

func Listen(r chi.Router, addr string, tls bool) {
	//todo: Change this
	if tls {
		go func() { _ = ListenTLS(r, addr) }()
		go func() { _ = ListenQUIC(r, addr) }()
	} else {
		go func() { _ = ListenHttp(r, addr) }()
	}
}
