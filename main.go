package main

import (
	"flag"
	"time"

 	_ "github.com/bmhatfield/go-runtime-metrics"
)

func main() {
	flag.Parse()

	benchmarkUnBufferedChannel()
	time.Sleep(1 * time.Minute)
	benchmarkBufferedChannel()
	time.Sleep(1 * time.Minute)
	benchmarkBufferedChannelNConsumers(3)
	time.Sleep(1 * time.Minute)
	benchmarkNBufferedChannelNConsumers(3, 3)
}