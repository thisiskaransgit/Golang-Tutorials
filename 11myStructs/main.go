package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance in golang; no super or parent

	hitesh := User{"Hitesh", "hitesh1234@go.dev", true, 18}
	fmt.Println(hitesh)
	fmt.Printf("The details of him are : %+v \n", hitesh)
	fmt.Printf("His name is %v and his email is %v \n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
