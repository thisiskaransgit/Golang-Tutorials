package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all the languages : ", languages)
	fmt.Println("JS is short for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the languages : ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v the value is %v \n", key, value)
	}
}
