package main

import (
	"fmt"
	"math/rand"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	fmt.Println(doorNumber)
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number!")
	}
}

func main() {
	awardPrize()
}
