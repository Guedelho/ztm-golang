//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func totalCost(products [4]Product) {
	var (
		totalCost     float32
		totalProducts int
	)
	for i := 0; i < len(products); i++ {
		product := products[i]
		if product.name != "" {
			totalCost += product.price
			totalProducts += 1
		}
	}
	fmt.Println(products[totalProducts-1], totalProducts, totalCost)
}

func main() {
	shoppingList := [4]Product{
		{"shoes", 17.5},
		{"shirt", 19.98},
		{"short", 24.93},
	}
	totalCost(shoppingList)
	shoppingList[3] = Product{"tie", 9.87}
	totalCost(shoppingList)
}
