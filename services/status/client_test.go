package status

import (
	"github.com/p2pNG/core/services/discovery"
	"testing"
)

func Test_exchangeFileHash(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_exchangeFileHash",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exchangeFileHash(); (err != nil) != tt.wantErr {
				t.Errorf("exchangeFileHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exchangeFileInfoHash(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_exchangeFileInfoHash",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exchangeFileInfoHash(); (err != nil) != tt.wantErr {
				t.Errorf("exchangeFileInfoHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exchangePeers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_exchangePeers",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exchangePeers(); (err != nil) != tt.wantErr {
				t.Errorf("exchangePeers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exchangeSeeds(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test_exchangeSeeds",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := exchangeSeeds(); (err != nil) != tt.wantErr {
				t.Errorf("exchangeSeeds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_visitPeers(t *testing.T) {
	type args struct {
		fn func(peer discovery.PeerInfo) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Test_visitPeers",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := visitPeers(tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("visitPeers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
