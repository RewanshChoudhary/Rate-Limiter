package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"

	readingdata "github.com/RewanshChoudhary/System-Monitor/readingData"
)

func main() {
	fmt.Println("System data sending data ")

	runCronJob()

	

}
func runCronJob() {
	schedule := gocron.NewScheduler(time.Local)
	schedule.Every(1).Second().Do(collectStatsAndPushToProducer)

	schedule.StartAsync()

}

func collectStatsAndPushToProducer() {

	stats, err := readingdata.ReadMemoryStatus()

	if err != nil {
		fmt.Printf("Error reading stats: %v\n", err)
		return
	}
	_, err = readingdata.SendStatsToProducer(stats)
	if err != nil {
		fmt.Printf("Error reading stats: %v\n", err)
		return
	}
}
