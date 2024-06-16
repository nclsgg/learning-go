package main

import (
	"fmt"
	"math/cmplx"
	"strconv"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	test := strconv.FormatUint(MaxInt, 10)
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)
	fmt.Printf("Type: %T value: %v\n", test, test)
}
