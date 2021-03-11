package bitperms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermissionValue_Serialize(t *testing.T) {
	tests := []struct {
		name string
		pv   PermissionValue
		want string
	}{
		{
			name: "Serialize-01",
			pv:   1902837283,
			want: "1902837283",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pv.Serialize())
		})
	}
}

func TestPermissionValue_Deserialize(t *testing.T) {
	tests := []struct {
		name       string
		serialized string
		want       PermissionValue
	}{
		{
			name:       "Deserialize-01",
			serialized: "192743628728",
			want:       192743628728,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Deserialize(tt.serialized)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestPermissionValue_HasFlag(t *testing.T) {
	type args struct {
		flag uint64
	}
	tests := []struct {
		name string
		pv   PermissionValue
		args args
		want bool
	}{
		{
			name: "HasFlag.01",
			pv:   0b00010010,
			args: args{
				flag: 0b00010000,
			},
			want: true,
		},
		{
			name: "HasFlag.02",
			pv:   0b11111111,
			args: args{
				flag: 0b11111111,
			},
			want: true,
		},
		{
			name: "HasFlag.03",
			pv:   0b1010101011101011101,
			args: args{
				flag: 0b0000001000000010000,
			},
			want: true,
		},
		{
			name: "HasFlag.04",
			pv:   0b0111111111111111111111111111111111111111111111111111111111111111,
			args: args{
				flag: 0b0011100000000000000011000000000000100000000000000000000000100000,
			},
			want: true,
		},
		{
			name: "HasFlag.05",
			pv:   0b11111110,
			args: args{
				flag: 0b00000001,
			},
			want: false,
		},
		{
			name: "HasFlag.06",
			pv:   0b11111111,
			args: args{
				flag: 0b00000000,
			},
			want: true,
		},
		{
			name: "HasFlag.07",
			pv:   0b0011100000000000000011000000000000100000000001100000000000100000,
			args: args{
				flag: 0b0011110000001000000000000000000000100000000001100000000000100000,
			},
			want: false,
		},
		{
			name: "HasFlag.08",
			pv:   0b01001000,
			args: args{
				flag: 0b01000000 | 0b00001000,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pv.HasFlag(tt.args.flag))
		})
	}
}

func TestPermissionValue_HasFlags(t *testing.T) {
	type args struct {
		flags []uint64
	}
	tests := []struct {
		name string
		pv   PermissionValue
		args args
		want bool
	}{
		{
			name: "HasFlags.01",
			pv:   0b11111111,
			args: args{
				flags: []uint64{0b00000001, 0b10000000, 0b00010000, 0b00000100},
			},
			want: true,
		},
		{
			name: "HasFlags.02",
			pv:   0b11111110,
			args: args{
				flags: []uint64{0b00000001, 0b10000000},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.pv.HasFlags(tt.args.flags...))
		})
	}
}
