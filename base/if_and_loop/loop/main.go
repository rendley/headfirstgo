package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("before continute")
		break
		fmt.Println("after continue")
	}
	fmt.Println("AFTER LOOP")
}
