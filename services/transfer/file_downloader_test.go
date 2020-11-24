package transfer

import (
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/storage"
	"testing"
)

func TestNewFileDownloaderByFileInfoHash(t *testing.T) {
	type args struct {
		fileInfoHash string
		desFilePath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				fileInfoHash: storage.TestFileInfoHash,
				desFilePath:  storage.TestDownloadFilePath,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.RemoveFilePathAll(tt.args.desFilePath)
			downloader, err := NewFileDownloaderByFileInfoHash(tt.args.fileInfoHash, tt.args.desFilePath)
			if (err != nil) != tt.wantErr || downloader == nil {
				t.Errorf("NewFileDownloaderByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = downloader.DownloadFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFileDownloaderByFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
