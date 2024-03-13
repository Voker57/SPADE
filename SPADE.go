package SPADE

import (
	"SPADE/utils"
	"crypto/rand"
	"math/big"
)

// SPADE
// n: number of users
// q: size of plaintext vector
// q: modulus
// g: group generator
type SPADE struct {
	numUser   int
	ptVecSize int
	q         *big.Int
	g         *big.Int
}

func NewSpade() *SPADE {
	return &SPADE{
		numUser:   0,
		ptVecSize: 0,
		q:         nil,
		g:         nil,
	}
}

// setup
func (spade *SPADE) setup(ptVecSize, numUsers int) ([]*big.Int, []*big.Int) {
	spade.ptVecSize = ptVecSize
	spade.numUser = numUsers
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

// register
func (spade *SPADE) register(alpha *big.Int) *big.Int {
	g := spade.g
	q := spade.q
	regKey := new(big.Int).Exp(g, alpha, q)
	return regKey
}

// encrypt encrypts a vector of integers
// pk: public encryption key
// alpha: user's secret key
// data: a vector of integers
func (spade *SPADE) encrypt(pk []*big.Int, alpha *big.Int, data []int) [][]*big.Int {
	q := spade.q
	g := spade.g

	dataSize := len(data)
	if dataSize != spade.ptVecSize {
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
			new(big.Int).Exp(pk[i], alpha, q),
			new(big.Int).Exp(new(big.Int).Exp(g, r, q), mI, q))
		// c_i = [cI0, cI1]
		c[i] = []*big.Int{cI0, cI1}
	}

	return c
}

// keygen generates the decryption keys for specific query value
// sk: master secret key vector
// regKeys: users' registration keys
// value: query value from analyst
// id: user's id
func (spade *SPADE) keygen(id, value int, sk []*big.Int, regKeys []*big.Int) []*big.Int {
	q := spade.q
	g := spade.g
	a := regKeys[id]

	dk := make([]*big.Int, spade.ptVecSize)
	for i := 0; i < spade.ptVecSize; i++ {
		dk[i] = new(big.Int).Exp(g, new(big.Int).Mul(a, new(big.Int).Sub(new(big.Int).SetInt64(int64(value)), sk[i])), q)
	}
	return dk
}

// decrypt decrypts the ciphertext for specific query value
// dk: decryption keys for query value
// value: query value from analyst
// ciphertexts: users' ciphertexts
func (spade *SPADE) decrypt(dk []*big.Int, value int, ciphertexts [][]*big.Int) []*big.Int {
	q := spade.q
	results := make([]*big.Int, spade.ptVecSize)
	for i := 0; i < spade.ptVecSize; i++ {
		ci := ciphertexts[i]
		vb := new(big.Int).Neg(new(big.Int).SetInt64(int64(value)))
		yi := new(big.Int).Mul(dk[i], new(big.Int).Mul(ci[1], new(big.Int).Exp(ci[0], vb, q)))
		yi.Mod(yi, q)
		results[i] = yi
	}
	return results
}

// RandomElementInZMod generates a random element from Zq with respect to the modulus q
func (spade *SPADE) RandomElementInZMod() *big.Int {
	r, err := rand.Int(rand.Reader, spade.q)
	utils.HandleError(err)
	return r
}
