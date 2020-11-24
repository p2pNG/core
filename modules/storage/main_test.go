package storage

import (
	"os"
	"testing"
)

// TestMain open db
func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}
