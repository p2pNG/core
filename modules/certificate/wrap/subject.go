/*
 * Copyright (c) 2019 MengYX.
 */

package wrap

import (
	"crypto/sha256"
	"crypto/x509/pkix"
	"encoding/asn1"
	"github.com/p2pNG/core/internal/logging"
	"go.uber.org/zap"
)

// Subject is a pkix.Name wrapper for template use
type Subject pkix.Name

// CreateSubject create a most simple Subject
func CreateSubject(name string) *Subject {
	return &Subject{CommonName: name}
}

// SetLocation can simply set the location
func (s *Subject) SetLocation(country, province, city string) {
	s.Country = []string{country}
	if province != "" {
		s.Province = []string{province}
	}
	if city != "" {
		s.Locality = []string{city}
	}
}

// SetOrg can simply set the org name and org unit name
func (s *Subject) SetOrg(org, orgUnit string) {
	s.Organization = []string{org}
	if orgUnit != "" {
		s.OrganizationalUnit = []string{orgUnit}
	}
}

// SetSerial can simply set the serial of the subject
func (s *Subject) SetSerial(serial string) {
	s.SerialNumber = serial
}

// GetRaw return the raw pkix.Name after fill up the template
func (s *Subject) GetRaw() *pkix.Name {
	return (*pkix.Name)(s)
}

// GetKeyID return the hash sum of this Subject
func (s *Subject) GetKeyID() []byte {
	idHash := sha256.New()
	data, err := asn1.Marshal(*s)
	if err != nil {
		logging.Log().Error("compile Subject failed", zap.Error(err))
	}
	_, _ = idHash.Write(data)
	//todo: Use Another Method to Generate
	toEnc := "p2pNG-User-Id:" + s.CommonName
	_, _ = idHash.Write([]byte(toEnc))
	keyID := idHash.Sum(nil)
	return keyID
}
