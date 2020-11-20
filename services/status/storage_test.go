package status

import (
	"github.com/p2pNG/core/services/discovery"
	"net"
	"reflect"
	"testing"
	"time"
)

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
			name: "Test_saveSeeds",
			args: args{
				fileHashList: []string{
					"bdae991688970e57bba929524ba9b3d82eda5795",
				},
				peer: discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
			name: "TestGetPeerBySeedHash",
			args: args{
				fileInfoHash: "bdae991688970e57bba929524ba9b3d82eda5795",
			},
			wantPeers: []discovery.PeerInfo{
				discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
			name: "Test_saveSeeds",
			args: args{
				seedHashList: []string{
					"bdae991688970e57bba929524ba9b3d82eda5795",
				},
				peer: discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
				seedHash: "bdae991688970e57bba929524ba9b3d82eda5795",
			},
			wantPeers: []discovery.PeerInfo{
				discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
			name: "Test_saveSeeds",
			args: args{
				fileHashList: []string{
					"bdae991688970e57bba929524ba9b3d82eda5795",
				},
				peer: discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
			name: "TestGetPeerBySeedHash",
			args: args{
				fileHash: "bdae991688970e57bba929524ba9b3d82eda5795",
			},
			wantPeers: []discovery.PeerInfo{
				discovery.PeerInfo{
					Address:  net.ParseIP("192.168.1.101"),
					Port:     6060,
					DNS:      []string{"dns"},
					LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
				},
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
