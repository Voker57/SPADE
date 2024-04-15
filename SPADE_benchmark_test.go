package SPADE

import (
	"SPADE/utils"
	"fmt"
	"math/big"
	"testing"
)

func BenchmarkSpade(b *testing.B) {
	for _, tc := range TestVector {
		fmt.Println(TestString("SPADE", tc))
		benchmarkSpade(b, tc.n, tc.m, tc.l, tc.v)
	}
}

func benchmarkSpade(b *testing.B, n int, m int, l int64, v int) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	dummyData := utils.GenDummyData(n, m, l)

	// generate q, q = (2 ^ 128) + 1
	q := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
	q.Add(q, big.NewInt(1))
	// generate g
	g := RandomElementInZMod(q)

	spade := NewSpade(q, g)
	var sks, pks, dks, res []*big.Int
	var ciphertexts [][]*big.Int

	b.Run("Setup", func(b *testing.B) {
		b.ResetTimer()
		sks, pks = spade.Setup(n, m)
	})

	// create dummy registration keys
	alphas := make([]*big.Int, n)
	regKeys := make([]*big.Int, n)

	b.Run("Register", func(b *testing.B) {
		for i := 0; i < n; i++ {
			alphas[i] = RandomElementInZMod(q)
			regKeys[i] = spade.Register(alphas[i])
		}
	})

	b.Run("Encryption", func(b *testing.B) {
		b.ResetTimer()
		ciphertexts = spade.Encrypt(pks, alphas[0], dummyData[0])
	})

	b.Run("KeyDerivation", func(b *testing.B) {
		b.ResetTimer()
		dks = spade.KeyDerivation(0, v, sks, regKeys)
	})

	b.Run("Decryption", func(b *testing.B) {
		b.ResetTimer()
		res = spade.Decrypt(dks, v, ciphertexts)
	})

	if len(res) != len(dummyData[0]) {
		b.Errorf("Decrypt failed: invalid length of decrypted message slice")
	}
}
