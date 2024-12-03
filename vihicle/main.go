package main

import "fmt"

type Vihicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up!")
}
func (c Car) Brake() {
	fmt.Println("Stopping!")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up!")
}
func (t Truck) Brake() {
	fmt.Println("Stopping!")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

func main() {
	var vihicle Vihicle = Car("Toyota Yarvic")
	vihicle.Accelerate()
	vihicle.Steer("left")
	vihicle.Brake()

	vihicle = Truck("Ford F180")
	vihicle.Brake()
	vihicle.Accelerate()
}
