package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	Greeting()

	result := adder(3, 5)
	fmt.Println("The result is : ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 64, 4, 12, -5)
	fmt.Println("Pro result is : ", proRes)
	fmt.Println("Pro Messsge is : ", myMessage)

}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, " Hello My Pro Result"
}

func Greeting() {
	fmt.Println("Namastey from Golang")
}
