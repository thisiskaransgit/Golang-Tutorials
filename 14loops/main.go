package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// second syntax

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, days := range days {
		fmt.Printf("The index is %v and the days is %v \n", index, days)
	}

	roguevalue := 1

	for roguevalue < 10 {

		if roguevalue == 2 {
			goto lco
		}

		if roguevalue == 5 {
			roguevalue++
			// continue // skips value 	5
			break //breaks the loop
		}
		fmt.Println("Value is :", roguevalue)
		roguevalue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}
