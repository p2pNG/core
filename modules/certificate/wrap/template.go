/*
 * Copyright (c) 2019 MengYX.
 */

package wrap

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

// CertTemplate is a x509.Certificate wrapper for template use
type CertTemplate x509.Certificate

// CreateTemplate create a most simple CertTemplate
func CreateTemplate(subject *pkix.Name, keyID []byte, serial int64) *CertTemplate {
	template := CertTemplate{
		SerialNumber: big.NewInt(serial),
		Subject:      *subject,
		SubjectKeyId: keyID,
	}
	return &template
}

// SetExpire can simply set the expire date (start from today 00:00)
func (cert *CertTemplate) SetExpire(year, month, day int) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	cert.NotBefore = today
	cert.NotAfter = today.AddDate(year, month, day)
}

// SetConstraint can simply set if the cert can be a CA
// While cert is a CA, you can limit the sub CA path length
func (cert *CertTemplate) SetConstraint(isCA, LimitPathLen bool, pathLen int) {
	cert.BasicConstraintsValid = true
	cert.IsCA = isCA
	if isCA {
		if LimitPathLen {
			if 0 == pathLen {
				cert.MaxPathLenZero = true
			} else {
				cert.MaxPathLenZero = false
				cert.MaxPathLen = pathLen
			}
		} else {
			cert.MaxPathLenZero = false
			cert.MaxPathLen = 0
		}
	}
}

// SetUsage can simply set the cert basic usage
// input the index of the x509.KeyUsage (start from 1)
func (cert *CertTemplate) SetUsage(usage []int) {
	all := 0
	for idx := range usage {
		all = all | 1<<usage[idx]
	}
	cert.KeyUsage = x509.KeyUsage(all)
}

// SetExtUsage can simply set the cert of the list x509.ExtKeyUsage
func (cert *CertTemplate) SetExtUsage(usage []int) {
	var all []x509.ExtKeyUsage
	for e := range usage {
		all = append(all, x509.ExtKeyUsage(usage[e]))
	}
	cert.ExtKeyUsage = all
}

// GetRaw return the raw x509.Certificate
func (cert *CertTemplate) GetRaw() *x509.Certificate {
	return (*x509.Certificate)(cert)
}

// SetAlgorithm can simply set the algorithm that used for validate the cert itself
func (cert *CertTemplate) SetAlgorithm(category string, name string) {
	//仅用于验证证书本身 无关加密算法
	algorithm := 0
	switch category {

	case "ecdsa":
		algorithm = 9
	case "rsa":
		algorithm = 12
	case "ed25519":
		algorithm = 16
	}

	switch name {
	case "sha256":
		algorithm++
	case "sha384":
		algorithm += 2
	case "sha512":
		algorithm += 3
	}
	cert.SignatureAlgorithm = x509.SignatureAlgorithm(algorithm)
}
