package SleepSaFE

import (
	"SleepSaFE/utils"
	"math/big"
	"math/rand"
	"testing"
)

func TestSleepSaFE(t *testing.T) {
	q := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
	q.Add(q, big.NewInt(1))
	G := q

	var sk, pk, dk, res []*big.Int
	var alpha, c [][]*big.Int

	t.Run("Setup", func(t *testing.T) {
		sk, pk, alpha = Setup(G, 100, 50, 10)
	})

	Lx := make([][]int, 50)
	for j := 0; j < 50; j++ {
		Lx[j] = make([]int, 100)
		for i := 0; i < 100; i++ {
			Lx[j][i] = int(rand.Int63n(10) + 1)
		}
	}

	t.Run("Encryption", func(t *testing.T) {
		c = encrypt(pk, alpha, Lx, 0, G, utils.RandomElementInZMod(G))
	})

	t.Run("FEKeyGen", func(t *testing.T) {
		dk = keyGen(sk, alpha, 1, 0, G, utils.RandomElementInZMod(G))
	})

	t.Run("Decryption", func(t *testing.T) {
		res = decrypt(dk, 1, c, G)
	})

	if len(res) != len(Lx[0]) {
		t.Errorf("Decrypt failed: invalid length of decrypted message slice")
	}
}
