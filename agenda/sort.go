package agenda

import (
	"sync"

	"github.com/apognu/gocal"
	"github.com/j0hax/gobrief/config"
)

// EventEntry wraps gocal.Event and includes a CalendarName field
type EventEntry struct {
	gocal.Event
	Calendar *config.Calendar
}

// Agenda is a thread-safe minheap of calendar events.
type Agenda struct {
	queue []EventEntry
	mu    sync.RWMutex
}

func (el *Agenda) Len() int {
	el.mu.RLock()
	defer el.mu.RUnlock()
	return len(el.queue)
}

func (el *Agenda) Less(i, j int) bool {
	el.mu.RLock()
	defer el.mu.RUnlock()
	return el.queue[i].Start.Before(*el.queue[j].Start)
}

func (el *Agenda) Swap(i, j int) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.queue[j], el.queue[i] = el.queue[i], el.queue[j]
}

func (el *Agenda) Push(x any) {
	el.mu.Lock()
	defer el.mu.Unlock()
	el.queue = append(el.queue, x.(EventEntry))
}

func (el *Agenda) Pop() any {
	el.mu.Lock()
	defer el.mu.Unlock()
	old := el.queue
	n := len(old)
	x := old[n-1]
	el.queue = old[0 : n-1]
	return x
}
