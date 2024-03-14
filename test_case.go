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
	// n number of users
	n int
	// m plaintext vector size
	m int
	// l maximum value for data range
	l int64
	// query value for decryption key generation
	v int
}

var TestVector = []TestContext{
	{
		tc: S,
		n:  10,
		m:  100,
		l:  10,
		v:  1,
	},
	{
		tc: M,
		n:  100,
		m:  10000,
		l:  10,
		v:  1,
	},
	{
		tc: L,
		n:  1000,
		m:  100000,
		l:  10,
		v:  1,
	},
}

func TestString(name string, t TestContext) string {
	return fmt.Sprintf("=== %s: NumberOfUsers=%d, DataSize=%d, MaxDataValue=%d, QueryValue=%d",
		name, t.n, t.m, t.l, t.v)
}
