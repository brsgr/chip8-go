package chip8

import (
	"log"
	"time"
)

type CPU struct {
	// https://en.wikipedia.org/wiki/CHIP-8#Virtual_machine_description

	// memory
	// 0x000-0x1FF - Chip 8 interpreter (contains font set in emu)
	// 0x050-0x0A0 - Used for the built in 4x5 pixel font set (0-F)
	// 0x200-0xFFF - Program ROM and work RAM
	Memory [4096]byte

	// data Registers
	V [16]byte

	// Address Register
	I uint16

	// opcode
	// https://en.wikipedia.org/wiki/CHIP-8#Opcode_table
	op uint16

	// program counter
	pc uint16

	// display is 2048 pixles monochrome
	display [64][32]bool

	// timers
	DelayTimer int
	SoundTimer int

	// stack
	// https://en.wikipedia.org/wiki/CHIP-8#The_stack
	Stack [16]byte

	// stop channel stops the program
	Stop chan bool

	// clock pulls from a `tick` channel
	Clock <-chan time.Time
}

func NewCPU(ClockSpeed time.Duration) (*CPU, error) {

	c := &CPU{
		Clock: time.NewTicker(time.Second / ClockSpeed).C,
		Stop:  make(chan bool),
	}

	return c, nil
}

func (c *CPU) Run() error {
	for {
		select {
		case <-c.Stop:
			log.Println("stopping emulator")
		case e := <-c.Clock:
			log.Println(e)
			// run our cool stuff here
		}
	}
}
