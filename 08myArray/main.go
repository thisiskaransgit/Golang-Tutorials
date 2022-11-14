package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	// fruitList[0]= "Apple"

	fmt.Println("The fruit list is : ", fruitList)
	fmt.Println("Length of the list is : ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushrooms"}
	fmt.Println("Length of the list is :", len(vegList))
	fmt.Println("The vegetable List is : ", vegList)
}
