package main

import "fmt"

type Space struct {
	occupied bool
}

type Parkinglot struct {
	spaces []Space
}

func occupySpace(lot *Parkinglot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *Parkinglot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *Parkinglot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := Parkinglot{spaces: make([]Space, 10)}
	fmt.Println(lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println(lot)

	lot.vacateSpace(2)
	fmt.Println(lot)
}
