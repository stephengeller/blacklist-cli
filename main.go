package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/stephengeller/pomodoro/file_processor"
)

const dummyHosts = "/tmp/dummyhosts"

func main() {
	bl := flag.String("blacklist", "", "Path to blacklist file")
	mtw := flag.Duration("time", 0, "Number of minutes to wait")
	flag.Parse()

	blacklist := *bl
	minutesToWait := *mtw

	sitesToBlacklist, err := FileProcessor.ReadFile(blacklist)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = FileProcessor.AddLinesToFile(dummyHosts, sitesToBlacklist)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Timer running for %s...\n", minutesToWait.String())
	timer1 := time.NewTimer(minutesToWait)
	<-timer1.C
	fmt.Printf("%d minute timer expired\n", minutesToWait)
	err = FileProcessor.RemoveLinesFromFile(dummyHosts, sitesToBlacklist)

	if err != nil {
		log.Fatal(err)
		return
	}
}
