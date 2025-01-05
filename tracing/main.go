package main

import "fmt"

func main() {
	one()
}

func one() {
	defer fmt.Println("deferring func one()")
	two()
}

func two() {
	defer fmt.Println("deferring func two()")
	three()
}

func three() {
	defer fmt.Println("I'm with defer println1")
	defer fmt.Println("I'm with defer println2")
	fmt.Println("Hello i'm func three!")
	panic("This call stack's too deep for me!")
}
