//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a, b Coordinate
}

func calculateWidth(rectangle Rectangle) int {
	return rectangle.b.x - rectangle.a.x
}

func calculateLength(rectangle Rectangle) int {
	return rectangle.a.y - rectangle.b.y
}

func calculateArea(rectangle Rectangle) int {
	return calculateLength(rectangle) * calculateWidth(rectangle)
}

func calculatePerimeter(rectangle Rectangle) int {
	return (calculateLength(rectangle) * 2) + (calculateWidth(rectangle) * 2)
}

func main() {
	rectangle := Rectangle{Coordinate{0, 7}, Coordinate{10, 0}}
	fmt.Println(rectangle, calculateArea(rectangle), calculatePerimeter(rectangle))
	rectangle.a.y *= 2
	rectangle.b.x *= 2
	fmt.Println(rectangle, calculateArea(rectangle), calculatePerimeter(rectangle))
}
