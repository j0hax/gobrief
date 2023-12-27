package agenda

import (
	"container/heap"
	"sync"

	"github.com/j0hax/gobrief/config"
)

// NewAgenda fetches the specified calendars into a sorted list
func NewAgenda(cfg *config.Configuration) *Agenda {
	var wg sync.WaitGroup

	h := &Agenda{}

	heap.Init(h)

	for _, cal := range cfg.Calendars {
		wg.Add(1)
		go func(cal config.Calendar, days int) {
			defer wg.Done()
			for _, event := range fetchCal(cal.URL, days) {
				heap.Push(h, EventEntry{
					Event:    event,
					Calendar: &cal,
				})
			}
		}(cal, cfg.Days)
	}
	wg.Wait()

	return h
}
