package agenda

import (
	"slices"

	"github.com/apognu/gocal"
)

// EventEntry wraps gocal.Event and includes a CalendarName field
// in order to make sorting and ranging a map of event slices much easier.
type EventEntry struct {
	gocal.Event
	CalendarName string
}

// Sorted returns the events of an Agenda in a sorted order.
func (a Agenda) Sorted() []EventEntry {
	var allEvents []EventEntry

	// Append all items from map to a slice
	for name, events := range a {
		for _, event := range events {
			allEvents = append(allEvents, EventEntry{
				Event:        event,
				CalendarName: name,
			})
		}
	}

	// Sort the slice
	slices.SortFunc(allEvents, func(a, b EventEntry) int {
		return a.Start.Compare(*b.Start)
	})

	return allEvents
}
