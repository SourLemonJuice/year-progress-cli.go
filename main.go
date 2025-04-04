package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/SourLemonJuice/year-progress-cli.go/define"
	"github.com/SourLemonJuice/year-progress-cli.go/info"
	"github.com/SourLemonJuice/year-progress-cli.go/progress"
)

func main() {
	var err error

	helpFlag := flag.Bool("help", false, "Print help")
	versionFlag := flag.Bool("version", false, "Print Version")
	widthFlag := flag.Int("width", define.PROGRESS_DEFAULT_WIDTH, "Width")

	flag.Parse()

	switch {
	case *helpFlag:
		info.InfoHelp()
		os.Exit(0)
	case *versionFlag:
		info.InfoVersion()
		os.Exit(0)
	}

	meta, err := progress.PrintMetaNew(time.Now(), *widthFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	meta.PrintSection1()
	meta.PrintSectionProgress()
}
