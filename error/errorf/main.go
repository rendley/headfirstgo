package main

import (
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("A height of %0.2f is invalid", -2.22)
	fmt.Println(err.Error())
	fmt.Println(err)
	log.Fatal(err)
}
