package SPADE

import (
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkSpade(b *testing.B) {
	benchmarkSpade(b)
}

func benchmarkSpade(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	ptVecSize := 100
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

	b.Run("Setup", func(b *testing.B) {
		b.ResetTimer()
		sk, pk = spade.setup(ptVecSize, numUsers)
	})

	// create dummy registration keys
	regKeys := make([]*big.Int, numUsers)
	for i := 0; i < numUsers; i++ {
		alpha := spade.RandomElementInZMod()
		regKeys[i] = spade.register(alpha)
	}

	b.Run("Encryption", func(b *testing.B) {
		b.ResetTimer()
		ciphertexts = spade.encrypt(pk, regKeys[0], dummyData[0])
	})

	b.Run("FEKeyDer", func(b *testing.B) {
		b.ResetTimer()
		dk = spade.keygen(0, 1, sk, regKeys)
	})

	b.Run("Decryption", func(b *testing.B) {
		b.ResetTimer()
		res = spade.decrypt(dk, 1, ciphertexts)
	})

	if len(res) != len(dummyData[0]) {
		b.Errorf("Decrypt failed: invalid length of decrypted message slice")
	}
}
