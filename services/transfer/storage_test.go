package transfer

import (
	"github.com/p2pNG/core/modules/storage"
	"reflect"
	"testing"
)

var testFileInfo = storage.FileInfo{
	Size:        1024,
	Hash:        "bdae991688970e57bba929524ba9b3d82eda5795",
	PieceLength: 1024,
	PieceHash: []string{
		"bdae991688970e57bba929524ba9b3d82eda5795",
	},
	WellKnown: nil,
}

var testSeedInfo = storage.SeedInfo{
	Title: "TestSeedInfoTitle",
	Files: []storage.SeedFileItem{
		{
			Path:            "F:\\www\\TestSeed\\TestFile.txt",
			Size:            1024,
			Hash:            testFileHash,
			RecFileInfoHash: "bdae991688970e57bba929524ba9b3d82eda5795",
			RecPieceLength:  1024,
		},
	},
	ExtraInfo: nil,
	WellKnown: nil,
}

var testFileInfoHash = storage.HashFileInfo(testFileInfo)
var testSeedInfoHash = storage.HashSeedInfo(testSeedInfo)
var testFileHash = "bdae991688970e57bba929524ba9b3d82eda5795"

func TestSaveFileInfo(t *testing.T) {
	type args struct {
		file storage.FileInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveFileInfo",
			args: args{
				file: testFileInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFileInfo(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("SaveFileInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveSeedInfo(t *testing.T) {
	type args struct {
		seed storage.SeedInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveSeedInfo",
			args: args{
				seed: testSeedInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveSeedInfo(tt.args.seed); (err != nil) != tt.wantErr {
				t.Errorf("SaveSeedInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileInfoByFileInfoHash(t *testing.T) {
	type args struct {
		fileInfoHash string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles storage.FileInfo
		wantErr   bool
	}{
		{
			name: "TestGetFileInfoByFileInfoHash",
			args: args{
				fileInfoHash: testFileInfoHash,
			},
			wantFiles: testFileInfo,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := GetFileInfoByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfoByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("GetFileInfoByFileInfoHash() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func TestGetSeedInfo(t *testing.T) {
	type args struct {
		seedInfoHash string
	}
	tests := []struct {
		name      string
		args      args
		wantSeeds storage.SeedInfo
		wantErr   bool
	}{
		{
			name: "TestGetSeedInfo",
			args: args{
				seedInfoHash: testSeedInfoHash,
			},
			wantSeeds: testSeedInfo,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSeeds, err := GetSeedInfo(tt.args.seedInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSeedInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSeeds, tt.wantSeeds) {
				t.Errorf("GetSeedInfo() gotSeeds = %v, want %v", gotSeeds, tt.wantSeeds)
			}
		})
	}
}

func TestGetFileInfoByFileHash(t *testing.T) {
	type args struct {
		fileHash string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []storage.FileInfo
		wantErr   bool
	}{
		{
			name: "TestGetFileInfoByFileHash",
			args: args{
				fileHash: testFileHash,
			},
			wantFiles: []storage.FileInfo{
				testFileInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := GetFileInfoByFileHash(tt.args.fileHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfoByFileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("GetFileInfoByFileHash() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func TestGetFileHashList(t *testing.T) {
	tests := []struct {
		name             string
		wantFileHashList []string
		wantErr          bool
	}{
		{
			name: "TestGetFileHashList",
			wantFileHashList: []string{
				testFileHash,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileHashList, err := GetFileHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileHashList, tt.wantFileHashList) {
				t.Errorf("GetFileHashList() gotFileHashList = %v, want %v", gotFileHashList, tt.wantFileHashList)
			}
		})
	}
}

func TestGetSeedInfoHashList(t *testing.T) {
	tests := []struct {
		name                 string
		wantSeedInfoHashList []string
		wantErr              bool
	}{
		{
			name: "TestGetSeedInfoHashList",
			wantSeedInfoHashList: []string{
				testSeedInfoHash,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSeedInfoHashList, err := GetSeedInfoHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSeedInfoHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSeedInfoHashList, tt.wantSeedInfoHashList) {
				t.Errorf("GetSeedInfoHashList() gotSeedInfoHashList = %v, want %v", gotSeedInfoHashList, tt.wantSeedInfoHashList)
			}
		})
	}
}

func TestGetFileInfoHashList(t *testing.T) {
	tests := []struct {
		name                 string
		wantFileInfoHashList []string
		wantErr              bool
	}{
		{
			name: "TestGetFileInfoHashList",
			wantFileInfoHashList: []string{
				testFileInfoHash,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfoHashList, err := GetFileInfoHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfoHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfoHashList, tt.wantFileInfoHashList) {
				t.Errorf("GetFileInfoHashList() gotFileInfoHashList = %v, want %v", gotFileInfoHashList, tt.wantFileInfoHashList)
			}
		})
	}
}
