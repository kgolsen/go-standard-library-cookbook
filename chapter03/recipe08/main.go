package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func main() {

	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))
	for i := 0; i < 5; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Println("math/rand generated non-equal sequence")
		} else {
			fmt.Printf("math/rand 1: %d, math/rand 2: %d\n", rnd1, rnd2)
		}
	}

	for i := 0; i < 5; i++ {
		safeNum1 := NewCryptoRand()
		safeNum2 := NewCryptoRand()
		if safeNum1 == safeNum2 {
			fmt.Println("crypto/rand generated equal numbers")
			break
		} else {
			fmt.Printf("crypto/rand 1: %d, crypto/rand 2: %d\n", safeNum1, safeNum2)
		}
	}

}

func NewCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
