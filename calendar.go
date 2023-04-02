package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/apognu/gocal"
	"github.com/jwalton/gchalk"
)

// fetchCal fetches iCal data for the specified days from a URL.
//
// This function is designed to run concurrently, and as such writes events
// into a channel.
func fetchCal(url string, numDays int, ch chan<- gocal.Event, wg *sync.WaitGroup) {
	defer wg.Done()

	// Grab iCal
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	// Look up events up to numDays from now
	start, end := time.Now(), time.Now().AddDate(0, 0, numDays)

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

// printCal prints the list of Events to stdout using custom formatting
func printCal(events []gocal.Event) {
	// Init custom formatting
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	dateCol := gchalk.WithBold().Yellow
	timeCol := gchalk.Red
	eventCol := gchalk.WithItalic().Blue

	for _, e := range events {
		// Establish formats
		date := e.Start.Format("Mon 02 Jan")
		time := e.Start.Format("15:04")

		// Check if the event is all-day
		_, allday := e.RawStart.Params["VALUE"]
		if allday {
			time = "all-day"
		}

		fmt.Fprintf(w, "%s\t%s\t%s\n", dateCol(date), timeCol(time), eventCol(e.Summary))
	}

	w.Flush()
}
