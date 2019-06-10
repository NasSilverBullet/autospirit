package main

import (
	"fmt"
	"log"

	"github.com/NasSilverBullet/auto-insert_timestamps-on-teamsprit/pkg/webdriver"
)

func main() {
	fmt.Println("timestamps")
	d, err := webdriver.NewWebDriver()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
}
