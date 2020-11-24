package services

import (
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services/discovery"
	"testing"
)

func TestPeerInfoToStringAddr(t *testing.T) {
	type args struct {
		info discovery.PeerInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestPeerInfoToStringAddr",
			args: args{storage.TestPeerInfo},
			want: "https://127.0.0.1:6480",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PeerInfoToStringAddr(tt.args.info); got != tt.want {
				t.Errorf("PeerInfoToStringAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
