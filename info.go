package main

import (
	"flag"
	"fmt"
)

func InfoHelp() error {
	fmt.Println("Usage: " + APP_NAME + " [--help] [--version]")
	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()

	return nil
}

func InfoVersion() error {
	fmt.Println("v0.0.0")

	return nil
}
