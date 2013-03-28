package timetext

import (
	"fmt"
	"strings"
)

func Duration(delta int64) string {
	const year_text = 'y'
	const day_text = 'd'
	const hour_text = 'h'
	const minute_text = 'm'
	const second_text = 's'

	const SECONDS_PER_MINUTE = 60
	const SECONDS_PER_HOUR = 60 * SECONDS_PER_MINUTE
	const SECONDS_PER_DAY = 24 * SECONDS_PER_HOUR
	const SECONDS_PER_YEAR = 365 * SECONDS_PER_DAY

	if delta == 0 {
		return "0s"
	}

	years := delta / SECONDS_PER_YEAR
	delta = delta % SECONDS_PER_YEAR

	days := delta / SECONDS_PER_DAY
	delta = delta % SECONDS_PER_DAY

	hours := delta / SECONDS_PER_HOUR
	delta = delta % SECONDS_PER_HOUR

	minutes := delta / SECONDS_PER_MINUTE
	delta = delta % SECONDS_PER_MINUTE

	seconds := delta

	var timeChunk [5]string
	idx := 0

	if years > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%c", years, year_text)
		idx++
	}
	if idx > 0 || days > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%c", days, day_text)
		idx++
	}
	if idx > 0 || hours > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%c", hours, hour_text)
		idx++
	}
	if idx > 0 || minutes > 0 {
		timeChunk[idx] = fmt.Sprintf("%2d%c", minutes, minute_text)
		idx++
	}
	if idx > 0 || seconds > 0 {
		fmtStr := "%2d%c"
		if idx == 0 {
			fmtStr = "%d%c"
		}
		timeChunk[idx] = fmt.Sprintf(fmtStr, seconds, second_text)
		idx++
	}

	return strings.Join(timeChunk[0:idx], " ")
}
