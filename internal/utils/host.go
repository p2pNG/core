package utils

import (
	"encoding/base64"
	"math/rand"
	"os"
)

// Return the Hostname along with a random string
func GetHostname() string {
	h, err := os.Hostname()
	if err != nil {
		b := make([]byte, 12)
		_, err := rand.Read(b)
		if err != nil {
			return ""
		}
		return base64.RawStdEncoding.EncodeToString(b)
	}
	return h
}
