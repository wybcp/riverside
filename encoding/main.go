package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonTest()
}

type address struct {
	Type    string
	City    string
	Country string
}

func jsonTest() {
	pa := address{"private", "cd", "cn"}
	js, _ := json.Marshal(pa)
	fmt.Printf("json foarmat is %s", js)
}
