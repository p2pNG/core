package transfer

import (
	"github.com/p2pNG/core/modules/storage"
	"reflect"
	"testing"
)

func TestDownloadSeed(t *testing.T) {
	type args struct {
		seedInfo storage.SeedInfo
		seedPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestDownloadSeed",
			args: args{
				seedInfo: storage.TestSeedInfo,
				seedPath: storage.TestSeedPath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadSeed(tt.args.seedInfo, tt.args.seedPath); (err != nil) != tt.wantErr {
				t.Errorf("DownloadSeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
