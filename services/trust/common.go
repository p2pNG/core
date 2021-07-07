package trust

import (
	"context"
	"crypto"
	"crypto/x509"
	"github.com/go-chi/chi"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/certificate"
	"github.com/smallstep/certificates/acme"
	"github.com/smallstep/certificates/api"
	"github.com/smallstep/certificates/authority"
	"net/http"
	"time"
)

type coreTrustConfig struct {
	IssuerFile   string
	RootFile     string
	DbFilename   string
	OCSPDuration string
}

type coreTrustPlugin struct {
	config       coreTrustConfig
	authority    *authority.Authority
	acme         *acme.Authority
	acmeHandler  api.RouterHandler
	rootCert     *x509.Certificate
	caCert       *x509.Certificate
	caSigner     crypto.Signer
	ocspDuration time.Duration
}

func (p *coreTrustPlugin) Init() (err error) {
	p.ocspDuration, err = time.ParseDuration(p.config.OCSPDuration)
	if err != nil {
		return
	}
	if err = p.initCerts(); err != nil {
		return
	}
	if err = p.initAcme(); err != nil {
		return
	}
	logging.Log().Info("Core Trust Plugin Init OK!")
	return
}

func (p *coreTrustPlugin) initCerts() (err error) {
	// todo: replace with get only
	p.rootCert, err = certificate.GetCert(p.config.RootFile, "Test CA")
	if err != nil {
		return err
	}
	// todo: replace with get only
	p.caCert, err = certificate.GetCert(p.config.IssuerFile, "Test CA")
	if err != nil {
		return err
	}
	// todo: replace with get only
	caKey, err := certificate.GetCertKey(p.config.IssuerFile)
	if err != nil {
		return err
	}
	p.caSigner, err = x509.ParseECPrivateKey(caKey)
	return err
}

func (p *coreTrustPlugin) Config() interface{} {
	return &p.config
}

func (p *coreTrustPlugin) PluginInfo() *core.PluginInfo {
	return &core.PluginInfo{
		Name:    "github.com/p2pNG/core/services/trust",
		Version: "0.0.0",
		Prefix:  "/trust",
		Buckets: []string{},
	}
}

type contextType int

const (
	contextTypePlugin contextType = iota
)

func (p *coreTrustPlugin) GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextTypePlugin, p)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	router.Get("/ocsp", ocspResponder)
	p.acmeHandler.Route(router)
	return router
}

func init() {
	core.RegisterRouterPlugin(&coreTrustPlugin{})
}
