package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(2)
	b := big.NewInt(20)

	sum := new(big.Int).Add(a, b)
	difference := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Div(a, b)

	fmt.Printf("Сумма: %s\n", sum.String())
	fmt.Printf("Разность: %s\n", difference.String())
	fmt.Printf("Произведение: %s\n", product.String())
	fmt.Printf("Частное: %s\n", quotient.String())
}
