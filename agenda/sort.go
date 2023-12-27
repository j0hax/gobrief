package agenda

import (
	"sync"

	"github.com/apognu/gocal"
)

// EventEntry wraps gocal.Event and includes a CalendarName field
type EventEntry struct {
	gocal.Event
	CalendarName string
}

// EventHeap is a minheap of calendar events.
type EventHeap struct {
	queue []EventEntry
	mu    sync.Mutex
}

func (el *EventHeap) Len() int {
	el.mu.Lock()
	defer el.mu.Unlock()
	return len(el.queue)
}

func (el *EventHeap) Less(i, j int) bool {
	el.mu.Lock()
	defer el.mu.Unlock()
	return el.queue[i].Start.Before(*el.queue[j].Start)
}

func (el *EventHeap) Swap(i, j int) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.queue[j], el.queue[i] = el.queue[i], el.queue[j]
}

func (el *EventHeap) Push(x any) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.queue = append(el.queue, x.(EventEntry))
}

func (el *EventHeap) Pop() any {
	el.mu.Lock()
	defer el.mu.Unlock()
	old := el.queue
	n := len(old)
	x := old[n-1]
	el.queue = old[0 : n-1]
	return x
}
