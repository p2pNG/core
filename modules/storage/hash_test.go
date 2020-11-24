package storage

import (
	"testing"
)

func TestHashFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name         string
		args         args
		wantFileHash string
		wantErr      bool
	}{
		{
			name: "TestHashFile",
			args: args{
				filePath: TestFilePath,
			},
			wantFileHash: TestFileHash,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileHash, err := HashFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileHash != tt.wantFileHash {
				t.Errorf("HashFile() gotFileHash = %v, want %v", gotFileHash, tt.wantFileHash)
			}
		})
	}
}

func TestHashFileInfo(t *testing.T) {
	type args struct {
		fileInfo FileInfo
	}
	tests := []struct {
		name             string
		args             args
		wantFileInfoHash string
	}{
		{
			name: "TestHashFileInfo",
			args: args{
				fileInfo: TestFileInfo,
			},
			wantFileInfoHash: TestFileInfoHash,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFileInfoHash := HashFileInfo(tt.args.fileInfo); gotFileInfoHash != tt.wantFileInfoHash {
				t.Errorf("HashFileInfo() = %v, want %v", gotFileInfoHash, tt.wantFileInfoHash)
			}
		})
	}
}

func TestHashFilePieceInBytes(t *testing.T) {
	type args struct {
		piece []byte
	}
	tests := []struct {
		name          string
		args          args
		wantPieceHash string
		wantErr       bool
	}{
		{
			name: "TestHashFilePieceInBytes",
			args: args{
				piece: TestPiece,
			},
			wantPieceHash: TestPieceHash,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPieceHash, err := HashFilePieceInBytes(tt.args.piece)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashFilePieceInBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPieceHash != tt.wantPieceHash {
				t.Errorf("HashFilePieceInBytes() gotPieceHash = %v, want %v", gotPieceHash, tt.wantPieceHash)
			}
		})
	}
}

func TestHashSeedInfo(t *testing.T) {
	type args struct {
		seedInfo SeedInfo
	}
	tests := []struct {
		name             string
		args             args
		wantSeedInfoHash string
	}{
		{
			name: "TestHashSeedInfo",
			args: args{
				seedInfo: TestSeedInfo,
			},
			wantSeedInfoHash: TestSeedInfoHash,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSeedInfoHash := HashSeedInfo(tt.args.seedInfo); gotSeedInfoHash != tt.wantSeedInfoHash {
				t.Errorf("HashSeedInfo() = %v, want %v", gotSeedInfoHash, tt.wantSeedInfoHash)
			}
		})
	}
}
