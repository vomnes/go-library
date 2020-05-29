package lib

import (
	"reflect"
	"strings"
)

// RemoveCharacters clear the selected characters from a string
func RemoveCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}

// StringInArray take a string or an array of strings and a array of string as parameter
// If the first argument is an array of string, the function return true if
// at list one of the elements array in the array of the second argument
// Return true if the string is in the array of string else false
func StringInArray(a interface{}, list []string) bool {
	var elements []string
	typeA := reflect.TypeOf(a)
	if typeA.String() == "string" {
		elements = []string{a.(string)}
	}
	if typeA.String() == "[]string" {
		elements = a.([]string)
	}
	if len(elements) == 0 {
		return false
	}
	for _, b := range list {
		for _, elem := range elements {
			if b == elem {
				return true
			}
		}
	}
	return false
}

func Strsub(input string, start int, end int) string {
	var output string
	if start < 0 || end < 0 {
		return ""
	}
	for i := start; i < start+end; i++ {
		output += string(input[i])
	}
	return output
}

func TrimStringFromString(s, sub string) string {
	if idx := strings.Index(s, sub); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimStringFromLastString(s, sub string) string {
	if idx := strings.LastIndexAny(s, sub); idx != -1 {
		return s[:idx]
	}
	return s
}
