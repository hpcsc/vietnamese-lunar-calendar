package main

import (
	"syscall/js"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
	"github.com/hpcsc/vietnamese-lunar-calendar/internal/ics"
	"github.com/hungtrd/amlich"
)

func convertSolarToLunar(this js.Value, args []js.Value) interface{} {
	year := args[0].Int()
	month := args[1].Int()
	day := args[2].Int()

	lunarDay, lunarMonth, lunarYear, _ := amlich.Solar2Lunar(day, month, year, 7)

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

	solarDay, solarMonth, solarYear := amlich.Lunar2Solar(day, month, year, 0, 7)

	return map[string]interface{}{
		"year":  solarYear,
		"month": solarMonth,
		"day":   solarDay,
	}
}

func generateICS(this js.Value, args []js.Value) interface{} {
	yearsAhead := args[0].Int()
	customEvents := args[1].String()

	startYear := time.Now().Year()
	gen := calendar.NewGenerator(startYear, yearsAhead)
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
