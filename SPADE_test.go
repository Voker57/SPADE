package SPADE

import (
	"math/big"
	"testing"
)

func TestSpade(t *testing.T) {
	q, _ := new(big.Int).SetString("340282366920938463463374607431768211458", 10)
	g, _ := new(big.Int).SetString("167784994160571689805255858220993457294", 10)
	n := 1

	spade := NewSpade(q, g, n)

	// manual setup from set 'random' value for sk
	sk0, _ := new(big.Int).SetString("157222500695008773376422121395749431412", 10)
	sks := []*big.Int{sk0}
	pks := []*big.Int{new(big.Int).Exp(spade.g, sk0, spade.q)}

	alpha := new(big.Int).SetInt64(19) // set 'random' alpha
	regKey := spade.Register(alpha)

	given_value := 5
	fakeR := new(big.Int).SetInt64(1)

	// Encrypt is modified to inject only 'random' fakeR instead of r_i
	ciphertexts := spade.Encrypt(pks, alpha, []int{given_value}, fakeR)

	value := 5

	dks := spade.KeyDerivation(0, value, sks, regKey)

	spade.Decrypt(dks, value, ciphertexts)
}

/*
--- FAIL: TestSpade (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x508097]

goroutine 17 [running]:
testing.tRunner.func1.2({0x537260, 0x666880})
	/usr/lib/go-1.22/src/testing/testing.go:1631 +0x24a
testing.tRunner.func1()
	/usr/lib/go-1.22/src/testing/testing.go:1634 +0x377
panic({0x537260?, 0x666880?})
	/usr/lib/go-1.22/src/runtime/panic.go:770 +0x132
math/big.(*Int).Mul(0xc0001cdd88, 0xc0001a40c0, 0x0)
	/usr/lib/go-1.22/src/math/big/int.go:194 +0x97
SPADE.(*SPADE).Decrypt(0xc0001cdf08, {0xc000182030, 0x1, 0xc0001cdee0?}, 0x5, {0xc0001a20d8, 0x1, 0x5b6f80?})
	/home/voker57/Sources/SPADE/SPADE.go:112 +0x359
SPADE.TestSpade(0xc0001ac4e0?)
	/home/voker57/Sources/SPADE/SPADE_test.go:33 +0x292
testing.tRunner(0xc0001ac4e0, 0x56ab00)
	/usr/lib/go-1.22/src/testing/testing.go:1689 +0xfb
created by testing.(*T).Run in goroutine 1
	/usr/lib/go-1.22/src/testing/testing.go:1742 +0x390
FAIL	SPADE	0.004s
FAIL
*/
