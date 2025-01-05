package main

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f lit", l)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mil", m)
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	fmt.Println(Gallons(12.998503))
	fmt.Println(Liters(12.43729847))
	fmt.Println(Milliliters(12.43729847))

	fmt.Printf("%0.2f gal\n", Gallons(12.80949385))
	fmt.Printf("%0.2f L\n", Liters(12.34890284))
	fmt.Printf("%0.2f M\n", Milliliters(12.38409433))

	coffeePot := CoffeePot("Luxbrew")
	fmt.Println(coffeePot)

}
