package main

import "math/big"

func Fib(num int) string {
	a1 := big.NewInt(0)
	a2 := big.NewInt(1)

	for i := 0; i < num; i++ {
		tmp := big.NewInt(0).Add(a1, a2)
		a1 = a2
		a2 = tmp
	}
	return a1.String()
}
