package main

import "fmt"

var packageVar = "package"

func main() {
	var functionVar = "function"
	if true {
		var conditionalVal = "conditional"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVal)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	//fmt.Println(conditionalVar) - not for scope main
}
