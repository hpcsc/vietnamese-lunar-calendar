# Vietnamese Lunar Calendar

Generates an ICS (iCalendar) file with Vietnamese lunar calendar events.

## Events Included

### Default Events (when no custom events specified)
- **Tết Nguyên Đán** - Vietnamese Lunar New Year
- **Tết Thượng Nguyên** - Lantern Festival (Rằm tháng Giêng)
- **Giỗ Tổ Hùng Vương** - Hung Kings' Commemoration
- **Tết Đoan Ngọ** - Duong Ngoc (Mùng 5 tháng 5)
- **Vu Lan** - Ghost Festival (Rằm tháng 7)
- **Tết Trung Thu** - Mid-Autumn Festival (Rằm tháng 8)
- **Mùng 1** - First day of each lunar month (except when another event falls on that day)

## Usage

### Generate with Default Events

```bash
go run main.go
```

### Generate with Custom Events

```bash
go run main.go -events "4/5:XXX,15/8/2026:My Birthday"
```

Format: `day/month:title` (recurring yearly) or `day/month/year:title` (single year)

Example:
- `4/5:XXX` - Custom event on day 4, month 5 (recurs every year)
- `15/8/2026:My Birthday` - Birthday on 15th day of 8th lunar month in 2026 only

### Options

| Flag | Default | Description |
|------|---------|-------------|
| `-years` | 10 | Number of years ahead to generate |
| `-output` | vietnamese-lunar-calendar.ics | Output ICS file path |
| `-events` | (none) | Custom lunar events (day/month:title for recurring, day/month/year:title for single year) |

### Import to Calendar

1. Open Google Calendar or Apple Calendar
2. Import the `.ics` file
3. Events will be added to your calendar

**Note:** Event titles include the lunar date (e.g., "Tết Nguyên Đán (1/1)"). First day of month events show as "Mùng 1 Tháng X (Âm lịch)" without the lunar date suffix.

## GitHub Actions

The project uses GitHub Actions to automatically generate the ICS file:

- **Every 2 months**: Automatically generates a new ICS file with default events
- **On release**: Attaches the ICS file as a release asset

### Manual Run

1. Go to Actions tab
2. Select "Generate Lunar Calendar"
3. Click "Run workflow"
4. Optionally specify custom events in the `years_ahead` field (uses `-events` flag)

### Download Latest Release

Check the [Releases](https://github.com/hpcsc/vietnamese-lunar-calendar/releases) page for the latest ICS file.

## Timezone

All events are generated in Asia/Hanoi timezone (UTC+7).

## License

MIT
