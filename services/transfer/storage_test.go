package transfer

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/discovery"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	err := database.OpenDB("testing_database")
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	db, err := database.GetDBEngine()
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	err = database.InitBuckets(db, []string{services.SeedHashToSeedDB, services.FileInfoHashToFileDB, services.FileHashToFileDB, services.FileInfoHashToLocalFileDB, "discovery_registry"})
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	m.Run()
}

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
			if err := saveFileInfo(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("saveFileInfo() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := saveSeedInfo(tt.args.seed); (err != nil) != tt.wantErr {
				t.Errorf("saveSeedInfo() error = %v, wantErr %v", err, tt.wantErr)
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
			gotFiles, err := getFileInfoByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileInfoByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("getFileInfoByFileInfoHash() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
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
			gotSeeds, err := getSeedInfoBySeedInfoHash(tt.args.seedInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSeedInfoBySeedInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSeeds, tt.wantSeeds) {
				t.Errorf("getSeedInfoBySeedInfoHash() gotSeeds = %v, want %v", gotSeeds, tt.wantSeeds)
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
			gotFiles, err := getFileInfoByFileHash(tt.args.fileHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileInfoByFileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("getFileInfoByFileHash() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
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
			if err := saveLocalFileInfo(tt.args.fileInfoHash, tt.args.localFileInfo); (err != nil) != tt.wantErr {
				t.Errorf("saveLocalFileInfo() error = %v, wantErr %v", err, tt.wantErr)
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
			gotLocalFileInfo, err := getLocalFileInfoByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalFileInfoByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLocalFileInfo, tt.wantLocalFileInfo) {
				t.Errorf("getLocalFileInfoByFileInfoHash() gotLocalFileInfo = %v, want %v", gotLocalFileInfo, tt.wantLocalFileInfo)
			}
		})
	}
}

func TestGetPeerBySeedHash(t *testing.T) {
	type args struct {
		seedHash string
	}
	tests := []struct {
		name      string
		args      args
		wantPeers []discovery.PeerInfo
		wantErr   bool
	}{
		{
			name: "TestGetPeerBySeedHash",
			args: args{
				seedHash: storage.TestSeedInfoHash,
			},
			wantPeers: []discovery.PeerInfo{
				storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPeers, err := getPeerBySeedHash(tt.args.seedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPeerBySeedHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("getPeerBySeedHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}

func TestGetPeerByFileInfoHash(t *testing.T) {
	type args struct {
		fileInfoHash string
	}
	tests := []struct {
		name      string
		args      args
		wantPeers []discovery.PeerInfo
		wantErr   bool
	}{
		{
			name: "TestGetPeerByFileInfoHash",
			args: args{
				fileInfoHash: storage.TestFileInfoHash,
			},
			wantPeers: []discovery.PeerInfo{
				storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPeers, err := getPeerByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPeerByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("getPeerByFileInfoHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}

func TestGetPeerByFileHash(t *testing.T) {
	type args struct {
		fileHash string
	}
	tests := []struct {
		name      string
		args      args
		wantPeers []discovery.PeerInfo
		wantErr   bool
	}{
		{
			name: "TestGetPeerByFileHash",
			args: args{
				fileHash: storage.TestFileHash,
			},
			wantPeers: []discovery.PeerInfo{
				storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPeers, err := getPeerByFileHash(tt.args.fileHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPeerByFileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("getPeerByFileHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}
