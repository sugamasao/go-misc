package main

import (
	"io/ioutil"
	"log"
	"reflect"

	yaml "gopkg.in/yaml.v2"
)

type Person struct {
	Name string
	Age  int
}

type Country struct {
	Japan Person
	USA   Person
}

func (c *Country) findPerson(fieldName string) Person {
	pC := reflect.ValueOf(*c)
	f := pC.FieldByName(fieldName)
	if !f.IsValid() {
		log.Fatalf("Notfound Field Name. [%s]", fieldName)
	}

	return f.Interface().(Person)
}

func main() {
	raw, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var c Country
	if err := yaml.Unmarshal(raw, &c); err != nil {
		log.Fatalf("Failed Unmarshal[%v]", err)
	}

	log.Printf("Japan -> %v", c.Japan)
	log.Printf("USA   -> %v", c.USA)

	fieldName := "Japan" // ここを動的に変更したい
	person := c.findPerson(fieldName)
	log.Printf("%s -> %v", fieldName, person)
}
