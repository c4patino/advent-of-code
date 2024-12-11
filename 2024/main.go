//go:generate go run ./cmd/update
package main

import (
	"log"

	"cpatino.com/advent-of-code/2024/cmd"
)

func main() {
	log.SetFlags(0)

	cmd.Execute()
}
