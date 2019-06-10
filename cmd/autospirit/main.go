package main

import (
	"log"

	"github.com/NasSilverBullet/autospirit/internal/cmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := cmd.NewAutoSpiritCommand()
	err := c.Execute()
	return err
}
