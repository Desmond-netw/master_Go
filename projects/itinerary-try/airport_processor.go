package main

import (
	"regexp"
)

func processAirportCodes(content string, airportData *AirportData) string {
	// Patterns for airport codes
	iataPattern := `#([A-Z]{3})`
	icaoPattern := `##([A-Z]{4})`
	cityIATAPattern := `\*#([A-Z]{3})`
	cityICAOPattern := `\*##([A-Z]{4})`

	// Process city names first (with * prefix)
	content = regexp.MustCompile(cityIATAPattern).ReplaceAllStringFunc(content, func(match string) string {
		code := match[2:] // Remove "*#"
		if info, exists := airportData.ByIATA[code]; exists {
			return info.Municipality
		}
		return match
	})

	content = regexp.MustCompile(cityICAOPattern).ReplaceAllStringFunc(content, func(match string) string {
		code := match[3:] // Remove "*##"
		if info, exists := airportData.ByICAO[code]; exists {
			return info.Municipality
		}
		return match
	})

	// Process airport names
	content = regexp.MustCompile(iataPattern).ReplaceAllStringFunc(content, func(match string) string {
		code := match[1:] // Remove "#"
		if info, exists := airportData.ByIATA[code]; exists {
			return info.Name
		}
		return match
	})

	content = regexp.MustCompile(icaoPattern).ReplaceAllStringFunc(content, func(match string) string {
		code := match[2:] // Remove "##"
		if info, exists := airportData.ByICAO[code]; exists {
			return info.Name
		}
		return match
	})

	return content
}
