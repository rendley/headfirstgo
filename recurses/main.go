package main

import "fmt"

func recurses() {
	fmt.Println("Oh, no, I'm stuck!")
	recurses()
}

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
	fmt.Printf("Returning form count(%d, %d) call\n", start, end)
}

func main() {
	//recurses()
	count(1, 5)
}
