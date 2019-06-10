package main

import (
	"fmt"
	"log"

	"github.com/NasSilverBullet/autospirit/pkg/webdriver"
)

func main() {
	d, err := webdriver.NewWebDriver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
}
