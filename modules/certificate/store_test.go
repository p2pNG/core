package certificate

import (
	"golang.org/x/crypto/pkcs12"
	"testing"
)

func TestCreateCertBundle(t *testing.T) {
	t.Run("create Cert Bundle", func(t *testing.T) {
		got, err := createCertBundle("unit-test", "Test Subject")
		if err != nil {
			t.Errorf("GetCertBundle() error = %v", err)
			return
		}
		_, _, err = pkcs12.Decode(got, "")
		if err != nil {
			t.Errorf("Decoder pfx error = %v", err)
			return
		}
	})
}

// todo: Add more test
