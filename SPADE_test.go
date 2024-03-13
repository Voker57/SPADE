package SPADE

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func TestSpade(t *testing.T) {
	ptVecSize := 1000
	numUsers := 50

	spade := NewSpade()
	var sk, pk, dk, res []*big.Int
	var ciphertexts [][]*big.Int

	// generate random data for test
	dummyData := make([][]int, numUsers)
	for j := 0; j < numUsers; j++ {
		dummyData[j] = make([]int, ptVecSize)
		for i := 0; i < ptVecSize; i++ {
			dummyData[j][i] = int(rand.Int63n(10) + 1)
		}
	}

	t.Run("Setup", func(t *testing.T) {
		sk, pk = spade.setup(ptVecSize, numUsers)
	})

	// create dummy registration keys
	regKeys := make([]*big.Int, numUsers)
	for i := 0; i < numUsers; i++ {
		alpha := spade.RandomElementInZMod()
		regKeys[i] = spade.register(alpha)
	}

	t.Run("Encryption", func(t *testing.T) {
		ciphertexts = spade.encrypt(pk, regKeys[0], dummyData[0])
	})

	t.Run("FEKeyGen", func(t *testing.T) {
		dk = spade.keygen(0, 1, sk, regKeys)
	})

	t.Run("Decryption", func(t *testing.T) {
		res = spade.decrypt(dk, 1, ciphertexts)
	})

	if len(res) != len(dummyData[0]) {
		t.Errorf("Decrypt failed: invalid length of decrypted message slice")
	}
	fmt.Println(res)
}
