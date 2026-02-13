package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hungtrd/amlich"
)

var (
	yearsAhead   = flag.Int("years", 10, "Number of years ahead to generate")
	outputFile   = flag.String("output", "vietnamese-lunar-calendar.ics", "Output ICS file path")
	customEvents = flag.String("events", "", "Custom lunar events in format 'day/month:title' (recurring) or 'day/month/year:title' (single year), can be repeated")
)

func main() {
	flag.Parse()

	startYear := time.Now().Year()

	var events []LunarEvent
	if *customEvents != "" {
		events = parseCustomEvents(*customEvents, startYear, *yearsAhead)
	} else {
		events = generateLunarEvents(startYear, *yearsAhead)
	}

	icsContent := generateICS(events)

	err := os.WriteFile(*outputFile, []byte(icsContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write ICS file: %v", err)
	}

	fmt.Printf("Generated ICS file with %d events for years %d-%d to %s\n",
		len(events), startYear, startYear+*yearsAhead, *outputFile)
}

type LunarDate struct {
	Day   int
	Month int
	Show  bool
}

type LunarEvent struct {
	Title       string
	Date        time.Time
	LunarDate   LunarDate
	Description string
}

func generateLunarEvents(startYear, yearsAhead int) []LunarEvent {
	var events []LunarEvent

	for year := startYear; year < startYear+yearsAhead; year++ {
		events = append(events, getEventsForYear(year)...)
	}

	return events
}

func getEventsForYear(year int) []LunarEvent {
	var events []LunarEvent

	// Find Tết Nguyên Đán (Lunar New Year)
	tetDate := findTet(year)
	events = append(events, LunarEvent{
		Title:       "Tết Nguyên Đán",
		Date:        tetDate,
		LunarDate:   LunarDate{Day: 1, Month: 1, Show: true},
		Description: "Tết Nguyên Đán - Vietnamese Lunar New Year",
	})

	// Find Tết Thượng Nguyên (Lantern Festival - 15th day of 1st lunar month)
	events = append(events, LunarEvent{
		Title:       "Tết Thượng Nguyên",
		Date:        tetDate.AddDate(0, 0, 14),
		LunarDate:   LunarDate{Day: 15, Month: 1, Show: true},
		Description: "Tết Thượng Nguyên - Rằm tháng Giêng",
	})

	// Giỗ Tổ Hùng Vương (10th day of 3rd lunar month)
	events = append(events, LunarEvent{
		Title:       "Giỗ Tổ Hùng Vương",
		Date:        findHungKingCommemoration(year),
		LunarDate:   LunarDate{Day: 10, Month: 3, Show: true},
		Description: "Giỗ Tổ Hùng Vương",
	})

	// Đoan Ngọ (5th day of 5th lunar month)
	events = append(events, LunarEvent{
		Title:       "Tết Đoan Ngọ",
		Date:        findDuongNgoc(year),
		LunarDate:   LunarDate{Day: 5, Month: 5, Show: true},
		Description: "Tết Đoan Ngọ - Mùng 5 tháng 5",
	})

	// Vu Lan (15th day of 7th lunar month)
	vuLan := findVuLan(year)
	events = append(events, LunarEvent{
		Title:       "Vu Lan",
		Date:        vuLan,
		LunarDate:   LunarDate{Day: 15, Month: 7, Show: true},
		Description: "Vu Lan - Rằm tháng 7",
	})

	// Tết Trung Thu (15th day of 8th lunar month)
	trungThu := findTrungThu(year)
	events = append(events, LunarEvent{
		Title:       "Tết Trung Thu",
		Date:        trungThu,
		LunarDate:   LunarDate{Day: 15, Month: 8, Show: true},
		Description: "Tết Trung Thu - Rằm tháng 8",
	})

	// First day of each lunar month (skip if already covered by another event)
	events = append(events, getFirstDayOfLunarMonths(year, events)...)

	return events
}

func findTet(year int) time.Time {
	for month := 1; month <= 2; month++ {
		dayStart := 21
		if month == 2 {
			dayStart = 1
		}
		dayEnd := 31
		if month == 2 {
			dayEnd = 20
		}

		for day := dayStart; day <= dayEnd; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == 1 && lunar.Day == 1 && lunar.Year == year {
				return t
			}
		}
	}
	return time.Date(year, 1, 21, 0, 0, 0, 0, time.UTC)
}

func findHungKingCommemoration(year int) time.Time {
	for month := 3; month <= 5; month++ {
		for day := 1; day <= 30; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == 3 && lunar.Day == 10 {
				return t
			}
		}
	}
	return time.Date(year, 4, 10, 0, 0, 0, 0, time.UTC)
}

func findDuongNgoc(year int) time.Time {
	for month := 4; month <= 7; month++ {
		for day := 1; day <= 30; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == 5 && lunar.Day == 5 {
				return t
			}
		}
	}
	return time.Date(year, 5, 5, 0, 0, 0, 0, time.UTC)
}

func findVuLan(year int) time.Time {
	for month := 7; month <= 9; month++ {
		for day := 1; day <= 30; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == 7 && lunar.Day == 15 {
				return t
			}
		}
	}
	return time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC)
}

func findTrungThu(year int) time.Time {
	for month := 8; month <= 10; month++ {
		for day := 1; day <= 30; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == 8 && lunar.Day == 15 {
				return t
			}
		}
	}
	return time.Date(year, 9, 15, 0, 0, 0, 0, time.UTC)
}

func getFirstDayOfLunarMonths(year int, existingEvents []LunarEvent) []LunarEvent {
	var events []LunarEvent

	existingLunarDates := make(map[string]bool)
	for _, e := range existingEvents {
		existingLunarDates[fmt.Sprintf("%d/%d", e.LunarDate.Day, e.LunarDate.Month)] = true
	}

	for month := 1; month <= 12; month++ {
		lunarDateKey := fmt.Sprintf("1/%d", month)
		if existingLunarDates[lunarDateKey] {
			continue
		}
		date := findLunarDate(year, month, 1)
		if !date.IsZero() {
			events = append(events, LunarEvent{
				Title:       fmt.Sprintf("Mùng 1 Tháng %d (Âm lịch)", month),
				Date:        date,
				LunarDate:   LunarDate{Day: 1, Month: month, Show: false},
				Description: fmt.Sprintf("Mùng 1 tháng %d âm lịch", month),
			})
		}
	}

	return events
}

func findLunarDate(year, lunarMonth, lunarDay int) time.Time {
	for month := 1; month <= 12; month++ {
		for day := 1; day <= 31; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == lunarMonth && lunar.Day == lunarDay {
				return t
			}
		}
	}
	return time.Time{}
}

func parseCustomEvents(eventsStr string, startYear, yearsAhead int) []LunarEvent {
	var events []LunarEvent

	parts := strings.Split(eventsStr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		kv := strings.Split(part, ":")
		if len(kv) != 2 {
			log.Printf("Invalid format: %s, expected day/month:title or day/month/year:title", part)
			continue
		}

		datePart := strings.TrimSpace(kv[0])
		title := strings.TrimSpace(kv[1])

		dateParts := strings.Split(datePart, "/")
		if len(dateParts) == 2 {
			// recurring event: day/month:title
			var day, month int
			fmt.Sscanf(dateParts[0], "%d", &day)
			fmt.Sscanf(dateParts[1], "%d", &month)

			for year := startYear; year < startYear+yearsAhead; year++ {
				date := findLunarDate(year, month, day)
				if !date.IsZero() {
					events = append(events, LunarEvent{
						Title:       title,
						Date:        date,
						LunarDate:   LunarDate{Day: day, Month: month, Show: true},
						Description: fmt.Sprintf("%s - Ngày %d tháng %d âm lịch", title, day, month),
					})
				}
			}
		} else if len(dateParts) == 3 {
			// single event: day/month/year:title
			var day, month, year int
			fmt.Sscanf(dateParts[0], "%d", &day)
			fmt.Sscanf(dateParts[1], "%d", &month)
			fmt.Sscanf(dateParts[2], "%d", &year)

			date := findLunarDate(year, month, day)
			if !date.IsZero() {
				events = append(events, LunarEvent{
					Title:       title,
					Date:        date,
					LunarDate:   LunarDate{Day: day, Month: month, Show: true},
					Description: fmt.Sprintf("%s - Ngày %d tháng %d năm %d âm lịch", title, day, month, year),
				})
			}
		} else {
			log.Printf("Invalid date format: %s, expected day/month:title or day/month/year:title", datePart)
		}
	}

	return events
}

func generateICS(events []LunarEvent) string {
	buf := &strings.Builder{}

	buf.WriteString("BEGIN:VCALENDAR\r\n")
	buf.WriteString("VERSION:2.0\r\n")
	buf.WriteString("PRODID:-//Vietnamese Lunar Calendar//EN\r\n")
	buf.WriteString("CALSCALE:GREGORIAN\r\n")
	buf.WriteString("METHOD:PUBLISH\r\n")
	buf.WriteString("X-WR-CALNAME:Vietnamese Lunar Calendar\r\n")
	buf.WriteString("X-WR-TIMEZONE:Asia/Hanoi\r\n")

	for _, e := range events {
		dateStr := e.Date.Format("20060102")

		buf.WriteString("BEGIN:VEVENT\r\n")
		buf.WriteString(fmt.Sprintf("UID:vnlunar-%s-%s@lunar-calendar\r\n",
			e.Title, dateStr))
		buf.WriteString("DTSTAMP:" + time.Now().UTC().Format("20060102T150405Z") + "\r\n")
		buf.WriteString(fmt.Sprintf("DTSTART;VALUE=DATE:%s\r\n", dateStr))
		buf.WriteString(fmt.Sprintf("DTEND;VALUE=DATE:%s\r\n", dateStr))
		summary := e.Title
		if e.LunarDate.Show {
			summary = fmt.Sprintf("%s (%d/%d)", e.Title, e.LunarDate.Day, e.LunarDate.Month)
		}
		buf.WriteString(fmt.Sprintf("SUMMARY:%s\r\n", summary))
		if e.Description != "" {
			buf.WriteString(fmt.Sprintf("DESCRIPTION:%s\r\n", e.Description))
		}
		buf.WriteString("STATUS:CONFIRMED\r\n")
		buf.WriteString("END:VEVENT\r\n")
	}

	buf.WriteString("END:VCALENDAR\r\n")

	return buf.String()
}
