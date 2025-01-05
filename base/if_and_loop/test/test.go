package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\t fomerly surrounded by sapce \n"
	fmt.Println(strings.TrimSpace(s), s, s)
}
