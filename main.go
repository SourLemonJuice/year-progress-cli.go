package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const APP_NAME string = "year-progress.go"
const PROGRESS_DEFAULT_WIDTH int = 80 * 0.6

func main() {
	var err error

	helpFlag := flag.Bool("help", false, "Print help")
	versionFlag := flag.Bool("version", false, "Print Version")
	widthFlag := flag.Int("width", PROGRESS_DEFAULT_WIDTH, "Width")

	flag.Parse()

	switch {
	case *helpFlag:
		InfoHelp()
		os.Exit(0)
	case *versionFlag:
		InfoVersion()
		os.Exit(0)
	}

	meta, err := PrintMetaNew(time.Now(), *widthFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	meta.PrintSection1()
	meta.PrintSectionProgress()
}
