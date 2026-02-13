package lunar

import (
	"time"

	"github.com/hungtrd/amlich"
)

type Date struct {
	Day   int
	Month int
}

type Range struct {
	StartMonth int
	StartDay   int
	EndMonth   int
	EndDay     int
}

type FindOption func(*findConfig)

type findConfig struct {
	range_ Range
}

func WithRange(r Range) FindOption {
	return func(c *findConfig) {
		c.range_ = r
	}
}

var (
	Tet                   = Date{Month: 1, Day: 1}
	HungKingCommemoration = Date{Day: 10, Month: 3}
	DuongNgoc             = Date{Day: 5, Month: 5}
	VuLan                 = Date{Day: 15, Month: 7}
	TrungThu              = Date{Day: 15, Month: 8}

	TetRange = Range{StartMonth: 1, StartDay: 21, EndMonth: 2, EndDay: 20}
)

func FindLunarDate(year int, ld Date, opts ...FindOption) time.Time {
	cfg := &findConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	startMonth, endMonth := 1, 12
	startDay, endDay := 1, 31

	if cfg.range_.StartMonth > 0 {
		startMonth = cfg.range_.StartMonth
		startDay = cfg.range_.StartDay
	}
	if cfg.range_.EndMonth > 0 {
		endMonth = cfg.range_.EndMonth
		endDay = cfg.range_.EndDay
	}

	for month := startMonth; month <= endMonth; month++ {
		dayStart := 1
		dayEnd := 31
		if month == startMonth {
			dayStart = startDay
		}
		if month == endMonth {
			dayEnd = endDay
		}

		for day := dayStart; day <= dayEnd; day++ {
			t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			lunar := amlich.New(t.In(amlich.VietnamLocation()))
			if lunar.Month == ld.Month && lunar.Day == ld.Day {
				return t
			}
		}
	}
	return time.Time{}
}
