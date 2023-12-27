package agenda

import (
	"log"
	"net/http"
	"time"

	"github.com/apognu/gocal"
)

func fetchCal(url string, numDays int) []gocal.Event {
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

	return c.Events
}
