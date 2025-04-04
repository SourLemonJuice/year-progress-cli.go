package info

import (
	"flag"
	"fmt"

	"github.com/SourLemonJuice/year-progress-cli.go/define"
)

func InfoHelp() error {
	fmt.Printf("Usage: " + define.APP_NAME + " [--help] [--version] [--width]\n")
	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	return nil
}

func InfoVersion() error {
	fmt.Println("v0.0.0")

	return nil
}
