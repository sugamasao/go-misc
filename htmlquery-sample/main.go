package main

import (
	"fmt"
	"os"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func toFile() *html.Node {
	f, err := os.Open("index.html")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	doc, err := htmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	d := doc
	return d
}

func main() {
	doc := toFile()

	for i, n := range htmlquery.Find(doc, "//p") {
		for _, a := range htmlquery.Find(n, "//a") {
			fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
		}
	}
}
