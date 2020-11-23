package transfer

import (
	"github.com/p2pNG/core/modules/storage"
	"reflect"
	"testing"
)

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
				file: storage.TestFileInfo,
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
				seed: storage.TestSeedInfo,
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
				fileInfoHash: storage.TestFileInfoHash,
			},
			wantFiles: storage.TestFileInfo,
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
				seedInfoHash: storage.TestSeedInfoHash,
			},
			wantSeeds: storage.TestSeedInfo,
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
				fileHash: storage.TestFileHash,
			},
			wantFiles: []storage.FileInfo{
				storage.TestFileInfo,
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
				storage.TestFileHash,
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
				storage.TestSeedInfoHash,
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
				storage.TestFileInfoHash,
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

func TestSaveLocalFileInfo(t *testing.T) {
	type args struct {
		fileInfoHash  string
		localFileInfo storage.LocalFileInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveLocalFileInfo",
			args: args{
				fileInfoHash:  storage.TestFileInfoHash,
				localFileInfo: storage.TestLocalFileInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveLocalFileInfo(tt.args.fileInfoHash, tt.args.localFileInfo); (err != nil) != tt.wantErr {
				t.Errorf("SaveLocalFileInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLocalFileInfoByFileInfoHash(t *testing.T) {
	type args struct {
		fileInfoHash string
	}
	tests := []struct {
		name              string
		args              args
		wantLocalFileInfo storage.LocalFileInfo
		wantErr           bool
	}{
		{
			name: "TestGetLocalFileInfoByFileInfoHash",
			args: args{
				fileInfoHash: storage.TestFileInfoHash,
			},
			wantLocalFileInfo: storage.TestLocalFileInfo,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocalFileInfo, err := GetLocalFileInfoByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalFileInfoByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLocalFileInfo, tt.wantLocalFileInfo) {
				t.Errorf("GetLocalFileInfoByFileInfoHash() gotLocalFileInfo = %v, want %v", gotLocalFileInfo, tt.wantLocalFileInfo)
			}
		})
	}
}
