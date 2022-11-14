package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in Golang")
	var mynumberOne int = 2
	var mynumberTwo float64 = 4.6

	// fmt.Println("The sum is : ", mynumberOne+mynumberTwo)
	fmt.Println("The sum is : ", mynumberOne+int(mynumberTwo))

	// random number
	/*
		rand.Seed(time.Now().UnixNano())
		fmt.Println("the number is ", rand.Intn(5)+1) // You will get the same number
	*/

	// random with crypto

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)

}
