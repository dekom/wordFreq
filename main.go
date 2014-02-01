package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		return
	}

	Run(args[1])
}
