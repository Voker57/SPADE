package SPADE

import (
	"SPADE/utils"
	"fmt"
	"math/big"
	"testing"
)

func TestSpade(t *testing.T) {
	for _, tc := range TestVector {
		fmt.Println(TestString("SPADE", tc))
		testSpade(t, tc.n, tc.m, tc.l, tc.v)
	}
	//tc := TestVector[0]
	//fmt.Println(TestString("SPADE", tc))
	//testSpade(t, tc.n, tc.m, tc.l, tc.v)
}

func testSpade(t *testing.T, n int, m int, l int64, v int) {
	dummyData := utils.GenDummyData(n, m, l)

	spade := NewSpade()
	var sks, pks, dks, res []*big.Int
	var ciphertexts [][]*big.Int

	t.Run("Setup", func(t *testing.T) {
		sks, pks = spade.setup(n, m)
	})

	// create dummy registration keys
	alphas := make([]*big.Int, n)
	regKeys := make([]*big.Int, n)

	// to test one user registration
	t.Run("Register", func(t *testing.T) {
		alphas[0] = spade.RandomElementInZMod()
		regKeys[0] = spade.register(alphas[0])
	})

	// do the registration for the rest of users
	for i := 1; i < n; i++ {
		alphas[i] = spade.RandomElementInZMod()
		regKeys[i] = spade.register(alphas[i])
	}

	t.Run("Encryption", func(t *testing.T) {
		ciphertexts = spade.encrypt(pks, alphas[0], dummyData[0])
	})

	t.Run("keyDerivation", func(t *testing.T) {
		dks = spade.keyDerivation(0, v, sks, regKeys)
	})

	t.Run("Decryption", func(t *testing.T) {
		res = spade.decrypt(dks, v, ciphertexts)
	})

	if len(res) != len(dummyData[0]) {
		t.Errorf("Decrypt failed: invalid length of decrypted message slice")
	}

	//fmt.Println("data: ", dummyData[0])
	//fmt.Println("res: ", res)
	verifyResults(dummyData[0], res, v)
}

func verifyResults(originalData []int, res []*big.Int, v int) {
	nMatchEls := 0
	for i := 0; i < len(originalData); i++ {
		if originalData[i] == v {
			tmp := new(big.Int).SetInt64(int64(originalData[i]))
			if res[i].Cmp(tmp) != 0 {
				// the element from results vector is not equal to the one from original data
				// which means that we are not getting the correct results!!
				nMatchEls++
			}
		}
	}
	if nMatchEls != 0 {
		fmt.Println("=== FAIL: there are ", nMatchEls, " elements from the results vector, that are not equal to the original data!")
	} else {
		fmt.Println("=== PASS: Hooray!")
	}
}
