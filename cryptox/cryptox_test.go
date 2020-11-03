package cryptox

import (
	"testing"
)

func TestCrc32IEEE(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "case1", args: args{data: []byte("1289898988989")}, want: 1538950961},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc32IEEE(tt.args.data); got != tt.want {
				t.Errorf("Crc32IEEE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha256Encode(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{param: "case00000"}, want: "386a2f3368200c0df560f876c875e1b25c4b43a4e4ec8c63dbc6d2b53218fab7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256Encode(tt.args.param); got != tt.want {
				t.Errorf("Sha256Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
