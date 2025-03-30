package main

import (
	"fmt"
	"os"
	"time"
)

const APP_NAME string = "year-progress.go"
const PROGRESS_DEFAULT_WIDTH int = 80 * 0.6

func main() {
	var err error
	meta, err := PrintMetaNew(time.Now(), PROGRESS_DEFAULT_WIDTH)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	meta.PrintSection1()
	meta.PrintSectionProgress()
}
