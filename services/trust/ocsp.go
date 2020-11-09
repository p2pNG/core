package trust

import (
	"bytes"
	"crypto/x509"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/certificate"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/request"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/ocsp"
	"io/ioutil"
	"net/http"
	"time"
)

func ocspResponder(w http.ResponseWriter, r *http.Request) {
	reqData, _ := ioutil.ReadAll(r.Body)
	ocspReq, _ := ocsp.ParseRequest(reqData)
	caCert, _ := certificate.GetCert("ca", utils.GetHostname()+" CA")
	caKey, _ := certificate.GetCertKey("ca")
	key, _ := x509.ParseECPrivateKey(caKey)
	clientCert, _ := getStoredCert([]byte("issued"), ocspReq.SerialNumber.Bytes())
	tpl := ocsp.Response{
		Status:             ocsp.Good,
		SerialNumber:       ocspReq.SerialNumber,
		ProducedAt:         time.Time{},
		ThisUpdate:         time.Time{},
		NextUpdate:         time.Time{},
		RevokedAt:          time.Time{},
		RevocationReason:   0,
		Certificate:        nil,
		TBSResponseData:    nil,
		Signature:          nil,
		SignatureAlgorithm: 0,
		IssuerHash:         0,
		RawResponderName:   nil,
		ResponderKeyHash:   nil,
		Extensions:         nil,
		ExtraExtensions:    nil,
	}
	respData, _ := ocsp.CreateResponse(caCert, clientCert, tpl, key)
	_, _ = w.Write(respData)
}
func QueryOCSP() (*ocsp.Response, error) {
	client, err := getClientCert()
	issuer, err := getStoredCert([]byte("my"), []byte("issuer"))
	spew.Dump(err)
	ocspReqContent, err := ocsp.CreateRequest(client, issuer, nil)
	_, tlsClient, err := request.GetDefaultHTTPClient()
	resp, err := tlsClient.Post(issuer.OCSPServer[0], "application/ocsp-request", bytes.NewReader(ocspReqContent))
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
