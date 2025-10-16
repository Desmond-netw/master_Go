package main

import (
	"fmt"
	"regexp"
	"time"
)

func processDatesAndTimes(content string) string {
	// Patterns for date and time formats
	datePattern := `D\(([^)]+)\)`
	time12Pattern := `T12\(([^)]+)\)`
	time24Pattern := `T24\(([^)]+)\)`

	// Process dates
	content = regexp.MustCompile(datePattern).ReplaceAllStringFunc(content, func(match string) string {
		dateStr := match[2 : len(match)-1] // Extract content inside D(...)
		formatted, err := formatDate(dateStr)
		if err != nil {
			return match
		}
		return formatted
	})

	// Process 12-hour times
	content = regexp.MustCompile(time12Pattern).ReplaceAllStringFunc(content, func(match string) string {
		timeStr := match[4 : len(match)-1] // Extract content inside T12(...)
		formatted, err := formatTime12(timeStr)
		if err != nil {
			return match
		}
		return formatted
	})

	// Process 24-hour times
	content = regexp.MustCompile(time24Pattern).ReplaceAllStringFunc(content, func(match string) string {
		timeStr := match[4 : len(match)-1] // Extract content inside T24(...)
		formatted, err := formatTime24(timeStr)
		if err != nil {
			return match
		}
		return formatted
	})

	return content
}

func formatDate(isoString string) (string, error) {
	// Try to parse the ISO 8601 format
	t, err := time.Parse(time.RFC3339, isoString)
	if err != nil {
		return "", err
	}

	// Format as DD-Mmm-YYYY
	return t.Format("02 Jan 2006"), nil
}

func formatTime12(isoString string) (string, error) {
	t, offset, err := parseTimeWithOffset(isoString)
	if err != nil {
		return "", err
	}

	// Format as 12-hour time with AM/PM
	timeStr := t.Format("3:04PM")
	return fmt.Sprintf("%s (%s)", timeStr, formatOffset(offset)), nil
}

func formatTime24(isoString string) (string, error) {
	t, offset, err := parseTimeWithOffset(isoString)
	if err != nil {
		return "", err
	}

	// Format as 24-hour time
	timeStr := t.Format("15:04")
	return fmt.Sprintf("%s (%s)", timeStr, formatOffset(offset)), nil
}

func parseTimeWithOffset(isoString string) (time.Time, string, error) {
	// Try to parse the ISO 8601 format
	t, err := time.Parse(time.RFC3339, isoString)
	if err != nil {
		return time.Time{}, "", err
	}

	// Extract offset
	_, offset := t.Zone()
	offsetStr := formatOffsetFromSeconds(offset)

	return t, offsetStr, nil
}

func formatOffset(offset string) string {
	if offset == "Z" {
		return "+00:00"
	}
	return offset
}

func formatOffsetFromSeconds(offsetSeconds int) string {
	if offsetSeconds == 0 {
		return "Z"
	}

	hours := offsetSeconds / 3600
	minutes := (offsetSeconds % 3600) / 60

	sign := "+"
	if hours < 0 {
		sign = "-"
		hours = -hours
		minutes = -minutes
	}

	return fmt.Sprintf("%s%02d:%02d", sign, hours, minutes)
}
