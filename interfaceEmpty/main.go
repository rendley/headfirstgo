package main

import (
	"fmt"
	"headfirstgo/mypkg"
)

type Anything interface {
}

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(mypkg.Whistle)
	if ok {
		whistle.MakeSound()
	}
}

func main() {
	fmt.Println(2.333, "A string", true)

	AcceptAnything(2.343)
	AcceptAnything("a string")
	AcceptAnything(true)
	AcceptAnything(mypkg.Whistle("Too canary"))
}
