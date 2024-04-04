package usecases

import (
	"SPADE"
	"math/big"
)

type CSP struct {
	sks   []*big.Int
	pks   []*big.Int
	spade *SPADE.SPADE
}

func (csp CSP) Setup(numUser, maxVecSize int) {
	spd := SPADE.NewSpade()
	csp.sks, csp.pks = spd.Setup(numUser, maxVecSize)
	csp.spade = spd
}
