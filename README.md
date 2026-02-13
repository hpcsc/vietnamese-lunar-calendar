# Vietnamese Lunar Calendar

Generates an ICS (iCalendar) file with Vietnamese lunar calendar events for the next 10 years.

## Events Included

- **Tết Nguyên Đán** - Vietnamese Lunar New Year
- **Tết Thượng Nguyên** - Lantern Festival (Rằm tháng Giêng)
- **Giỗ Tổ Hùng Vương** - Hung Kings' Commemoration
- **Tết Đoan Ngọ** - Duong Ngoc (Mùng 5 tháng 5)
- **Vu Lan** - Ghost Festival (Rằm tháng 7)
- **Tết Trung Thu** - Mid-Autumn Festival (Rằm tháng 8)

## Usage

### Generate ICS File Locally

```bash
go run main.go
```

This will generate `vietnamese-lunar-calendar.ics` in the current directory.

### Import to Calendar

1. Open Google Calendar or Apple Calendar
2. Import the `.ics` file
3. Events will be added to your calendar

## GitHub Actions

The project uses GitHub Actions to automatically generate the ICS file:

- **Every 2 months**: Automatically generates a new ICS file
- **On release**: Attaches the ICS file as a release asset

### Manual Run

1. Go to Actions tab
2. Select "Generate Lunar Calendar"
3. Click "Run workflow"

### Download Latest Release

Check the [Releases](https://github.com/hpcsc/vietnamese-lunar-calendar/releases) page for the latest ICS file.

## Timezone

All events are generated in Asia/Hanoi timezone (UTC+7).

## License

MIT
