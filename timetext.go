package timetext

import (
	"fmt"
	"strings"
)

const (
	short = iota
	long
)

func TerseDuration(delta int64) string {
	return duration(delta, short, true)
}

func TerseLongDuration(delta int64) string {
	return duration(delta, long, true)
}

func Duration(delta int64) string {
	return duration(delta, short, false)
}

func LongDuration(delta int64) string {
	return duration(delta, long, false)
}

func duration(delta, style int64, terse bool) string {
	year_text := "y"
	week_text := "w"
	day_text := "d"
	hour_text := "h"
	minute_text := "m"
	second_text := "s"

	if style == long {
		year_text = " year"
		week_text = " week"
		day_text = " day"
		hour_text = " hour"
		minute_text = " minute"
		second_text = " second"
	}

	const SECONDS_PER_MINUTE = 60
	const SECONDS_PER_HOUR = 60 * SECONDS_PER_MINUTE
	const SECONDS_PER_DAY = 24 * SECONDS_PER_HOUR
	const SECONDS_PER_WEEK = 7 * SECONDS_PER_DAY
	const SECONDS_PER_YEAR = 365 * SECONDS_PER_DAY

	if delta == 0 {
		switch style {
		case short:
			return "0s"
		case long:
			return "0 seconds"
		}
	}

	years := delta / SECONDS_PER_YEAR
	delta = delta % SECONDS_PER_YEAR

	weeks := delta / SECONDS_PER_WEEK
	delta = delta % SECONDS_PER_WEEK

	days := delta / SECONDS_PER_DAY
	delta = delta % SECONDS_PER_DAY

	hours := delta / SECONDS_PER_HOUR
	delta = delta % SECONDS_PER_HOUR

	minutes := delta / SECONDS_PER_MINUTE
	delta = delta % SECONDS_PER_MINUTE

	seconds := delta

	const nChunks = 6 // years weeks days hours minutes seconds
	var timeChunk [nChunks]string
	idx := 0

	if years > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%v", years, year_text)
		if style == long && years > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}
	if (!terse && idx > 0) || weeks > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%v", weeks, week_text)
		if style == long && weeks > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}
	if (!terse && idx > 0) || days > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%v", days, day_text)
		if style == long && days > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}
	if (!terse && idx > 0) || hours > 0 {
		timeChunk[idx] = fmt.Sprintf("%d%v", hours, hour_text)
		if style == long && hours > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}
	if (!terse && idx > 0) || minutes > 0 {
		timeChunk[idx] = fmt.Sprintf("%2d%v", minutes, minute_text)
		if style == long && minutes > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}
	if (!terse && idx > 0) || seconds > 0 {
		fmtStr := "%2d%v"
		if idx == 0 {
			fmtStr = "%d%v"
		}
		timeChunk[idx] = fmt.Sprintf(fmtStr, seconds, second_text)
		if style == long && seconds > 1 {
			timeChunk[idx] += "s"
		}
		idx++
	}

	return strings.Join(timeChunk[0:idx], " ")
}
