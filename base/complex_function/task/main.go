package main

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}
func callWithArgumetns(passedFunction func(a string, b bool)) {
	passedFunction("This sentence is ", false)
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArgumetns(functionC)
	printReturnValue(functionB)

}
