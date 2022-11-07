package main

import "fmt"

type Preparer interface {
	PreparerDish()
}

type Chicken string
type Salad string

func (c Chicken) PreparerDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PreparerDish() {
	fmt.Println("cook salad")
}

func prepareDishes(dishes []Preparer) {
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		dish.PreparerDish()
	}
}

func main() {
	chicken := Chicken("Chicken")
	salad := Salad("Salad")
	dishes := []Preparer{chicken, salad}
	prepareDishes(dishes)
}
