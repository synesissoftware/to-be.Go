package to_be

import (
	"slices"
	"strings"
)

var _FALSEY_PRECISE_STRINGS = [...]string{
	"0",
	"FALSE",
	"False",
	"NO",
	"No",
	"OFF",
	"Off",
	"false",
	"no",
	"off",
}

var _TRUEY_PRECISE_STRINGS = [...]string{
	"1",
	"ON",
	"On",
	"TRUE",
	"True",
	"YES",
	"Yes",
	"on",
	"true",
	"yes",
}

var _FALSEY_LOWERCASE_STRINGS = [...]string{
	"false",
	"no",
	"off",
	"0",
}

var _TRUEY_LOWERCASE_STRINGS = [...]string{
	"true",
	"yes",
	"on",
	"1",
}

func StringIsFalsey(s string) bool {
	if _, found := slices.BinarySearch(_FALSEY_PRECISE_STRINGS[:], s); found {
		return true
	}

	l := strings.ToLower(strings.TrimSpace(s))

	if slices.Contains(_FALSEY_LOWERCASE_STRINGS[:], l) {
		return true
	}

	return false
}

func StringIsTruey(s string) bool {
	if _, found := slices.BinarySearch(_TRUEY_PRECISE_STRINGS[:], s); found {
		return true
	}

	l := strings.ToLower(strings.TrimSpace(s))

	if slices.Contains(_TRUEY_LOWERCASE_STRINGS[:], l) {
		return true
	}

	return false
}

func StringIsTruthy(s string) bool {
	if _, found := slices.BinarySearch(_TRUEY_PRECISE_STRINGS[:], s); found {
		return true
	}

	if _, found := slices.BinarySearch(_FALSEY_PRECISE_STRINGS[:], s); found {
		return true
	}

	l := strings.ToLower(strings.TrimSpace(s))

	if slices.Contains(_TRUEY_LOWERCASE_STRINGS[:], l) {
		return true
	}

	if slices.Contains(_FALSEY_LOWERCASE_STRINGS[:], l) {
		return true
	}

	return false
}
