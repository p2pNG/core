package status

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/storage"
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
	err = database.InitBuckets(db, []string{SeedHashToPeerDB, FileInfoHashToPeerDB, FileHashToPeerDB, "discovery_registry"})
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	m.Run()
}

func TestSaveFileInfoHash(t *testing.T) {
	type args struct {
		fileHashList []string
		peer         discovery.PeerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveFileInfoHash",
			args: args{
				fileHashList: []string{
					storage.TestFileInfoHash,
				},
				peer: storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFileInfoHash(tt.args.fileHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("SaveFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
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
			gotPeers, err := GetPeerByFileInfoHash(tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeerByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("GetPeerByFileInfoHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}

func TestSaveSeedInfoHash(t *testing.T) {
	type args struct {
		seedHashList []string
		peer         discovery.PeerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveSeedInfoHash",
			args: args{
				seedHashList: []string{
					storage.TestSeedInfoHash,
				},
				peer: storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveSeedInfoHash(tt.args.seedHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("SaveSeedInfoHash() error = %v, wantErr %v", err, tt.wantErr)
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
			gotPeers, err := GetPeerBySeedHash(tt.args.seedHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeerBySeedHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("GetPeerBySeedHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}

func TestSaveFileHash(t *testing.T) {
	type args struct {
		fileHashList []string
		peer         discovery.PeerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSaveFileHash",
			args: args{
				fileHashList: []string{
					storage.TestFileHash,
				},
				peer: storage.TestPeerInfo,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFileHash(tt.args.fileHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("SaveFileHash() error = %v, wantErr %v", err, tt.wantErr)
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
			gotPeers, err := GetPeerByFileHash(tt.args.fileHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeerByFileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPeers, tt.wantPeers) {
				t.Errorf("GetPeerByFileHash() gotPeers = %v, want %v", gotPeers, tt.wantPeers)
			}
		})
	}
}
