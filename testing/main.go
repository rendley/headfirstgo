package main

import (
	"fmt"
	"headfirstgo/testing/prose"
	"strings"
)

func main() {
	date := []string{"05", "14", "2018"}
	fmt.Println(len(date))
	fmt.Println(prose.JoinWithCommas(date))
	//without func JoinWithCommas
	s := []string{"state", "of", "the", "art"}
	fmt.Println(strings.Join(s, "-"))

	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))

}
