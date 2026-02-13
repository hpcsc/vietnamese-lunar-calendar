package calendar

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/lunar"
)

type LunarDate struct {
	Day   int
	Month int
	Show  bool
}

type Event struct {
	Title       string
	Date        time.Time
	LunarDate   LunarDate
	Description string
}

type Generator struct {
	startYear  int
	yearsAhead int
}

func NewGenerator(startYear, yearsAhead int) *Generator {
	return &Generator{
		startYear:  startYear,
		yearsAhead: yearsAhead,
	}
}

func (g *Generator) Generate(customEvents string) ([]Event, error) {
	if customEvents != "" {
		return g.parseCustomEvents(customEvents)
	}
	return g.generateDefaultEvents(), nil
}

func (g *Generator) generateDefaultEvents() []Event {
	var events []Event

	for year := g.startYear; year < g.startYear+g.yearsAhead; year++ {
		events = append(events, g.getEventsForYear(year)...)
	}

	return events
}

func (g *Generator) getEventsForYear(year int) []Event {
	var events []Event

	tetDate := lunar.FindLunarDate(year, lunar.Tet)
	events = append(events, Event{
		Title:       "Tết Nguyên Đán",
		Date:        tetDate,
		LunarDate:   LunarDate{Day: lunar.Tet.Day, Month: lunar.Tet.Month, Show: true},
		Description: "Tết Nguyên Đán - Vietnamese Lunar New Year",
	})

	events = append(events, Event{
		Title:       "Tết Thượng Nguyên",
		Date:        tetDate.AddDate(0, 0, 14),
		LunarDate:   LunarDate{Day: 15, Month: 1, Show: true},
		Description: "Tết Thượng Nguyên - Rằm tháng Giêng",
	})

	events = append(events, Event{
		Title:       "Giỗ Tổ Hùng Vương",
		Date:        lunar.FindLunarDate(year, lunar.HungKingCommemoration),
		LunarDate:   LunarDate{Day: lunar.HungKingCommemoration.Day, Month: lunar.HungKingCommemoration.Month, Show: true},
		Description: "Giỗ Tổ Hùng Vương",
	})

	events = append(events, Event{
		Title:       "Tết Đoan Ngọ",
		Date:        lunar.FindLunarDate(year, lunar.DuongNgoc),
		LunarDate:   LunarDate{Day: lunar.DuongNgoc.Day, Month: lunar.DuongNgoc.Month, Show: true},
		Description: "Tết Đoan Ngọ - Mùng 5 tháng 5",
	})

	vuLan := lunar.FindLunarDate(year, lunar.VuLan)
	events = append(events, Event{
		Title:       "Vu Lan",
		Date:        vuLan,
		LunarDate:   LunarDate{Day: lunar.VuLan.Day, Month: lunar.VuLan.Month, Show: true},
		Description: "Vu Lan - Rằm tháng 7",
	})

	trungThu := lunar.FindLunarDate(year, lunar.TrungThu)
	events = append(events, Event{
		Title:       "Tết Trung Thu",
		Date:        trungThu,
		LunarDate:   LunarDate{Day: lunar.TrungThu.Day, Month: lunar.TrungThu.Month, Show: true},
		Description: "Tết Trung Thu - Rằm tháng 8",
	})

	events = append(events, g.getFirstDayOfLunarMonths(year, events)...)

	return events
}

func (g *Generator) getFirstDayOfLunarMonths(year int, existingEvents []Event) []Event {
	var events []Event

	existingLunarDates := make(map[string]bool)
	for _, e := range existingEvents {
		existingLunarDates[fmt.Sprintf("%d/%d", e.LunarDate.Day, e.LunarDate.Month)] = true
	}

	for month := 1; month <= 12; month++ {
		lunarDateKey := fmt.Sprintf("1/%d", month)
		if existingLunarDates[lunarDateKey] {
			continue
		}
		date := lunar.FindLunarDate(year, lunar.Date{Month: month, Day: 1})
		if !date.IsZero() {
			events = append(events, Event{
				Title:       fmt.Sprintf("Mùng 1 Tháng %d (Âm lịch)", month),
				Date:        date,
				LunarDate:   LunarDate{Day: 1, Month: month, Show: false},
				Description: fmt.Sprintf("Mùng 1 tháng %d âm lịch", month),
			})
		}
	}

	return events
}

func (g *Generator) parseCustomEvents(eventsStr string) ([]Event, error) {
	var events []Event

	parts := strings.Split(eventsStr, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		kv := strings.Split(part, ":")
		if len(kv) != 2 {
			return nil, errors.New("invalid format: " + part + ", expected day/month:title or day/month/year:title")
		}

		datePart := strings.TrimSpace(kv[0])
		title := strings.TrimSpace(kv[1])

		if title == "" {
			return nil, errors.New("invalid format: " + part + ", title cannot be empty")
		}

		dateParts := strings.Split(datePart, "/")
		if len(dateParts) == 2 {
			var day, month int
			fmt.Sscanf(dateParts[0], "%d", &day)
			fmt.Sscanf(dateParts[1], "%d", &month)

			if day == 0 || month == 0 {
				return nil, errors.New("invalid date: " + datePart + ", day and month must be greater than 0")
			}

			for year := g.startYear; year < g.startYear+g.yearsAhead; year++ {
				date := lunar.FindLunarDate(year, lunar.Date{Month: month, Day: day})
				if !date.IsZero() {
					events = append(events, Event{
						Title:       title,
						Date:        date,
						LunarDate:   LunarDate{Day: day, Month: month, Show: true},
						Description: fmt.Sprintf("%s - Ngày %d tháng %d âm lịch", title, day, month),
					})
				}
			}
		} else if len(dateParts) == 3 {
			var day, month, year int
			fmt.Sscanf(dateParts[0], "%d", &day)
			fmt.Sscanf(dateParts[1], "%d", &month)
			fmt.Sscanf(dateParts[2], "%d", &year)

			if day == 0 || month == 0 {
				return nil, errors.New("invalid date: " + datePart + ", day and month must be greater than 0")
			}

			date := lunar.FindLunarDate(year, lunar.Date{Month: month, Day: day})
			if !date.IsZero() {
				events = append(events, Event{
					Title:       title,
					Date:        date,
					LunarDate:   LunarDate{Day: day, Month: month, Show: true},
					Description: fmt.Sprintf("%s - Ngày %d tháng %d năm %d âm lịch", title, day, month, year),
				})
			}
		} else {
			return nil, errors.New("invalid date format: " + datePart + ", expected day/month:title or day/month/year:title")
		}
	}

	return events, nil
}
