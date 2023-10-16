package main

import "math/big"

func Fib(num int) string {
	bigNum := big.NewInt(int64(num))
	bigA1 := big.NewInt(1)
	bigA2 := big.NewInt(1)
	return recursive(bigNum, bigA1, bigA2).String()
}

func recursive(num *big.Int, a1 *big.Int, a2 *big.Int) *big.Int {
	if num.Cmp(big.NewInt(2)) <= 0 {
		return a2
	}
	a1.Add(a1, a2)
	num.Sub(num, big.NewInt(1))

	return recursive(num, a2, a1)
}
