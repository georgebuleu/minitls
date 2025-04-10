package encoding

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestI2OSP_OS2IP_RoundTrip(t *testing.T) {
	for i := range 500 {
		// Generate a random bit length (up to 256 bits)
		bitLen := rand.Intn(256) + 1
		x := new(big.Int).Rand(rand.New(rand.NewSource(int64(i))), new(big.Int).Lsh(big.NewInt(1), uint(bitLen)))

		// Compute minimal xLen in bytes
		xLen := (x.BitLen() + 7) / 8

		// Add random padding
		padding := rand.Intn(5) // 0–4 extra bytes
		xLen += padding

		encoded, err := IntToOctetString(x, xLen)
		if err != nil {
			t.Fatalf("unexpected error on I2OSP(%v, %d): %v", x, xLen, err)
		}

		decoded := OctetStringToInt(encoded)
		if x.Cmp(decoded) != 0 {
			t.Errorf("Round trip failed: x=%v, decoded=%v", x, decoded)
		}
	}
}
