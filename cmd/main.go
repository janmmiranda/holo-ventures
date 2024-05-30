package main

import (
	"fmt"
	"os"

	"github.com/janmmiranda/holo-ventures/internal/tui"
)

func main() {
	_, err := tui.StartRepl()
	if err != nil {
		fmt.Print("Error occured %v", err)
		os.Exit(1)
	}
}
