/*
	Author: Charles Shook
	Description: Showing off iteration in Go!
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Iteration with a for loop")

	for i:= 0; i <= 20; i++ {
		fmt.Println(i)
	}

	fmt.Println("Iteration with a range for loop")
	var myarray [5]int
	myarray[0] = 20
	myarray[1] = 30
	myarray[2] = 40
	myarray[3] = 50
	myarray[4] = 60

	for _, item := range myarray {
		fmt.Println(item)
	}
}