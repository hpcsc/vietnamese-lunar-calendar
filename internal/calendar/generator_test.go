package calendar_test

import (
	"maps"
	"slices"
	"testing"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
	"github.com/stretchr/testify/require"
)

func findEventByTitle(events []calendar.Event, title string) *calendar.Event {
	for i := range events {
		if events[i].Title == title {
			return &events[i]
		}
	}
	return nil
}

func TestGenerator_Generate(t *testing.T) {
	t.Run("generates events for default years", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		events, err := gen.Generate("")

		require.NoError(t, err)
		require.NotEmpty(t, events)
	})

	t.Run("generates events for multiple years", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 2, "Asia/Hanoi")
		events, err := gen.Generate("")

		require.NoError(t, err)
		years := make(map[int]bool)
		for _, e := range events {
			years[e.Date.Year()] = true
		}
		require.ElementsMatch(t, []int{2026, 2027}, slices.Collect(maps.Keys(years)))
	})
}

func TestGenerator_CustomEvents(t *testing.T) {
	t.Run("parses recurring custom event", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 2, "Asia/Hanoi")
		events, err := gen.Generate("15/8:My Birthday")

		require.NoError(t, err)
		require.Len(t, events, 2)

		require.Equal(t, "My Birthday", events[0].Title)
		require.Equal(t, 2026, events[0].Date.Year())
		require.Equal(t, time.September, events[0].Date.Month())
		require.Equal(t, 25, events[0].Date.Day())

		require.Equal(t, "My Birthday", events[1].Title)
		require.Equal(t, 2027, events[1].Date.Year())
		require.Equal(t, time.September, events[1].Date.Month())
		require.Equal(t, 14, events[1].Date.Day())
	})

	t.Run("parses single year custom event", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 2, "Asia/Hanoi")
		events, err := gen.Generate("15/8/2027:One Time Event")

		require.NoError(t, err)
		require.Len(t, events, 1)
		require.Equal(t, "One Time Event", events[0].Title)
		require.Equal(t, 2027, events[0].Date.Year())
		require.Equal(t, time.September, events[0].Date.Month())
		require.Equal(t, 14, events[0].Date.Day())
	})

	t.Run("parses multiple custom events", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		events, err := gen.Generate("4/5:Event 1,15/8:Event 2")

		require.NoError(t, err)
		require.Len(t, events, 2)

		require.Equal(t, "Event 1", events[0].Title)
		require.Equal(t, 2026, events[0].Date.Year())
		require.Equal(t, time.June, events[0].Date.Month())
		require.Equal(t, 18, events[0].Date.Day())

		require.Equal(t, "Event 2", events[1].Title)
		require.Equal(t, 2026, events[1].Date.Year())
		require.Equal(t, time.September, events[1].Date.Month())
		require.Equal(t, 25, events[1].Date.Day())
	})

	t.Run("invalid format returns error", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		_, err := gen.Generate("invalid")

		require.Error(t, err)
	})

	t.Run("empty title returns error", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		_, err := gen.Generate("15/8:")

		require.Error(t, err)
	})

	t.Run("missing day returns error", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		_, err := gen.Generate("/8:Event")

		require.Error(t, err)
	})

	t.Run("missing month returns error", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		_, err := gen.Generate("15/:Event")

		require.Error(t, err)
	})
}

func TestGenerator_DefaultEvents(t *testing.T) {
	tests := []struct {
		name         string
		eventTitle   string
		expectedYear int
	}{
		{"includes Tet event", "Tết Nguyên Đán", 2026},
		{"includes Vu Lan event", "Vu Lan", 2026},
		{"includes Trung Thu event", "Tết Trung Thu", 2026},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
			events, err := gen.Generate("")

			require.NoError(t, err)
			event := findEventByTitle(events, tt.eventTitle)
			require.NotNil(t, event, "event %q not found", tt.eventTitle)
			require.Equal(t, tt.expectedYear, event.Date.Year())
		})
	}
}

func TestEvent_LunarDate(t *testing.T) {
	t.Run("default events have Show true", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		events, err := gen.Generate("")

		require.NoError(t, err)
		event := findEventByTitle(events, "Tết Nguyên Đán")
		require.NotNil(t, event)
		require.True(t, event.LunarDate.Show)
	})

	t.Run("first day of month events have Show false", func(t *testing.T) {
		gen := calendar.NewGenerator(2026, 1, "Asia/Hanoi")
		events, err := gen.Generate("")

		require.NoError(t, err)
		event := findEventByTitle(events, "Mùng 1 Tháng 2 (Âm lịch)")
		require.NotNil(t, event)
		require.False(t, event.LunarDate.Show)
	})
}
