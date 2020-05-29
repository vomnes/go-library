package lib

import (
	"testing"
)

var RemoveCharactersTests = []struct {
	toClear     string // input
	characters  string // input
	expected    string // expected result
	testContent string // test details
}{
	{" +33 6 00 00 00 ", " ", "+336000000", "Only spaces to clear"},
	{"+33 6 00 00 00", " +", "336000000", "Spaces and '+' to clear"},
	{"+33 6", "", "+33 6", "Nothing to clear"},
}

func TestRemoveCharacters(t *testing.T) {
	for _, tt := range RemoveCharactersTests {
		actual := RemoveCharacters(tt.toClear, tt.characters)

		if actual != tt.expected {
			t.Errorf("RemoveCharacters(%s, %s): expected %s, actual %s - Test type: \033[31m%s\033[0m", tt.toClear, tt.characters, tt.expected, actual, tt.testContent)
		}
	}
}

var stringInArrayTests = []struct {
	item_s      interface{} // input
	stringArray []string
	expected    bool   // expected result
	testContent string // test details
}{
	{"a", []string{"a", "b", "c"}, true, "Basic - string"},
	{"z", []string{"a", "b", "c"}, false, "Not in array"},
	{42, []string{"a", "b", "c"}, false, "Not a string or an array"},
	{[]string{"a"}, []string{"a", "b", "c"}, true, "Basic - []string"},
	{[]string{"z", "a"}, []string{"a", "b", "c"}, true, "Basic - []string"},
	{[]string{"a", "z"}, []string{"a", "b", "c"}, true, "Basic - []string"},
	{[]string{"y", "z"}, []string{"a", "b", "c"}, false, "Not in array"},
	{[]string{"y"}, []string{"a", "b", "c"}, false, "Not in array"},
}

func TestStringInArray(t *testing.T) {
	for _, tt := range stringInArrayTests {
		actual := StringInArray(tt.item_s, tt.stringArray)
		if actual != tt.expected {
			t.Errorf("StringInArray(%s, %s): expected %t, actual %t - Test type: \033[31m%s\033[0m", tt.item_s, tt.stringArray, tt.expected, actual, tt.testContent)
		}
	}
}
