package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	// No inheritance in golang; no super or parent

	karan := User{"Karan", "karan1234@go.dev", true, 18}
	fmt.Println(karan)
	fmt.Printf("The details of him are : %+v \n", karan)
	fmt.Printf("His name is %v and his email is %v \n", karan.Name, karan.Email)
	karan.GetStatus()
	karan.NewMail()
	fmt.Printf("His name is %v and his email is %v \n", karan.Name, karan.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

// Here the copy is being passed not the original value
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is :", u.Email)
}
