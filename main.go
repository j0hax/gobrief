package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/j0hax/gobrief/agenda"
	"github.com/j0hax/gobrief/config"
)

func customUsage() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "Usage: %s [CALENDAR]\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	cfg := config.LoadConfig()

	flag.Usage = customUsage
	list := flag.Bool("list", false, "list calendars")
	add := flag.Bool("add", false, "add calendar sources in the pattern [NAME] [URL]")
	del := flag.Bool("del", false, "remove calender by [NAME]")
	nDays := flag.Int("days", cfg.Days, "number of days to look ahead")
	flag.Parse()

	if *list {
		out := flag.CommandLine.Output()
		cfg.ListCalendars(out)
		os.Exit(0)
	}

	if *add {
		cfg.AddCalendar(flag.Args()...)
		cfg.SaveExit()
	} else if *del {
		cfg.DeleteCalendar(flag.Args()...)
		cfg.SaveExit()
	}

	if len(flag.Args()) > 0 {
		cfg.SelectCalendars(flag.Args()...)
	}

	a := agenda.NewAgenda(*nDays, cfg.Calendars)
	a.PrettyPrint(os.Stdout)
}
