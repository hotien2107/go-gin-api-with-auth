package utils

import (
	"regexp"
	"strings"
)

// Function check string has contain number
func IsContainNumber(str string) bool {
	digitRegex := regexp.MustCompile(`[0-9]`)
	return digitRegex.MatchString(str)
}

// Function check string has contain capital letter
func IsContainCapitalLetter(str string) bool {
	capitalRegex := regexp.MustCompile(`[A-Z]`)
	return capitalRegex.MatchString(str)
}

// Function check string has contain space
func IsContainSpace(str string) bool {
	spaceRegex := regexp.MustCompile(" ")
	return spaceRegex.FindAllIndex([]byte(str), -1) != nil
}

// Function check string is empty
func IsEmpty(str string) bool {
	strFormat := strings.TrimSpace(str)

	if len(strFormat) == 0 || strFormat == "" {
		return true
	}
	return false
}
