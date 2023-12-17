package utils

import (
	"crypto/rand"
	"math/big"
)

// HandleError check the error if err then panic
func HandleError(e error) {
	if e != nil {
		panic(e)
	}
}

// RandomElementInZMod g from Zq
func RandomElementInZMod(q *big.Int) *big.Int {
	r, err := rand.Int(rand.Reader, q)
	HandleError(err)
	return r
}
