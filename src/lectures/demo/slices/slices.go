package main

import "fmt"

func printRoute(title string, route []string) {
	fmt.Println("---", title, "---")
	for i := 0; i < len(route); i++ {
		element := route[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printRoute("Route 1", route)

	route = append(route, "Home")
	printRoute("Route 2", route)

	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")

	route = route[2:]
	printRoute("Route 3", route)
}
