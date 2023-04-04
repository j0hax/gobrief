package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
	"github.com/jwalton/gchalk"
)

// Entry encapsulates a single event with its parent calendar name
type Entry struct {
	CalendarName string
	Event        gocal.Event
}

// Fetch concurrenctly fetches iCal events from a map of calendar names and source URLs
//
// The returned list of events is sorted by date.
func Fetch(days int, cals []config.Calendar) []Entry {
	events := make([]Entry, 0, 64)
	results := make(chan Entry, len(cals))

	// fetch each URL concurrently
	var wg sync.WaitGroup
	for _, cal := range cals {
		wg.Add(1)
		go fetchCal(cal.Name, cal.URL, days, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Add each item to the list:
	// The range ends when all goroutines are finished and the channel is closed.
	for r := range results {
		events = append(events, r)
	}

	sort.Slice(events, func(i, j int) bool {
		iStart := events[i].Event.Start
		jStart := events[j].Event.Start
		return iStart.Before(*jStart)
	})

	return events
}

// fetchCal fetches iCal data for the specified days from a URL.
//
// This function is designed to run concurrently, and as such writes events
// into a channel.
func fetchCal(name, url string, numDays int, ch chan<- Entry, wg *sync.WaitGroup) {
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
		event := Entry{
			CalendarName: name,
			Event:        e,
		}
		ch <- event
	}
}

// printCal prints the list of Events to stdout using custom formatting
func printCal(events []Entry) {
	// Init custom formatting
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	dateCol := gchalk.WithBold().Yellow
	timeCol := gchalk.Red
	calCol := gchalk.Blue
	eventCol := gchalk.WithItalic().Green

	for _, e := range events {
		c := e.Event

		// Establish formats
		date := c.Start.Format("Mon 02 Jan")

		var duration string
		// Check if the event is all-day
		_, allday := c.RawStart.Params["VALUE"]
		if allday {
			duration = "all day"
		} else {
			start := c.Start.Format("15:04")
			end := c.End.Format("15:04")
			duration = fmt.Sprintf("%s - %s", start, end)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", dateCol(date), timeCol(duration), calCol(e.CalendarName), eventCol(c.Summary))
	}

	w.Flush()
}
