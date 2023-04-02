package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/apognu/gocal"
	"github.com/jwalton/gchalk"
)

func fetchCal(url string, numDays int, ch chan<- gocal.Event) {
	// Grab ical
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	// Go ahead 2 days
	start, end := time.Now(), time.Now().AddDate(0, 0, numDays)

	// Parse Ical
	c := gocal.NewParser(resp.Body)
	c.Start, c.End = &start, &end

	err = c.Parse()
	if err != nil {
		log.Panic(err)
	}

	for _, e := range c.Events {
		ch <- e
	}
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
