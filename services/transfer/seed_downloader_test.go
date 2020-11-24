package transfer

import (
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/storage"
	"testing"
)

func TestDownloadSeed(t *testing.T) {
	type args struct {
		seedInfoHash string
		seedPath     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestDownloadSeed",
			args: args{
				seedInfoHash: storage.TestSeedInfoHash,
				seedPath:     storage.TestSeedPath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.RemoveFilePathAll(tt.args.seedPath)
			if err := DownloadSeed(tt.args.seedInfoHash, tt.args.seedPath); (err != nil) != tt.wantErr {
				t.Errorf("DownloadSeed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
