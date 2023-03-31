package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/jwalton/gchalk"
)

func main() {

	url := flag.String("u", "https://calendar.google.com/calendar/ical/en.german%23holiday%40group.v.calendar.google.com/public/basic.ics", "iCal URL")
	days := flag.Int("d", 2, "Days to display")

	flag.Parse()

	// Grab ical
	resp, err := http.Get(*url)

	if err != nil {
		log.Fatal(err)
	}

	// Go ahead 2 days
	start, end := time.Now(), time.Now().AddDate(0, 0, *days)

	// Parse Ical
	c := gocal.NewParser(resp.Body)
	c.Start, c.End = &start, &end
	c.Parse()

	// Init custom formatting
	dateCol := gchalk.WithBold().Yellow
	timeCol := gchalk.Red
	eventCol := gchalk.WithItalic().Blue

	if len(c.Events) == 0 {
		fmt.Printf("No upcoming events in the next %d days\n", *days)
		os.Exit(0)
	}

	for _, e := range c.Events {

		// Establish formats
		date := e.Start.Format("Mon _2 Jan")
		time := e.Start.Format("15:04")

		// Remove StudIP Spam
		sum := strings.ReplaceAll(e.Summary, "Hauptveranst., ", "")

		fmt.Printf("%s %s %s\n", dateCol(date), timeCol(time), eventCol(sum))
	}
}
