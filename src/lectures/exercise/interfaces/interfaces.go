//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Motorcycle string
type Car string
type Truck string

type Lifter interface {
	Lift() int
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("%v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("%v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("%v", string(t))
}

func (m Motorcycle) Lift() int {
	return SmallLift
}

func (c Car) Lift() int {
	return StandardLift
}

func (t Truck) Lift() int {
	return LargeLift
}

func liftVehicle(l Lifter) {
	switch l.Lift() {
	case SmallLift:
		fmt.Printf("%v small\n", l)
	case StandardLift:
		fmt.Printf("%v standard\n", l)
	case LargeLift:
		fmt.Printf("%v large\n", l)
	}
}

func main() {
	moto := Motorcycle("Moto")
	car := Car("Car")
	truck := Truck("Truck")
	liftVehicle(moto)
	liftVehicle(car)
	liftVehicle(truck)
}
