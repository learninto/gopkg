package floatlib

import (
	"testing"

	"github.com/learninto/gopkg/stringx"
)

func TestDecimal(t *testing.T) {
	type args struct {
		f float64
		d int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{{
		name: stringx.Rand(),
		args: args{f: 2.2222, d: 2},
		want: 2.22,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decimal(tt.args.f, tt.args.d); got != tt.want {
				t.Errorf("Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
