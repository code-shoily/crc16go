package main

import (
	"fmt"

	"github.com/code-shoily/crc16go/lib"
)

func main() {
	bytes := []uint16{0x11, 0x31, 0x55, 0x17}
	fmt.Printf("%x\n", lib.Compute(bytes))
}
