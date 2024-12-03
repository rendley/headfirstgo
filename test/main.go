package main

import (
	"fmt"
	"headfirst/gadget"
	"headfirst/mypkg"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func play(n mypkg.NoiseMaker) {
	n.MakeSound()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapeRecorder{}
	playList(player, mixtape)
	fmt.Println("######################")

	//var value mypkg.MyInterface
	//value = mypkg.MyType(5)
	//value.MethodWithoutParameters()
	//value.MethodWithParameter(5.55)
	//fmt.Println(value.MethodWithReturnValue())
	//
	//var toy mypkg.NoiseMaker
	//toy = mypkg.Whistle("Toyco Canary")
	//toy.MakeSound()
	//toy = mypkg.Horn("Toyco Blaster")
	//toy.MakeSound()

}
