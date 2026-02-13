package lunar_test

import (
	"testing"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/lunar"
	"github.com/stretchr/testify/require"
)

func TestFindLunarDate(t *testing.T) {
	t.Run("Tet", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.Tet)
		require.Equal(t, 2026, result.Year())
		require.Equal(t, time.February, result.Month())
		require.Equal(t, 17, result.Day())
	})

	t.Run("HungKingCommemoration", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.HungKingCommemoration)
		require.Equal(t, 2026, result.Year())
		require.Equal(t, time.April, result.Month())
		require.Equal(t, 26, result.Day())
	})

	t.Run("DuongNgoc", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.DuongNgoc)
		require.Equal(t, 2026, result.Year())
		require.Equal(t, time.June, result.Month())
		require.Equal(t, 19, result.Day())
	})

	t.Run("VuLan", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.VuLan)
		require.Equal(t, 2026, result.Year())
		require.Equal(t, time.August, result.Month())
		require.Equal(t, 27, result.Day())
	})

	t.Run("TrungThu", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.TrungThu)
		require.Equal(t, 2026, result.Year())
		require.Equal(t, time.September, result.Month())
		require.Equal(t, 25, result.Day())
	})

	t.Run("invalid lunar month returns zero", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.Date{Month: 13, Day: 1})
		require.True(t, result.IsZero())
	})

	t.Run("invalid lunar day returns zero", func(t *testing.T) {
		result := lunar.FindLunarDate(2026, lunar.Date{Month: 10, Day: 35})
		require.True(t, result.IsZero())
	})

	t.Run("WithRange", func(t *testing.T) {
		t.Run("return date if within range", func(t *testing.T) {
			result := lunar.FindLunarDate(2026, lunar.Tet, lunar.WithRange(lunar.TetRange))
			require.Equal(t, 2026, result.Year())
			require.Equal(t, time.February, result.Month())
			require.Equal(t, 17, result.Day())
		})

		t.Run("return zero if outside of range", func(t *testing.T) {
			// The end of the range is 10/2, while we are looking for 17/2
			result := lunar.FindLunarDate(2026, lunar.Tet, lunar.WithRange(lunar.Range{StartMonth: 1, StartDay: 21, EndMonth: 2, EndDay: 10}))
			require.True(t, result.IsZero())
		})
	})
}
