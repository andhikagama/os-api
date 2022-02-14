package utils

import (
	"strings"
)

func IsDuplicateError(err error) bool {
	if strings.Contains(err.Error(), "Error 1062") {
		return true
	}

	return false
}

func IsInsufficientError(err error) bool {
	if strings.Contains(err.Error(), "Error 1264") {
		return true
	}

	return false
}
