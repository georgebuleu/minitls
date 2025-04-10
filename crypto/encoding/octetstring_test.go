package encoding

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_IntToOctetString(t *testing.T) {
	tests := []struct {
		name string
		x    *big.Int
		xLen int
		want []byte
		err  bool
	}{
		{
			name: "negative integer",
			x:    big.NewInt(-1),
			xLen: 4,
			want: nil,
			err:  true,
		},
		{
			name: "integer too large",
			x:    new(big.Int).Lsh(big.NewInt(1), uint(64)),
			xLen: 4, // too small for 2 ^ 64
			want: nil,
			err:  true,
		},
		{
			name: "12345 to bytes",
			x:    big.NewInt(123456),
			xLen: 5,
			want: []byte{0x00, 0x00, 0x01, 0xe2, 0x40},
			err:  false,
		},
		{
			name: "0 to 4 bytes",
			x:    big.NewInt(0),
			xLen: 4,
			want: []byte{0x00, 0x00, 0x00, 0x00},
			err:  false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IntToOctetString(tc.x, tc.xLen)

			if (err != nil) != tc.err {
				t.Errorf("expected error: %v, got: %v", tc.err, err)
			}
			if !tc.err && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InToOctetString(%v, %d) = %x, want %x", tc.x, tc.xLen, got, tc.want)
			}
		})
	}
}

func Test_OctetStringToInt(t *testing.T) {
	tests := []struct {
		name string
		x    []byte
		want *big.Int
	}{
		{
			name: "12345 as bytes",
			x:    []byte{0x00, 0x00, 0x01, 0xe2, 0x40},
			want: big.NewInt(123456),
		},
		{
			name: "zero input",
			x:    []byte{0x00, 0x00},
			want: big.NewInt(0),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := OctetStringToInt(tc.x)
			if got.Cmp(tc.want) != 0 {
				t.Errorf("OS2IP(%x) = %v, want %v", tc.x, got, tc.want)
			}

		})
	}
}
