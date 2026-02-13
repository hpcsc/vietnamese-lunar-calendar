package ics

import (
	"fmt"
	"strings"
	"time"

	"github.com/hpcsc/vietnamese-lunar-calendar/internal/calendar"
)

func Generate(events []calendar.Event) string {
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
