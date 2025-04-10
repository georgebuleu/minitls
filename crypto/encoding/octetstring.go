package encoding

import (
	"errors"
	"math/big"
)

// OS2IP converts an octet string to a nonnegative integer
func OctetStringToInt(x []byte) *big.Int {
	return new(big.Int).SetBytes(x)
}

//	I2OSP converts a nonnegative integer to an octet string of a
// specified length.
func IntToOctetString(x *big.Int, xLen int) ([]byte, error) {
	if x.Sign() < 0 {
		return nil, errors.New("negative integer")
	}

	limit := new(big.Int).Lsh(big.NewInt(1), uint(8*xLen))

	if x.Cmp(limit) >= 0 {
		return nil, errors.New("integer too large")
	}

	result := x.Bytes()
	pResult := make([]byte, xLen)
	copy(pResult[xLen-len(result):], result)

	return pResult, nil

}
