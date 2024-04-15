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

	//To manually select test case
	//tc := TestVector[0]
	//fmt.Println(TestString("SPADE", tc))
	//testSpade(t, tc.n, tc.m, tc.l, tc.v)
}

func testSpade(t *testing.T, n int, m int, l int64, v int) {
	dummyData := utils.GenDummyData(n, m, l)
	// generate q, q = (2 ^ 128) + 1
	q := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
	q.Add(q, big.NewInt(1))
	// generate g
	g := RandomElementInZMod(q)

	spade := NewSpade(q, g)
	var sks, pks, dks, res []*big.Int
	var ciphertexts [][]*big.Int

	t.Run("Setup", func(t *testing.T) {
		sks, pks = spade.Setup(n, m)
	})

	// initialize registration keys
	alphas := make([]*big.Int, n)
	regKeys := make([]*big.Int, n)

	// test only one user registration
	t.Run("Register", func(t *testing.T) {
		alphas[0] = RandomElementInZMod(q)
		regKeys[0] = spade.Register(alphas[0])
	})

	// do the registration for the rest of users
	for i := 1; i < n; i++ {
		alphas[i] = RandomElementInZMod(q)
		regKeys[i] = spade.Register(alphas[i])
	}

	t.Run("Encryption", func(t *testing.T) {
		ciphertexts = spade.Encrypt(pks, alphas[0], dummyData[0])
	})

	t.Run("KeyDerivation", func(t *testing.T) {
		dks = spade.KeyDerivation(0, v, sks, regKeys)
	})

	t.Run("Decryption", func(t *testing.T) {
		res = spade.Decrypt(dks, v, ciphertexts)
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
			// i: the index for match query value in the original dataset
			// check to see if the decrypted value from results for the same
			// index i is 1 or not, 1 means that the query value match there.
			one := new(big.Int).SetInt64(int64(1))
			if res[i].Cmp(one) != 0 {
				// the index from result vector does not match with the index for original dataset,
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
