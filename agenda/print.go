package agenda

import (
	"container/heap"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/jwalton/gchalk"
)

var dateCol = gchalk.WithBold().Yellow
var timeCol = gchalk.Red
var calCol = gchalk.Blue
var eventCol = gchalk.WithItalic().Green

// PrettyPrint prints the list of Events to output using custom formatting
func (h *Agenda) PrettyPrint(output io.Writer) {
	// Init custom formatting
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', 0)

	for h.Len() > 0 {
		event := heap.Pop(h).(EventEntry)
		date := event.Start.Format("Mon 02 Jan")

		var duration string
		// Check if the event is all-day
		_, allday := event.RawStart.Params["VALUE"]
		if allday {
			duration = "all day"
		} else {
			start := event.Start.Format("15:04")
			end := event.End.Format("15:04")
			duration = fmt.Sprintf("%s - %s", start, end)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", dateCol(date), timeCol(duration), calCol(event.Calendar.Name), eventCol(event.Summary))
	}

	w.Flush()
}
