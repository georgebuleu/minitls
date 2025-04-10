package crypto

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/georgebuleu/minitls/crypto/pkcs1"
)

const PublicExponent int64 = 65537
const KeyBits int = 1024

type RSAPublicKey struct {
	N *big.Int
	E *big.Int
}

type RSAPrivateKey struct {
	PublicKey RSAPublicKey
	D         *big.Int
	P         *big.Int
	Q         *big.Int
}

func GenerateRSAKeys(bits int, e int64) (*RSAPrivateKey, error) {

	pubE := big.NewInt(e)
	for {
		p, q, err := generatePair(bits)

		if err != nil {
			return nil, err
		}

		phi := computePhi(p, q)
		if new(big.Int).GCD(nil, nil, pubE, phi).Cmp(big.NewInt(1)) != 0 {
			continue
		}

		d := new(big.Int).ModInverse(big.NewInt(PublicExponent), phi)
		if d == nil {
			continue
		}

		n := new(big.Int).Mul(p, q)

		return &RSAPrivateKey{
			PublicKey: RSAPublicKey{
				N: n,
				E: big.NewInt(PublicExponent),
			},
			D: d,
			P: p,
			Q: q,
		}, nil
	}
}

func (pub *RSAPublicKey) Encrypt(msg []byte) (*big.Int, error) {
	padded, err := pkcs1.PadPKCS1v15(msg, KeyBits/8, rand.Reader)
	if err != nil {
		return nil, err
	}

	m := new(big.Int).SetBytes(padded)

	if m.Cmp(pub.N) != -1 || m.Sign() == -1 {
		return nil, errors.New("encryption error")
	}

	// Encrypt: c = m^e mod n
	c := m.Exp(m, pub.E, pub.N)

	return c, nil
}

func (pv *RSAPrivateKey) Decrypt(cipher []byte, keySize int) ([]byte, error) {
	c := new(big.Int).SetBytes(cipher)

	if c.Cmp(pv.PublicKey.N) != -1 || c.Sign() == -1 {
		return nil, errors.New("decryption error")
	}

	pm := c.Exp(c, pv.D, pv.PublicKey.N).Bytes()

	return pkcs1.UnpadPKCS1V15(pm, keySize/8)

}

func computePhi(p, q *big.Int) *big.Int {
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	q1 := new(big.Int).Sub(q, big.NewInt(1))
	return new(big.Int).Mul(p1, q1)
}

func generatePair(keyBits int) (p, q *big.Int, err error) {
	for {
		p, err := rand.Prime(rand.Reader, keyBits)
		if err != nil {
			return nil, nil, err
		}
		q, err := rand.Prime(rand.Reader, keyBits)

		if err != nil {
			return nil, nil, err
		}
		if p.Cmp(q) != 0 {
			return p, q, nil
		}
	}
}
