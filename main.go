package main

import (
	"log"

	"github.com/j0hax/gobrief/config"
)

func main() {
	cfg := config.LoadConfig()

	events := Fetch(cfg.Days, cfg.Calendars)

	printCal(events)

	err := cfg.Save()
	if err != nil {
		log.Panic(err)
	}
}
