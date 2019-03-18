package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func setup() []byte {
	raw, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return raw
}

func BenchmarkUseStruct(b *testing.B) {
	raw := setup()
	for i := 0; i < b.N; i++ {
		useStruct(raw)
	}
}

func BenchmarkUseInterface(b *testing.B) {
	raw := setup()
	for i := 0; i < b.N; i++ {
		useInterface(raw)
	}
}

func BenchmarkUseString(b *testing.B) {
	raw := setup()
	for i := 0; i < b.N; i++ {
		useStringField(raw)
	}
}
