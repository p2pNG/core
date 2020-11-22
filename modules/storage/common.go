package storage

import (
	"time"
)

// DefaultFilePieceLength used for the default parameter to split a file to several blocks
// MinFilePieceLength used for the min parameter to split a file to several blocks
const (
	separator              = ":"
	DefaultFilePieceLength = 4 * 1024 * 1024
	MinFilePieceLength     = 1024 * 1024
)

// FileInfo describes how to download a file
type FileInfo struct {
	Size        int64
	Hash        string
	PieceLength int64
	PieceHash   []string
	WellKnown   []string
}

// SeedFileItem describes base info of a file
type SeedFileItem struct {
	Path            string
	Size            int64
	Hash            string
	RecPieceLength  int64
	RecFileInfoHash string
}

// SeedInfo describes a p2pNG seed
type SeedInfo struct {
	Title     string
	Files     []SeedFileItem
	ExtraInfo map[string]string
	WellKnown []string
}

// LocalFileInfo describes a local file’ path and last modified time
type LocalFileInfo struct {
	FileInfo
	Path       string
	LastModify time.Time
}
