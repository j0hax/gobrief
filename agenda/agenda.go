package agenda

import (
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

// An agenda is a collection of calendars
type Agenda map[string][]gocal.Event

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

	// Sort all events in the calendar
	sort.Slice(c.Events, func(i, j int) bool {
		iStart := c.Events[i].Start
		jStart := c.Events[j].Start
		return iStart.Before(*jStart)
	})

	return c.Events
}

// NewAgenda fetches the specified calendars
func NewAgenda(daysAhead int, cals []config.Calendar) Agenda {
	agenda := make(map[string][]gocal.Event)

	var wg sync.WaitGroup
	for _, cal := range cals {
		wg.Add(1)
		go func(calName, url string, days int) {
			defer wg.Done()
			agenda[calName] = fetchCal(url, days)
		}(cal.Name, cal.URL, daysAhead)
	}
	wg.Wait()

	return agenda
}
