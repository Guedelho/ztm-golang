package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Println("Running computation")
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2+2=", compute(2, 2, add))
	fmt.Println("10-2=", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))
	mul := func(lhs, rhs int) int {
		fmt.Println("Multiplyting")
		return lhs * rhs
	}
	fmt.Println(compute(3, 3, mul))
}
