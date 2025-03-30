package main

import (
	"fmt"
	"time"
)

type PrintMeta struct {
	time        time.Time
	totalDays   int
	thisDay     int
	ratio       float32
	width       int
	beforeWidth int
}

func PrintMetaNew(t time.Time, width int) (meta PrintMeta, err error) {
	meta.time = t

	meta.thisDay = t.YearDay()

	if t.Year()%4 == 0 {
		meta.totalDays = 366
	} else {
		meta.totalDays = 365
	}

	meta.ratio = float32(meta.thisDay) / float32(meta.totalDays)

	meta.width = width
	if meta.width < 3 {
		return meta, fmt.Errorf(APP_NAME + ": width less than 3")
	}
	meta.beforeWidth = int(float32(width) * meta.ratio)

	return meta, nil
}

func (meta PrintMeta) PrintSection1() {
	fmt.Printf("It's day %d of the year(%d). That's already %.1f%%\n", meta.thisDay, meta.totalDays, meta.ratio*100)
}

func (meta PrintMeta) PrintSectionProgress() {
	fmt.Printf("[")

	for i := 0; i < meta.beforeWidth; i++ {
		fmt.Printf("-")
	}

	for i := 0; i < meta.width-meta.beforeWidth; i++ {
		fmt.Printf("#")
	}

	fmt.Printf("]\n")
}
