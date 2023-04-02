package main

import (
	"log"
	"sort"
	"sync"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

func main() {
	cfg := config.LoadConfig()
	events := make([]gocal.Event, 0, 64)
	results := make(chan gocal.Event)

	// get each
	var wg sync.WaitGroup
	for _, url := range cfg.Calendars {
		wg.Add(1)
		go fetchCal(url, cfg.Days, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		events = append(events, r)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(*events[j].Start)
	})

	printCal(events)

	err := cfg.Save()
	if err != nil {
		log.Panic(err)
	}
}
