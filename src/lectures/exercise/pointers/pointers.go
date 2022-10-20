//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag struct {
	status bool
}

func changeTagStatus(securityTag *SecurityTag, status bool) {
	securityTag.status = status
}

func checkout(tagList []SecurityTag) {
	for i := 0; i < len(tagList); i++ {
		changeTagStatus(&tagList[i], Inactive)
	}
}

func main() {
	tagList := []SecurityTag{{status: true}, {status: true}, {status: true}, {status: true}}
	fmt.Println(tagList)
	changeTagStatus(&tagList[0], Inactive)
	fmt.Println(tagList)
	checkout(tagList)
	fmt.Println(tagList)
}
