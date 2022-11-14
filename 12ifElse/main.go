package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		fmt.Println("exactly 10 login count")
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("The number is Even0")
	} else {
		fmt.Println("The number is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not less than 10")
	}

}
