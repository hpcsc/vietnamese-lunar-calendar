package main

import (
	"syscall/js"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
	"github.com/hpcsc/vietnamese-lunar-calendar/internal/ics"
	"github.com/hungtrd/amlich"
)

func getTimezoneOffset(tz string, year, month, day int) int {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return 7
	}
	t := time.Date(year, time.Month(month), day, 12, 0, 0, 0, loc)
	_, offset := t.Zone()
	return offset / 3600
}

func convertSolarToLunar(this js.Value, args []js.Value) interface{} {
	year := args[0].Int()
	month := args[1].Int()
	day := args[2].Int()
	timezone := args[3].String()

	tz := getTimezoneOffset(timezone, year, int(month), day)
	lunarDay, lunarMonth, lunarYear, _ := amlich.Solar2Lunar(day, month, year, tz)

	return map[string]interface{}{
		"day":   lunarDay,
		"month": lunarMonth,
		"year":  lunarYear,
	}
}

func convertLunarToSolar(this js.Value, args []js.Value) interface{} {
	year := args[0].Int()
	month := args[1].Int()
	day := args[2].Int()
	timezone := args[3].String()

	tz := getTimezoneOffset(timezone, year, int(month), day)
	solarDay, solarMonth, solarYear := amlich.Lunar2Solar(day, month, year, 0, tz)

	return map[string]interface{}{
		"year":  solarYear,
		"month": solarMonth,
		"day":   solarDay,
	}
}

func generateICS(this js.Value, args []js.Value) interface{} {
	yearsAhead := args[0].Int()
	customEvents := args[1].String()
	timezone := args[2].String()

	startYear := time.Now().Year()
	gen := calendar.NewGenerator(startYear, yearsAhead, timezone)
	events, err := gen.Generate(customEvents)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	icsContent := ics.Generate(events)
	return map[string]interface{}{
		"content": icsContent,
		"count":   len(events),
	}
}

func registerCallbacks() {
	js.Global().Set("convertSolarToLunar", js.FuncOf(convertSolarToLunar))
	js.Global().Set("convertLunarToSolar", js.FuncOf(convertLunarToSolar))
	js.Global().Set("generateICS", js.FuncOf(generateICS))
}

func main() {
	registerCallbacks()
	<-make(chan bool)
}
