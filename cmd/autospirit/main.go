package main

import (
	"log"
	"time"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	d, err := webdriver.NewWebDriver()
	if err != nil {
		return err
	}
	if err := d.Driver.Start(); err != nil {
		return err
	}
	defer func() {
		err = d.Driver.Stop()
	}()
	page, err := d.Driver.NewPage()
	if err != nil {
		return err
	}
	err = page.Navigate(d.RootURL)
	time.Sleep(3 * time.Second)
	return err
}
