package main

import (
	"fmt"
)

const (
	priceApple  float64 = 5.99
	pricePear   int     = 7
	amountMoney int     = 23
)

func main() {

	appelsAndPears := (priceApple * float64(9)) + (float64(pricePear) * float64(8))
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? :", "Have to spand -", appelsAndPears, "UAH")

	amountOfPears := amountMoney / pricePear
	fmt.Println("2. Скільки груш ми можемо купити? :", "We can buy -", amountOfPears, "Pears")

	amountOfApples := float64(amountMoney) / priceApple
	fmt.Println("3. Скільки яблук ми можемо купити? :", "We can buy -", amountOfApples, "Apples")

	sumAppelsAndPears := (priceApple*2)+(float64(pricePear)*2) <= float64(amountMoney)
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? :", sumAppelsAndPears, "-(We can't pay two apples and two pears)")
}
