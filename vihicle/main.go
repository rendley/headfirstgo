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

func TryVehicle(vehicle Vihicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("format cargo")
	}
}

func main() {
	TryVehicle(Truck("Fnord F180"))
	TryVehicle(Car("Toyota"))
}
