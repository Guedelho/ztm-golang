package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(x, y int) int {
	return x + y
}

func greet() {
	fmt.Println("Hello from greet function!")
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println(dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println(bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println(anotherBakersDozen)
}
