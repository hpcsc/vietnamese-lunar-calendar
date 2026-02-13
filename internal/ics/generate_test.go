package ics_test

import (
	"strings"
	"testing"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
	"github.com/hpcsc/vietnamese-lunar-calendar/internal/ics"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("generates valid ICS header", func(t *testing.T) {
		events := []calendar.Event{
			{
				Title:       "Test Event",
				Date:        time.Date(2026, time.February, 17, 0, 0, 0, 0, time.UTC),
				LunarDate:   calendar.LunarDate{Day: 1, Month: 1, Show: true},
				Description: "Test Description",
			},
		}

		result := ics.Generate(events)

		require.Contains(t, result, "BEGIN:VCALENDAR")
		require.Contains(t, result, "VERSION:2.0")
		require.Contains(t, result, "END:VCALENDAR")
		require.Equal(t, 1, strings.Count(result, "BEGIN:VCALENDAR"))
	})

	t.Run("includes event with correct date", func(t *testing.T) {
		eventDate := time.Date(2026, time.February, 17, 0, 0, 0, 0, time.UTC)
		events := []calendar.Event{
			{
				Title:       "Tết Nguyên Đán",
				Date:        eventDate,
				LunarDate:   calendar.LunarDate{Day: 1, Month: 1, Show: true},
				Description: "Vietnamese Lunar New Year",
			},
		}

		result := ics.Generate(events)

		require.Contains(t, result, "DTSTART;VALUE=DATE:20260217")
		require.Contains(t, result, "SUMMARY:Tết Nguyên Đán (1/1)")
	})

	t.Run("excludes lunar date in summary when Show is false", func(t *testing.T) {
		events := []calendar.Event{
			{
				Title:       "Mùng 1 Tháng 2 (Âm lịch)",
				Date:        time.Date(2026, time.February, 17, 0, 0, 0, 0, time.UTC),
				LunarDate:   calendar.LunarDate{Day: 1, Month: 2, Show: false},
				Description: "First day of lunar month 2",
			},
		}

		result := ics.Generate(events)

		require.Contains(t, result, "SUMMARY:Mùng 1 Tháng 2 (Âm lịch)")
		require.NotContains(t, result, "(1/2)")
	})

	t.Run("includes multiple events", func(t *testing.T) {
		events := []calendar.Event{
			{
				Title:       "Event 1",
				Date:        time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC),
				LunarDate:   calendar.LunarDate{Day: 1, Month: 1, Show: true},
				Description: "Description 1",
			},
			{
				Title:       "Event 2",
				Date:        time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC),
				LunarDate:   calendar.LunarDate{Day: 1, Month: 2, Show: true},
				Description: "Description 2",
			},
		}

		result := ics.Generate(events)

		require.Equal(t, 2, strings.Count(result, "BEGIN:VEVENT"))
		require.Contains(t, result, "SUMMARY:Event 1")
		require.Contains(t, result, "SUMMARY:Event 2")
	})

	t.Run("includes UID for each event", func(t *testing.T) {
		events := []calendar.Event{
			{
				Title:       "Test Event",
				Date:        time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC),
				LunarDate:   calendar.LunarDate{Day: 1, Month: 1, Show: true},
				Description: "Test",
			},
		}

		result := ics.Generate(events)

		require.Contains(t, result, "UID:vnlunar-")
	})
}
