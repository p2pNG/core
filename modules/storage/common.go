package storage

import "time"

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

// LocalFileInfo extends FileInfo with the filepath and modify time
type LocalFileInfo struct {
	FileInfo
	Path       string
	LastModify time.Time
}
