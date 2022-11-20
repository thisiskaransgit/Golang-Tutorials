package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in Golang")

	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("The length is : ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNillErr(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
