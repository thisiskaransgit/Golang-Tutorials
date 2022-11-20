package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `josn:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Video")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnlin.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "	LearnCodeOnlin.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnlin.in", "hit123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "*", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`
		{
	        "coursename": "ReactJs Bootcamp",
            "Price": 299,
            "Platform": "LearnCodeOnlin.in",
            "tags": ["web-dev","js"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(JsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(JsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and  value is %v and type is : %T \n", k, v, v)
	}
}
