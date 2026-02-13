package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
	"github.com/hpcsc/vietnamese-lunar-calendar/internal/ics"
)

var (
	yearsAhead   = flag.Int("years", 10, "Number of years ahead to generate")
	outputFile   = flag.String("output", "vietnamese-lunar-calendar.ics", "Output ICS file path")
	customEvents = flag.String("events", "", "Custom lunar events in format 'day/month:title' (recurring) or 'day/month/year:title' (single year), can be repeated")
	timezone     = flag.String("timezone", "Asia/Hanoi", "Timezone for lunar date calculation")
)

func main() {
	flag.Parse()

	startYear := time.Now().Year()

	gen := calendar.NewGenerator(startYear, *yearsAhead, *timezone)
	events, err := gen.Generate(*customEvents)
	if err != nil {
		log.Fatalf("Failed to generate events: %v", err)
	}

	icsContent := ics.Generate(events)

	err = os.WriteFile(*outputFile, []byte(icsContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write ICS file: %v", err)
	}

	fmt.Printf("Generated ICS file with %d events for years %d-%d to %s\n",
		len(events), startYear, startYear+*yearsAhead, *outputFile)
}
