package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := errors.New("height ca't be nagative")
	fmt.Println(err)
	log.Fatal(err)
}
