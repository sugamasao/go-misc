package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	exe, err := os.Executable()
	path := filepath.Dir(exe)
	dir := filepath.Join(path, "data")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}
}
