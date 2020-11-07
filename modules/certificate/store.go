package certificate

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"github.com/p2pNG/core/internal/utils"
	wrapCert "github.com/p2pNG/core/modules/certificate/wrap"
	"io/ioutil"
	rand2 "math/rand"
	"os"
	"path"
	"software.sslmate.com/src/go-pkcs12"
)

// GetCertFilename returns the file path of Certificate in the AppDataDir
func GetCertFilename(name string) string {
	basePath := path.Join(utils.AppDataDir(), "certificate")
	_ = os.MkdirAll(basePath, 0755)
	return path.Join(basePath, name+".cer")
}

// GetCertKeyFilename returns the file path of private key in the AppDataDir
func GetCertKeyFilename(name string) string {
	basePath := path.Join(utils.AppDataDir(), "certificate")
	_ = os.MkdirAll(basePath, 0755)
	return path.Join(basePath, name+".key")
}

// GetCertBundleFilename returns the file path of Certificate Bundle in the AppDataDir
func GetCertBundleFilename(name string) string {
	basePath := path.Join(utils.AppDataDir(), "certificate")
	_ = os.MkdirAll(basePath, 0755)
	return path.Join(basePath, name+".pfx")
}

// GetCertBundle return or create a Certificate Bundle, based on GetCert and GetCertKey.
// Especially for client certificate, user may need to import to system.
func GetCertBundle(name string, subject string) ([]byte, error) {
	certFile := GetCertBundleFilename(name)
	_, err := os.Stat(certFile)
	if os.IsNotExist(err) {
		return createCertBundle(name, subject)
	} else {
		return ioutil.ReadFile(certFile)
	}
}
func createCertBundle(name string, subject string) ([]byte, error) {
	privDer, err := GetCertKey(name)
	if err != nil {
		return nil, err
	}

	cert, err := GetCert(name, subject)
	if err != nil {
		return nil, err
	}
	priv, err := x509.ParseECPrivateKey(privDer)
	if err != nil {
		return nil, err
	}
	pfx, err := pkcs12.Encode(rand.Reader, priv, cert, nil, "")
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(GetCertBundleFilename(name), pfx, os.ModePerm)
	return pfx, err
}

// GetCert return or create a self-signed Certificate.
func GetCert(name, subject string) (*x509.Certificate, error) {
	certFile := GetCertFilename(name)
	_, err := os.Stat(certFile)
	var certDer []byte
	if os.IsNotExist(err) {
		certDer, err = createCert(name, subject)
		if err != nil {
			return nil, err
		}
	} else {
		certPem, err := ioutil.ReadFile(certFile)
		if err != nil {
			return nil, err
		}
		certDer, _ = Pem2Der(string(certPem))
	}
	return x509.ParseCertificate(certDer)
}

func createCert(name, subject string) ([]byte, error) {
	priv, err := GetCertKey(name)
	if err != nil {
		return nil, err
	}
	sub := wrapCert.CreateSubject(subject)
	x := wrapCert.CreateTemplate(sub.GetRaw(), sub.GetKeyId(), rand2.Int63())
	x.SetExpire(1, 0, 0)
	x.SetUsage([]int{0, 2})
	x.SetExtUsage([]int{1, 2})

	parCert := x.GetRaw()
	x.AuthorityKeyId = sub.GetKeyId()

	/*todo: Change This If need
	x.SetAlgorithm("ed25519", "")
	privObj, err := x509.ParsePKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}
	der, err := x509.CreateCertificate(rand.Reader, x.GetRaw(), parCert, privObj.(ed25519.PrivateKey).Public(), privObj)
	*/
	x.SetAlgorithm("ecdsa", "sha256")
	privObj, err := x509.ParseECPrivateKey(priv)
	if err != nil {
		return nil, err
	}
	der, err := x509.CreateCertificate(rand.Reader, x.GetRaw(), parCert, privObj.Public(), privObj)

	certPem := Der2Pem(der, "CERTIFICATE")

	_ = ioutil.WriteFile(GetCertFilename(name), []byte(certPem), os.ModePerm)

	return der, err
}

// GetCert return or create a private key
func GetCertKey(name string) (key []byte, err error) {
	privFile := GetCertKeyFilename(name)
	_, err = os.Stat(privFile)
	if os.IsNotExist(err) {
		key, err = createCertKey(name)
	} else {
		privPem, err := ioutil.ReadFile(privFile)
		if err != nil {
			return nil, err
		}
		key, _ = Pem2Der(string(privPem))
	}
	return
}

func createCertKey(name string) ([]byte, error) {
	//priv := wrapCert.CreateKey("ed25519", 0)
	priv := wrapCert.CreateKey("ecdsa", 256)
	privPem := Der2Pem(priv, "PRIVATE KEY")
	_ = ioutil.WriteFile(GetCertKeyFilename(name), []byte(privPem), os.ModePerm)

	return priv, nil
}

// Der2Pem encodes der data to pem format
func Der2Pem(data []byte, title string) string {
	return string(pem.EncodeToMemory(&pem.Block{Type: title, Bytes: data}))
}

// Pem2Der decodes pem data to der format
func Pem2Der(PEMString string) ([]byte, string) {
	block, _ := pem.Decode([]byte(PEMString))
	if block == nil {
		return nil, ""
	}
	return block.Bytes, block.Type
}
