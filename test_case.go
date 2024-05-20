package SPADE

import "fmt"

type TestCase int

const (
	S TestCase = iota
	M
	L
)

type TestContext struct {
	// tc test case
	tc TestCase
	// m number of users
	m int
	// n maximum plaintext vector size
	n int
	// l maximum value for data range
	l int64
	// query value for decryption key generation
	v int
}

var TestVector = []TestContext{
	{
		tc: S,
		m:  10,
		n:  1000,
		l:  10,
		v:  1,
	},
	{
		tc: M,
		m:  100,
		n:  10000,
		l:  10,
		v:  1,
	},
	{
		tc: L,
		m:  1000,
		n:  50000,
		l:  10,
		v:  1,
	},
	{
		tc: L,
		m:  10000,
		n:  100000,
		l:  10,
		v:  1,
	},
}

func TestString(name string, t TestContext) string {
	return fmt.Sprintf("=== %s: NumberOfUsers=%d, DataSize=%d, MaxDataValue=%d, QueryValue=%d",
		name, t.m, t.n, t.l, t.v)
}
