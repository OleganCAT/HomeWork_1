package main

import (
	"fmt"
)

var (
	PriceApple  float64 = 5.99
	PricePear   int     = 7
	AmountMoney int     = 23
)

func main() {

	var purchases float64 = (PriceApple * float64(9)) + (float64(PricePear) * float64(8))
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? :", "Have to spand -", purchases, "UAH")
}
