package pkcs1

import (
	"io"
	"reflect"
	"testing"
)

func Test_PadPKCS1v15(t *testing.T) {
	tests := []struct {
		name    string
		msg     []byte
		keySize int
		rng     io.Reader
		want    []byte
		err     bool
	}{
		{
			name:    "long message",
			msg:     []byte{0x00, 0x04},
			keySize: 12, // very small keysize
			rng:     nil,
			want:    nil,
			err:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := PadPKCS1v15(tc.msg, tc.keySize, nil)

			if (err != nil) != tc.err {
				t.Fatalf("expected error: %v, got: %v", tc.err, err)
			}
			if !tc.err && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PadPKCS1V15(%v, %d) = %x, want %x", tc.msg, tc.keySize, got, tc.want)
			}
		})
	}
}

func Test_UnpadPKCS1v15(t *testing.T) {
	tests := []struct {
		name    string
		pm      []byte
		keySize int
		want    []byte
		err     bool
	}{
		{
			name:    "first byte not 0",
			pm:      []byte{0x01, 0x04},
			keySize: 256,
			want:    nil,
			err:     true,
		},
		{
			name:    "second byte not 2",
			pm:      []byte{0x01, 0x04},
			keySize: 256,
			want:    nil,
			err:     true,
		},
		{
			name:    "no 0 padding",
			pm:      []byte{0x00, 0x02, 0x2, 0x3},
			keySize: 256,
			want:    nil,
			err:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := UnpadPKCS1V15(tc.pm, tc.keySize)

			if (err != nil) != tc.err {
				t.Errorf("expected error: %v, got: %v", tc.err, err)
			}
			if !tc.err && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("UnpadPKCS1V15(%v, %d) = %x, want %x", tc.pm, tc.keySize, got, tc.want)
			}
		})
	}
}
