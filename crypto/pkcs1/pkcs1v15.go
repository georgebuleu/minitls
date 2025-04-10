package pkcs1

import (
	"bytes"
	"errors"
	"io"
)

const (
	// According to PKCS#1 v1.5, the padding string (PS) must be at least 8 bytes
	pkcs1MinPaddingLen = 8

	// PKCS#1 v1.5 has 3 fixed bytes in the encoded message (EM):
	//   0x00 (1 byte) | 0x02 (1 byte) | 0x00 (1 byte)
	pkcs1FixedOverhead = 3

	// Therefore, max message size is: keySize - 3 - min PS = keySize - 11
	pkcs1TotalOverhead = pkcs1FixedOverhead + pkcs1MinPaddingLen // = 11
)

func PadPKCS1v15(msg []byte, keySize int, rng io.Reader) ([]byte, error) {
	mLen := len(msg)
	if mLen > keySize-pkcs1TotalOverhead {
		return nil, errors.New("message too long")
	}

	psLen := keySize - mLen - pkcs1FixedOverhead

	ps := make([]byte, psLen)
	for i := 0; i < psLen; {
		if _, err := rng.Read(ps[i : i+1]); err != nil {
			return nil, err
		}
		if ps[i] != 0x00 {
			i++
		}
	}

	em := make([]byte, 0, keySize)
	em = append(em, 0x00, 0x02)
	em = append(em, ps...)
	em = append(em, 0x00)
	em = append(em, msg...)
	return em, nil
}

// EM = 0x00 || 0x02 || PS || 0x00 || M
func UnpadPKCS1V15(pm []byte, keySize int) ([]byte, error) {
	if pm[0] != 0x00 || pm[1] != 0x02 || !bytes.Contains(pm[2:], []byte{0x00}) {
		return nil, errors.New("decryption error")
	}

	ps := pm[bytes.IndexByte(pm[2:], 0x00):]

	if len(ps) < 8 {
		return nil, errors.New("decryption error")
	}

	return ps, nil
}
