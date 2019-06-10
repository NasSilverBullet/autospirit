package main

import (
	"log"
	"time"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
)

func main() {
	d, err := webdriver.NewWebDriver()
	if err != nil {
		log.Fatal(err)
	}
	if err := d.Driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer func() {
		err = d.Driver.Stop()
	}()
	page, err := d.Driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	if err := page.Navigate(d.RootURL); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	time.Sleep(3 * time.Second)
}
