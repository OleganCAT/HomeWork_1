package main

import (
	"fmt"
)

var (
	priceApple  float64 = 5.99
	pricePear   int     = 7
	amountMoney int     = 23
)

func main() {

	appelsAndPears := (priceApple * float64(9)) + (float64(pricePear) * float64(8))
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? :", "Have to spand -", appelsAndPears, "UAH")

	numberOfPears := amountMoney / pricePear
	fmt.Println("2. Скільки груш ми можемо купити? :", "We can buy -", numberOfPears, "Pears")

}
