package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("hello!")
	return fmt.Errorf("I don't want to talk")
	fmt.Println("Nice weather, eh?")
	return nil

}

func main() {
	//err := Socialize()
	//if err != nil {
	//	log.Fatal(err)
	//}

	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
