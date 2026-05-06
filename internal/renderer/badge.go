package renderer

// BadgeLevel represents the severity or status level of a badge.
type BadgeLevel string

const (
	BadgeLevelSuccess BadgeLevel = "success"
	BadgeLevelWarning BadgeLevel = "warning"
	BadgeLevelDanger  BadgeLevel = "danger"
	BadgeLevelInfo    BadgeLevel = "info"
)

// Badge holds display metadata for a summary badge in the HTML report.
type Badge struct {
	Label string
	Value string
	Level BadgeLevel
}

// buildBadges returns a slice of Badge values derived from the report summary
// counts. Each badge communicates a distinct drift metric to the reader.
func buildBadges(managed, unmanaged, missing, changed int) []Badge {
	return []Badge{
		{
			Label: "Managed",
			Value: itoa(managed),
			Level: BadgeLevelSuccess,
		},
		{
			Label: "Unmanaged",
			Value: itoa(unmanaged),
			Level: badgeLevelForCount(unmanaged),
		},
		{
			Label: "Missing",
			Value: itoa(missing),
			Level: badgeLevelForCount(missing),
		},
		{
			Label: "Drifted",
			Value: itoa(changed),
			Level: badgeLevelForCount(changed),
		},
	}
}

// badgeLevelForCount maps a non-zero count to a warning/danger level and a
// zero count to success, giving a quick visual signal in the report.
func badgeLevelForCount(n int) BadgeLevel {
	switch {
	case n == 0:
		return BadgeLevelSuccess
	case n < 5:
		return BadgeLevelWarning
	default:
		return BadgeLevelDanger
	}
}

// itoa converts an int to its decimal string representation without importing
// strconv at the call site.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := make([]byte, 0, 10)
	if n < 0 {
		buf = append(buf, '-')
		n = -n
	}
	var tmp [10]byte
	i := len(tmp)
	for n > 0 {
		i--
		tmp[i] = byte('0' + n%10)
		n /= 10
	}
	buf = append(buf, tmp[i:]...)
	return string(buf)
}
