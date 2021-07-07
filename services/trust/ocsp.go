package trust

import (
	"bytes"
	"crypto/x509"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/request"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/ocsp"
	"io/ioutil"
	"net/http"
	"time"
)

func ocspResponder(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ocspReq, err := ocsp.ParseRequest(reqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p := r.Context().Value(contextTypePlugin).(coreTrustPlugin)

	tpl := queryOCSP(p, ocspReq)

	respData, err := ocsp.CreateResponse(p.caCert, tpl.Certificate, tpl, p.caSigner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(respData)
}

func queryOCSP(ctx coreTrustPlugin, req *ocsp.Request) (resp ocsp.Response) {
	resp = ocsp.Response{}
	if req == nil {
		resp.Status = ocsp.ServerFailed
		return
	}
	db := ctx.authority.GetDatabase()
	cert, err := db.GetCertificate(req.SerialNumber.String())
	if err != nil {
		resp.Status = ocsp.Unknown
		return
	}
	resp.SerialNumber = req.SerialNumber
	resp.Certificate = cert
	resp.ThisUpdate = time.Now()
	rvk, err := db.IsRevoked(req.SerialNumber.String())
	if err != nil {
		resp.Status = ocsp.ServerFailed
		return
	}
	if rvk {
		resp.Status = ocsp.Revoked
		// todo: Revoke Reason and Revoke At
		return
	}
	resp.NextUpdate = time.Now().Add(ctx.ocspDuration)
	resp.Status = ocsp.Good
	return
}

func QueryOCSP() (*ocsp.Response, error) {
	client, err := getClientCert()
	issuer, err := getStoredCert([]byte("my"), []byte("issuer"))
	spew.Dump(err)
	ocspReqContent, err := ocsp.CreateRequest(client, issuer, nil)
	_, tlsClient, err := request.GetDefaultHTTPClient()
	if err != nil {
		return nil, err
	}
	resp, err := tlsClient.Post(issuer.OCSPServer[0], "application/ocsp-request", bytes.NewReader(ocspReqContent))
	if err != nil {
		return nil, err
	}
	ocspResp, err := ioutil.ReadAll(resp.Body)
	return ocsp.ParseResponse(ocspResp, issuer)
}

func getStoredCert(bucket, name []byte) (*x509.Certificate, error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return nil, err
	}
	var certByte []byte
	_ = db.View(func(tx *bolt.Tx) error {
		certBucket := tx.Bucket(append([]byte("certs_"), bucket...))
		certByte = certBucket.Get(name)
		return nil
	})
	if certByte == nil {
		return nil, ErrCertNotFoundInDB
	}
	return x509.ParseCertificate(certByte)
}

func saveCertToStore(bucket, name string, cert *x509.Certificate) (err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return
	}
	return db.Update(func(tx *bolt.Tx) error {
		certBucket := tx.Bucket([]byte("certs_" + bucket))
		return certBucket.Put([]byte(name), cert.Raw)
	})
}

var (
	ErrCertNotFoundInDB = errors.New("certificate not exist")
)

func getClientCert() (*x509.Certificate, error) {
	client, err := getStoredCert([]byte("my"), []byte("client"))

	if err == ErrCertNotFoundInDB {
		client, err = askForClientCert()
		if err != nil {
			return nil, err
		}
		return client, saveCertToStore("my", "client", client)
	}
	return client, err
}
func askForClientCert() (*x509.Certificate, error) {

	return nil, nil
}
