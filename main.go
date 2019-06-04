package main

import (
	"fmt"
	"github.com/brsgr/chip8-go/chip8"
	"log"
)

func main() {
	chip, err := chip8.NewCPU(60)
	if err != nil {
		log.Fatalln("failed to init CPU")
	}
	fmt.Println(chip)

	chip.Run()
}
