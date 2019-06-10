package main

import (
	"fmt"
	"log"

	"github.com/NasSilverBullet/teamspirit/pkg/webdriver"
)

func main() {
	d, err := webdriver.NewWebDriver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
}
