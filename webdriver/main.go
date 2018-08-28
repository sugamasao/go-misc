package main

import (
	"log"
	"os"

	"github.com/sclevine/agouti"
)

func main() {
	log.Println("start!")
	// Chromeを利用することを宣言
	agoutiDriver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		}),
		agouti.Debug,
	)
	log.Println("start!!")
	agoutiDriver.Start()
	log.Println("start!!!")
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()
	log.Println("start!!!!")
	if err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
	log.Fatalf("page %v", page)

	// 自動操作
	err = page.Navigate("https://qiita.com/")
	if err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
	log.Println(page.Title())
	page.Screenshot("Screenshot01.png")

	// err := page.FindByLink("Google 検索").Click()
	// log.Println(page.Title())
	// page.Screenshot("Screenshot02.png")
	log.Println("end")
}
