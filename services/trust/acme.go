package trust

import (
	"crypto/x509"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/certificate"
	"github.com/smallstep/certificates/api"
	"github.com/smallstep/certificates/db"
	"path"
	"time"

	"github.com/smallstep/certificates/acme"
	acmeAPI "github.com/smallstep/certificates/acme/api"
	"github.com/smallstep/certificates/authority"
	"github.com/smallstep/certificates/authority/provisioner"
	"github.com/smallstep/nosql"
)

func InitAcme() (api.RouterHandler, error) {

	cfg := &authority.Config{
		AuthorityConfig: &authority.AuthConfig{

			Provisioners: provisioner.List{&provisioner.ACME{
				Name: "acme",
				Type: provisioner.TypeACME.String(),
				Claims: &provisioner.Claims{
					MinTLSDur:     &provisioner.Duration{Duration: 5 * time.Minute},
					MaxTLSDur:     &provisioner.Duration{Duration: 24 * time.Hour * 365},
					DefaultTLSDur: &provisioner.Duration{Duration: 12 * time.Hour},
				},
			}},
		},
		DB: &db.Config{Type: "bbolt", DataSource: path.Join(utils.AppDataDir(), "ca.db")},
	}
	caCert, err := certificate.GetCert("ca", "Test CA")
	if err != nil {
		return nil, err
	}
	caKey, err := certificate.GetCertKey("ca")
	if err != nil {
		return nil, err
	}
	caSinger, err := x509.ParseECPrivateKey(caKey)
	if err != nil {
		return nil, err
	}
	xAuthority, err := authority.NewEmbedded(
		authority.WithConfig(cfg),
		authority.WithX509Signer(caCert, caSinger),
		authority.WithX509RootCerts(caCert),
	)
	if err != nil {
		return nil, err
	}
	xAcme, err := acme.New(xAuthority, acme.AuthorityOptions{DB: xAuthority.GetDatabase().(nosql.DB), DNS: "localhost"})
	if err != nil {
		return nil, err
	}
	// create the router for the ACME endpoints
	acmeRouterHandler := acmeAPI.New(xAcme)
	return acmeRouterHandler, nil
}
