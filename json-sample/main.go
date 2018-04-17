package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {
	raw, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		log.Fatal(err)
	}
	// json to go https://mholt.github.io/json-to-go/
	type Animal struct {
		Name  string `json:"Name"`
		Order string `json:"Order"`
	}
	var animals []Animal
	json.Unmarshal(raw, &animals)

	for i, v := range animals {
		log.Printf("[%d] %s -> %s", i, v.Name, v.Order)
	}
}
