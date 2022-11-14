package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Class on Pointers")

	var ptr *int
	fmt.Println("The value is : ", ptr)

	myNumber := 23

	var ptr2 = &myNumber
	fmt.Println("The value is : ", *ptr2)

	*ptr2 += 2
	fmt.Println("The new value is :  ", myNumber)

}
