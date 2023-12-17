package SleepSaFE

import (
	"SleepSaFE/utils"
	"crypto/rand"
	"math/big"
)

type Parameters struct {
	n int
	q big.Int
	g big.Int
}

func Setup(q *big.Int, n, m, t int) ([]*big.Int, []*big.Int, [][]*big.Int) {
	g := utils.RandomElementInZMod(q)
	sk := make([]*big.Int, n)
	pk := make([]*big.Int, n)
	alpha := make([][]*big.Int, m)

	for i := 0; i < n; i++ {
		sk[i] = utils.RandomElementInZMod(q)
		pk[i] = new(big.Int).Exp(g, sk[i], q)
	}

	for j := 0; j < m; j++ {
		alpha[j] = make([]*big.Int, n)
		for i := 0; i < n; i++ {
			alpha[j][i] = utils.RandomElementInZMod(q)
		}
	}

	return sk, pk, alpha
}

func encrypt(pk []*big.Int, alpha [][]*big.Int, Lx [][]int, j int, q *big.Int, g *big.Int) [][]*big.Int {
	x := Lx[j]
	c := make([][]*big.Int, len(x))
	a := alpha[j]

	for i := 0; i < len(x); i++ {
		ri, _ := rand.Int(rand.Reader, q)
		if ri.Bit(0) == 0 {
			ri.Add(ri, big.NewInt(1))
		}

		ci0 := new(big.Int).Exp(g, ri, q)
		ci1 := new(big.Int).Exp(pk[i], a[i], q)
		ci1.Mul(ci1, new(big.Int).Exp(new(big.Int).Exp(g, ri, q), new(big.Int).SetInt64(int64(x[i])), q))
		ci := []*big.Int{ci0, ci1}
		c[i] = ci
	}

	return c
}

func keyGen(sk []*big.Int, alpha [][]*big.Int, v, j int, q *big.Int, g *big.Int) []*big.Int {
	a := alpha[j]
	dk := make([]*big.Int, len(sk))
	for i := 0; i < len(sk); i++ {
		dk[i] = new(big.Int).Exp(g, new(big.Int).Mul(a[i], new(big.Int).Sub(new(big.Int).SetInt64(int64(v)), sk[i])), q)
	}
	return dk
}

func decrypt(dk []*big.Int, v int, c [][]*big.Int, q *big.Int) []*big.Int {
	y := make([]*big.Int, len(dk))
	for i := 0; i < len(dk); i++ {
		ci := c[i]
		yi := new(big.Int).Mul(new(big.Int).Exp(ci[1], new(big.Int).Neg(new(big.Int).SetInt64(int64(v))), q), dk[i])
		yi.Mod(yi, q)
		y[i] = yi
	}
	return y
}
