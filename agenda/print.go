package agenda

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jwalton/gchalk"
)

var dateCol = gchalk.WithBold().Yellow
var timeCol = gchalk.Red
var calCol = gchalk.Blue
var eventCol = gchalk.WithItalic().Green

// printCal prints the list of Events to stdout using custom formatting
func (a Agenda) PrintCal() {
	// Init custom formatting
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	for _, event := range a.Sorted() {
		// Establish formats
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

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", dateCol(date), timeCol(duration), calCol(event.CalendarName), eventCol(event.Summary))
	}

	w.Flush()
}
