/*
 * Copyright (c) 2019 MengYX.
 */

package wrap

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func createEcdsaKey(bit uint) *ecdsa.PrivateKey {
	var curve elliptic.Curve
	switch bit {
	case 224:
		curve = elliptic.P224()
	case 256:
		curve = elliptic.P256()
	case 384:
		curve = elliptic.P384()
	case 521:
		curve = elliptic.P521()
	default:
		return nil
	}
	key, _ := ecdsa.GenerateKey(curve, rand.Reader)
	return key
}

func createRSAKey(bit uint) *rsa.PrivateKey {
	if bit < 64 || bit > 4096 {
		return nil
	}
	key, _ := rsa.GenerateKey(rand.Reader, int(bit))
	return key
}

func createEd25519Key() ed25519.PrivateKey {
	_, priv, _ := ed25519.GenerateKey(rand.Reader)
	return priv
}

// CreateKey create specific x509 private key, and returns the DER data
func CreateKey(category string, bit uint) []byte {
	var key []byte
	switch category {
	case "rsa":
		keyObj := createRSAKey(bit)
		if keyObj != nil {
			key = x509.MarshalPKCS1PrivateKey(keyObj)
		}
	case "ecdsa":
		keyObj := createEcdsaKey(bit)
		if keyObj != nil {
			key, _ = x509.MarshalECPrivateKey(keyObj)
		}
	case "ed25519":
		keyObj := createEd25519Key()
		if keyObj != nil {
			key, _ = x509.MarshalPKCS8PrivateKey(keyObj)
		}
	}
	return key
}
