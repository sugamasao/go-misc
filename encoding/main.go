package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func eucToUTF8(text string) (string, error) {
	buf := new(bytes.Buffer)
	r := transform.NewReader(strings.NewReader(text), japanese.EUCJP.NewDecoder())
	io.Copy(buf, r)
	return buf.String(), nil
}

func readFile(path string) string {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%s open Failed %v", path, err)
	}
	return string(str)
}

const eucPath = "./testdata/euc.txt"

func main() {
	ret, err := eucToUTF8(readFile(eucPath))
	if err != nil {
		log.Fatalf("eucToUTF8 Failed %v", err)
	}
	log.Println(ret)
}
