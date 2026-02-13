package main

import (
	"syscall/js"

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

func registerCallbacks() {
	js.Global().Set("convertSolarToLunar", js.FuncOf(convertSolarToLunar))
	js.Global().Set("convertLunarToSolar", js.FuncOf(convertLunarToSolar))
}

func main() {
	registerCallbacks()
	<-make(chan bool)
}
