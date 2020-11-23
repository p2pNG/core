package storage

import (
	"reflect"
	"testing"
)

func TestStatLocalFile(t *testing.T) {
	type args struct {
		filepath    string
		pieceLength int64
		wellKnown   []string
	}
	tests := []struct {
		name    string
		args    args
		wantLf  *LocalFileInfo
		wantErr bool
	}{
		{
			name: "TestStatLocalFile",
			args: args{
				filepath:    TestFilePath,
				pieceLength: TestPieceLength,
				wellKnown:   TestFileWellKnown,
			},
			wantLf:  &TestLocalFileInfo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLf, err := StatLocalFile(tt.args.filepath, tt.args.pieceLength, tt.args.wellKnown)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatLocalFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLf, tt.wantLf) {
				t.Errorf("StatLocalFile() gotLf = %v, want %v", gotLf, tt.wantLf)
			}
		})
	}
}

func Test_statFile(t *testing.T) {
	type args struct {
		filepath    string
		pieceLength int64
		wellKnown   []string
	}
	tests := []struct {
		name         string
		args         args
		wantFileInfo *FileInfo
		wantErr      bool
	}{
		{
			name: "Test_statFile",
			args: args{
				filepath:    TestFilePath,
				pieceLength: TestPieceLength,
				wellKnown:   TestFileWellKnown,
			},
			wantFileInfo: &TestFileInfo,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileInfo, err := statFile(tt.args.filepath, tt.args.pieceLength, tt.args.wellKnown)
			if (err != nil) != tt.wantErr {
				t.Errorf("statFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("statFile() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
		})
	}
}
