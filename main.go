package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/stephengeller/pomodoro/file_processor"
	"github.com/stephengeller/pomodoro/internal"
)

const dummyHosts = "/tmp/dummyhosts"

var logger = &internal.DefaultLogger{}

func main() {
	bl := flag.String("blacklist", "", "Path to blacklist file")
	mtw := flag.Duration("time", 0, "Number of minutes to wait")
	flag.Parse()

	blacklist := *bl
	minutesToWait := *mtw

	sitesToBlacklist, err := FileProcessor.ReadFile(blacklist)
	internal.Check(logger, err)

	err = FileProcessor.AddLinesToFile(dummyHosts, sitesToBlacklist)
	internal.Check(logger, err)

	fmt.Printf("Timer running for %s...\n", minutesToWait.String())

	timer1 := time.NewTimer(minutesToWait)
	<-timer1.C

	fmt.Printf("%d minute timer expired\n", minutesToWait)

	err = FileProcessor.RemoveLinesFromFile(dummyHosts, sitesToBlacklist)
	internal.Check(logger, err)
}
