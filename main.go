package main

import (
	"github.com/brsgr/chip8-go/chip8"
	"log"
	"time"
)

var (
	ClockSpeed = time.Duration(60) // default 60hz
)

func main() {
	chip, err := chip8.NewCPU(60)
	if err != nil {
		log.Fatalln("failed to init CPU")
	}

	chip.Run()
}
