package status

import (
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services/discovery"
	"reflect"
	"testing"
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
			if err := saveFileInfoHash(tt.args.fileHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("saveFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := saveSeedInfoHash(tt.args.seedHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("saveSeedInfoHash() error = %v, wantErr %v", err, tt.wantErr)
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
			if err := saveFileHash(tt.args.fileHashList, tt.args.peer); (err != nil) != tt.wantErr {
				t.Errorf("saveFileHash() error = %v, wantErr %v", err, tt.wantErr)
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
			gotFileHashList, err := getFileHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileHashList, tt.wantFileHashList) {
				t.Errorf("getFileHashList() gotFileHashList = %v, want %v", gotFileHashList, tt.wantFileHashList)
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
			gotSeedInfoHashList, err := getSeedInfoHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("getSeedInfoHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSeedInfoHashList, tt.wantSeedInfoHashList) {
				t.Errorf("getSeedInfoHashList() gotSeedInfoHashList = %v, want %v", gotSeedInfoHashList, tt.wantSeedInfoHashList)
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
			gotFileInfoHashList, err := getFileInfoHashList()
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileInfoHashList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfoHashList, tt.wantFileInfoHashList) {
				t.Errorf("getFileInfoHashList() gotFileInfoHashList = %v, want %v", gotFileInfoHashList, tt.wantFileInfoHashList)
			}
		})
	}
}

func Test_savePeerPieceInfoList(t *testing.T) {
	type args struct {
		ppInfoList map[string]storage.PeerPieceInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test_savePeerPieceInfoList",
			args: args{
				ppInfoList: storage.TestPPInfoList,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := savePeerPieceInfoList(tt.args.ppInfoList); (err != nil) != tt.wantErr {
				t.Errorf("savePeerPieceInfoList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPeerPieceInfoList(t *testing.T) {
	tests := []struct {
		name           string
		wantPpInfoList map[string]storage.PeerPieceInfo
		wantErr        bool
	}{
		{
			name:           "Test_getPeerPieceInfoList",
			wantPpInfoList: storage.TestPPInfoList,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPpInfoList, err := getPeerPieceInfoList()
			if (err != nil) != tt.wantErr {
				t.Errorf("getPeerPieceInfoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPpInfoList, tt.wantPpInfoList) {
				t.Errorf("getPeerPieceInfoList() gotPpInfoList = %v, want %v", gotPpInfoList, tt.wantPpInfoList)
			}
		})
	}
}
