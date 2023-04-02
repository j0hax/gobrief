package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apognu/gocal"
	"github.com/jwalton/gchalk"
)

func fetchCal(url string, numDays int) ([]gocal.Event, error) {
	// Grab ical
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Go ahead 2 days
	start, end := time.Now(), time.Now().AddDate(0, 0, numDays)

	// Parse Ical
	c := gocal.NewParser(resp.Body)
	c.Start, c.End = &start, &end

	err = c.Parse()
	if err != nil {
		return nil, err
	}

	return c.Events, nil
}

func printCal(events []gocal.Event) {
	// Init custom formatting
	dateCol := gchalk.WithBold().Yellow
	timeCol := gchalk.Red
	eventCol := gchalk.WithItalic().Blue
	for _, e := range events {
		// Establish formats
		date := e.Start.Format("Mon _2 Jan")
		time := e.Start.Format("15:04")

		fmt.Printf("%s %s %s\n", dateCol(date), timeCol(time), eventCol(e.Summary))
	}
}
