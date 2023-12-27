package agenda

import (
	"container/heap"
	"sync"

	"github.com/j0hax/gobrief/config"
)

// NewAgenda fetches the specified calendars into a sorted list
func NewAgenda(daysAhead int, cals []config.Calendar) *EventHeap {
	var wg sync.WaitGroup

	h := &EventHeap{}

	heap.Init(h)

	for _, cal := range cals {
		wg.Add(1)
		go func(calName, url string, days int) {
			defer wg.Done()
			for _, event := range fetchCal(url, days) {
				heap.Push(h, EventEntry{
					Event:        event,
					CalendarName: calName,
				})
			}
		}(cal.Name, cal.URL, daysAhead)
	}
	wg.Wait()

	return h
}
