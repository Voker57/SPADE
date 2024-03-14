package SPADE

import (
	"SPADE/utils"
	"crypto/rand"
	"math/big"
)

type SPADE struct {
	// n number of users
	n int
	// m maximum size of plaintext vector
	m int
	// q modulus
	q *big.Int
	// g generator
	g *big.Int
}

// NewSpade returns new instantiation of spade with "nil" values
func NewSpade() *SPADE {
	return &SPADE{
		n: 0,
		m: 0,
		q: nil,
		g: nil,
	}
}

// setup generates q and g based on plaintext vector size m and number of users n,
// then generates master public keys "pks" (encryption key) and master secret keys "sks",
// the number of keys for both "pks" and "sks" is bounded to the m, returns the "sks" and "pks"
func (spade *SPADE) setup(numUsers, ptVecSize int) ([]*big.Int, []*big.Int) {
	spade.n = numUsers
	spade.m = ptVecSize
	// q = (2 ^ 128) + 1
	spade.q = new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
	spade.q.Add(spade.q, big.NewInt(1))

	spade.g = spade.RandomElementInZMod()
	sks := make([]*big.Int, ptVecSize)
	pks := make([]*big.Int, ptVecSize)

	// generate secret and public keys
	for i := 0; i < ptVecSize; i++ {
		sks[i] = spade.RandomElementInZMod()
		pks[i] = new(big.Int).Exp(spade.g, sks[i], spade.q)
	}

	return sks, pks
}

// register accepts user's token "alpha" as input and generate user's registration key "regKey",
// which later on will be used by Curator for generating the decryption keys per query (check keyDerivation),
// returns "regKey"
func (spade *SPADE) register(alpha *big.Int) *big.Int {
	g := spade.g
	q := spade.q
	regKey := new(big.Int).Exp(g, alpha, q)
	return regKey
}

// encrypt encrypts a vector of integers "data" using master public key "pks" and user's "alpha",
// returns Elgamal style ciphertext vector c = [[C0, C1], ..., [C0, C1]]
func (spade *SPADE) encrypt(pks []*big.Int, alpha *big.Int, data []int) [][]*big.Int {
	q := spade.q
	g := spade.g

	dataSize := len(data)
	if dataSize != spade.m {
		panic("=== The input vector length does not matches the setup parameters!")
	}

	c := make([][]*big.Int, dataSize)

	for i := 0; i < dataSize; i++ {
		r := spade.RandomElementInZMod()
		// Ensure ri is odd
		if r.Bit(0) == 0 {
			r.Add(r, big.NewInt(1))
		}

		// cI0 = g^(r_i+alpha)
		cI0 := new(big.Int).Exp(g, new(big.Int).Add(r, alpha), q)
		// cI1 = (pk^alpha)*((g^r_i)^m_i)
		mI := new(big.Int).SetInt64(int64(data[i]))
		cI1 := new(big.Int).Mul(
			new(big.Int).Exp(pks[i], alpha, q),
			new(big.Int).Exp(new(big.Int).Exp(g, r, q), mI, q))
		// c_i = [cI0, cI1]
		c[i] = []*big.Int{cI0, cI1}
	}

	return c
}

// keyDerivation generates the decryption keys for specific query value "v",
// where the query is to be executed for a specific user "id",
// by using master secret key vector "sks", users' registration keys "regKeys"
// returns decryption keys "dk"
func (spade *SPADE) keyDerivation(id, value int, sks []*big.Int, regKeys []*big.Int) []*big.Int {
	q := spade.q
	//g := spade.g
	regKey := regKeys[id]

	dk := make([]*big.Int, spade.m)
	for i := 0; i < spade.m; i++ {
		//dk[i] = new(big.Int).Exp(g, new(big.Int).Mul(regKey, new(big.Int).Sub(new(big.Int).SetInt64(int64(value)), sks[i])), q)
		dk[i] = new(big.Int).Exp(regKey, new(big.Int).Sub(new(big.Int).SetInt64(int64(value)), sks[i]), q)
	}
	return dk
}

// decrypt decrypts the "ciphertexts" using decryption keys "dk" and value query "v",
// note: the value "v" must be the same value where the "dk" generated for (check keyDerivation)
func (spade *SPADE) decrypt(dk []*big.Int, value int, ciphertexts [][]*big.Int) []*big.Int {
	q := spade.q
	results := make([]*big.Int, spade.m)
	for i := 0; i < spade.m; i++ {
		ci := ciphertexts[i]
		vb := new(big.Int).Neg(new(big.Int).SetInt64(int64(value)))
		yi := new(big.Int).Mul(dk[i], new(big.Int).Mul(ci[1], new(big.Int).Exp(ci[0], vb, q)))
		yi.Mod(yi, q)
		results[i] = yi
	}
	return results
}

// RandomElementInZMod generates a random element from "Zq" correspond to q
func (spade *SPADE) RandomElementInZMod() *big.Int {
	r, err := rand.Int(rand.Reader, spade.q)
	utils.HandleError(err)
	return r
}
