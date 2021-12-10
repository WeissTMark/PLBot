/*
	Author: Charles Shook
	Description: Showing off input and output in Go!
*/

package main

import (
	"fmt"
)

func main() {
	var name string;

	fmt.Print("Enter your name :> ")
	fmt.Scanf("%s", &name)

	fmt.Println("Your name is: ", name)
}