package config

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"text/tabwriter"

	"golang.org/x/exp/slices"
)

// ListCalendars prints the list of calendars in the configuration
// to the interface implemented by out.
func (cfg *Configuration) ListCalendars(out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	for _, j := range cfg.Calendars {
		fmt.Fprintf(w, "%s\t%s\t\n", j.Name, j.URL)
	}

	w.Flush()
}

// AddCalendar adds a new named calendar to the configuration.
//
// It expects calendar name and URL in alternating order.
func (cfg *Configuration) AddCalendar(args ...string) {
	count := len(args)

	if count%2 != 0 {
		log.Fatal("Arguments must be in form [name] [URL]!")
	}

	for i := 0; i < len(args)-1; i += 2 {
		k := args[i]
		v := args[i+1]

		item := Calendar{Name: k, URL: v}
		cfg.Calendars = append(cfg.Calendars, item)
		log.Printf("Saved calendar as '%s'\n", k)
	}
}

// DeleteCalendar removes a calendar from the configuration.
//
// It expects calendar names to be passed via args
func (cfg *Configuration) DeleteCalendar(args ...string) {
	for _, name := range args {
		for i, cal := range cfg.Calendars {
			if name == cal.Name {
				cfg.Calendars = append(cfg.Calendars[:i], cfg.Calendars[i+1:]...)
				log.Printf("Deleted calendar '%s'\n", name)
				break
			}
		}
	}
}

// SelectCalendars processes arguments into a map of calendars:
//
// - If an argument matches the name of a configured calendar, it will be included
//
// - If valid URL is passed, it will be included
//
// All other calendars will be filtered from the configuration.
func (cfg *Configuration) SelectCalendars(args ...string) {
	newCalendars := make([]Calendar, len(args))

	for i, arg := range args {
		// Search for calendar with matching name
		idx := slices.IndexFunc(cfg.Calendars, func(c Calendar) bool {
			return c.Name == arg
		})

		// If the calendar exists, append it.
		// Otherwise, parse the URL
		if idx >= 0 {
			newCalendars = append(newCalendars, cfg.Calendars[idx])
		} else {
			value, err := url.ParseRequestURI(arg)
			if err != nil {
				log.Printf("Neither a configured calendar nor valid URL: '%s'\n", arg)
				break
			}
			cal := Calendar{
				Name: fmt.Sprintf("URL%d", i),
				URL:  value.String(),
			}
			newCalendars = append(newCalendars, cal)
		}
	}

	cfg.Calendars = newCalendars
}
