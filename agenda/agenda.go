package agenda

import (
	"sync"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

// An agenda is a collection of calendars
type Agenda map[string][]gocal.Event

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
