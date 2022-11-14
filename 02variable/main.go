package main

import "fmt"

const LoginToken string = "gibbrish"

func main() {
	var username string = "Karan"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var SmallValue uint8 = 255
	fmt.Println(SmallValue)
	fmt.Printf("variable is of type: %T \n", SmallValue)

	var SmallFloat float32 = 255.345678
	fmt.Println(SmallFloat)
	fmt.Printf("variable is of type: %T \n", SmallFloat)

	var BigFloat float64 = 255.345678234567
	fmt.Println(BigFloat)
	fmt.Printf("variable is of type: %T \n", BigFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var AnotherVariable string
	fmt.Println(AnotherVariable)
	fmt.Printf("Variable is of type: %T \n", AnotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	/*
		You can use the this syntax for decalaration within a method to declare local variables
		But you cannot use them to declare the global variables
	*/
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
