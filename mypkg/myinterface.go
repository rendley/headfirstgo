package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}
func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}

type NoiseMaker interface {
	MakeSound()
}
type Whistle string
type Horn string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}
