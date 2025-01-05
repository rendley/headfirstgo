package main

import (
	"fmt"
	"log"
)

//type error interface {
//	Error() string
//}

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	//err := fmt.Errorf("a heigh of %0.2f is invalid", -2.333)
	//fmt.Println(err.Error())
	//fmt.Println(err)

	//var err error
	//err = ComedyError("wha's a programmer's favorite beer? Logger!")
	//fmt.Println(err)

	var err error
	err = checkTemperature(121.34, 100)
	if err != nil {
		log.Fatal(err)
	}
}
