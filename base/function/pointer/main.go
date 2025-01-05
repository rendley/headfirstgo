package main

import "fmt"

func createPointer() *float64 {
	var myFloat = 98.4
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(number *int) {
	*number *= 2
}

func main() {
	var myInt int
	myInt = 42
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
	fmt.Println("####################")
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
	fmt.Println("####################")

	var myBool bool = true
	printPointer(&myBool)

	amount := 6
	double(&amount)
	fmt.Println(amount)

}
