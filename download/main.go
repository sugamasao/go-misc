package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch() (string, error) {
	var url string = "http://example.com/index.html"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	ret, _ := Fetch()
	fmt.Printf("ret => %v", ret)
}
