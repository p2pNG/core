package transfer

import (
	"github.com/p2pNG/core/modules/storage"
	"reflect"
	"testing"
)

func TestQueryFileInfoByFileInfoHash(t *testing.T) {
	type args struct {
		peerAddr     string
		fileInfoHash string
	}
	tests := []struct {
		name         string
		args         args
		wantFileInfo *storage.FileInfo
		wantErr      bool
	}{
		{
			name: "TestQueryFileInfoByFileInfoHash",
			args: args{
				peerAddr:     storage.TestPeerAddr,
				fileInfoHash: storage.TestFileInfoHash,
			},
			wantFileInfo: &storage.TestFileInfo,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, err := QueryFileInfoByFileInfoHash(tt.args.peerAddr, tt.args.fileInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFileInfoByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("QueryFileInfoByFileInfoHash() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
		})
	}
}

func TestQuerySeedInfoBySeedInfoHash(t *testing.T) {
	type args struct {
		peerAddr     string
		seedInfoHash string
	}
	tests := []struct {
		name         string
		args         args
		wantSeedInfo *storage.SeedInfo
		wantErr      bool
	}{
		{
			name: "TestQuerySeedInfoBySeedInfoHash",
			args: args{
				peerAddr:     storage.TestPeerAddr,
				seedInfoHash: storage.TestSeedInfoHash,
			},
			wantSeedInfo: &storage.TestSeedInfo,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSeedInfo, err := QuerySeedInfoBySeedInfoHash(tt.args.peerAddr, tt.args.seedInfoHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuerySeedInfoBySeedInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSeedInfo, tt.wantSeedInfo) {
				t.Errorf("QuerySeedInfoBySeedInfoHash() gotSeedInfo = %v, want %v", gotSeedInfo, tt.wantSeedInfo)
			}
		})
	}
}

func TestQueryFileInfoByFileHash(t *testing.T) {
	type args struct {
		peerAddr string
		fileHash string
	}
	tests := []struct {
		name         string
		args         args
		wantFileInfo *storage.FileInfo
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, err := QueryFileInfoByFileHash(tt.args.peerAddr, tt.args.fileHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFileInfoByFileHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("QueryFileInfoByFileHash() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
		})
	}
}
