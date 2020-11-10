package trust

import (
	"github.com/p2pNG/core/internal/utils"
	"github.com/smallstep/certificates/db"
	"path"
	"time"

	"github.com/smallstep/certificates/acme"
	acmeAPI "github.com/smallstep/certificates/acme/api"
	"github.com/smallstep/certificates/authority"
	"github.com/smallstep/certificates/authority/provisioner"
	"github.com/smallstep/nosql"
)

// todo: replace smallstep/certificates with p2pNG/certificates
//todo: add new challenge type for common user
func (p *coreTrustPlugin) initAcme() (err error) {
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
		DB: &db.Config{Type: "bbolt", DataSource: path.Join(utils.AppDataDir(), p.config.DbFilename)},
	}

	p.authority, err = authority.NewEmbedded(
		authority.WithConfig(cfg),
		authority.WithX509Signer(p.caCert, p.caSigner),
		authority.WithX509RootCerts(p.rootCert),
	)
	if err != nil {
		return err
	}
	p.acme, err = acme.New(p.authority, acme.AuthorityOptions{
		DB: p.authority.GetDatabase().(nosql.DB), DNS: "localhost"})
	if err != nil {
		return err
	}
	p.acmeHandler = acmeAPI.New(p.acme)
	return nil
}
