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

type Country2 map[string]Person

func (c *Country) findPerson(fieldName string) Person {
	pC := reflect.ValueOf(*c)
	f := pC.FieldByName(fieldName)
	if !f.IsValid() {
		log.Fatalf("Notfound Field Name. [%s]", fieldName)
	}

	return f.Interface().(Person)
}

func useStruct(raw []byte) Person {
	var c Country
	log.SetPrefix("[Struct]")
	if err := yaml.Unmarshal(raw, &c); err != nil {
		log.Fatalf("Failed Unmarshal[%v]", err)
	}

	// log.Printf("Japan -> %v", c.Japan)
	// log.Printf("USA   -> %v", c.USA)

	fieldName := "Japan" // ここを動的に変更したい
	return c.findPerson(fieldName)
	// log.Printf("%s -> %v", fieldName, person)
}

func useInterface(raw []byte) Person {
	c := make(map[interface{}]Person)
	log.SetPrefix("[Interface]")
	if err := yaml.Unmarshal(raw, &c); err != nil {
		log.Fatalf("Failed Unmarshal[%v]", err)
	}

	fieldName := "japan"
	// log.Printf("Japan -> %v", c[fieldName])
	// log.Printf("UK    -> %v", c["UK"].Name)
	// if (Person{}) == c["hoge"] {
	// 	log.Printf("This Stcut is nil")
	// }
	return c[fieldName]
}

func useStringField(raw []byte) Person {
	var c Country2
	log.SetPrefix("[string]")
	if err := yaml.Unmarshal(raw, &c); err != nil {
		log.Fatalf("Failed Unmarshal[%v]", err)
	}

	fieldName := "japan"
	// log.Printf("Japan -> %v", c[fieldName])
	// log.Printf("UK    -> %v", c["UK"].Name)
	// if (Person{}) == c["hoge"] {
	// 	log.Printf("This Stcut is nil")
	// }
	return c[fieldName]
}

func main() {
	raw, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		log.Fatal(err)
	}
	useStruct(raw)
	useInterface(raw)
	useStringField(raw)
}
