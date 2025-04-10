package crypto

import (
	"testing"
)

func TestGenerateRSAKeys(t *testing.T) {
	pk, err := GenerateRSAKeys(1024, PublicExponent)
	if err != nil {
		t.Fatal(err)
	}

	if pk.D == nil {
		t.Errorf("D is nil")
	}

	if pk.Q == nil {
		t.Errorf("Q is nil")
	}

	if pk.PublicKey.N == nil {
		t.Errorf("PublicKey n is nil")
	}

	if pk.PublicKey.E == nil {
		t.Errorf("PublicKey E is nil")
	}

}



