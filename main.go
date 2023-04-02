package main

import (
	"sort"
	"sync"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

func main() {
	cfg := config.LoadConfig()
	events := make([]gocal.Event, 0, 64)

	eventChan := make(chan gocal.Event)

	// get each
	var wg sync.WaitGroup
	for _, url := range cfg.Calendars {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fetchCal(url, cfg.Days, eventChan)
		}(url)
	}

	go func() {
		for {
			events = append(events, <-eventChan)
		}
	}()

	wg.Wait()

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(*events[j].Start)
	})

	printCal(events)

	defer cfg.Save()
}
