package config

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"text/tabwriter"
)

// ListCalendars prints the list of calendars in the configuration
// to the interface implemented by out.
func (cfg *Configuration) ListCalendars(out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	for i, j := range cfg.Calendars {
		fmt.Fprintf(w, "%s\t%s\t\n", i, j)
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
		cfg.Calendars[k] = v
		log.Printf("Saved calendar as '%s'\n", k)
	}
}

// DeleteCalendar removes a calendar from the configuration.
//
// It expects calendar names to be passed via args
func (cfg *Configuration) DeleteCalendar(args ...string) {
	for _, name := range args {
		for key := range cfg.Calendars {
			if name == key {
				delete(cfg.Calendars, key)
				log.Printf("Deleted calendar '%s'\n", key)
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
	newCalendars := make(map[string]string)

	// Check if the item is a URL
	for i, arg := range args {
		value, ok := cfg.Calendars[arg]
		if ok {
			newCalendars[arg] = value
		} else {
			value, err := url.ParseRequestURI(arg)
			if err == nil {
				name := fmt.Sprintf("url%d", i)
				newCalendars[name] = value.String()
				break
			} else {
				log.Printf("Neither a configured calendar nor valid URL: '%s'\n", arg)
			}
		}
	}

	cfg.Calendars = newCalendars
}
