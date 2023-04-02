package main

import (
	"log"
	"sort"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

func main() {
	cfg := config.LoadConfig()
	events := make([]gocal.Event, 0, 64)

	// get each
	for _, url := range cfg.Calendars {
		results, err := fetchCal(url, cfg.Days)
		if err != nil {
			log.Println(err)
		}
		events = append(events, results...)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Start.Before(*events[j].Start)
	})

	printCal(events)

	defer cfg.Save()
}
