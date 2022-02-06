package anagrampkg

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnagramFanout(t *testing.T) {
	x := make(map[string][]string)
	result := map[string][]string{
		"aciot": {"tacio"},
		"act":   {"tac", "act", "cat"},
		"actx":  {"catx", "xcat", "atxc"},
		"aimn":  {"aimn", "iamn", "main"},
		"cotx":  {"xcot"},
	}
	originalList := []string{"tacio", "tac", "act", "cat", "catx", "xcat", "atxc", "aimn", "iamn", "main", "xcot"}
	// given section
	t.Log("Anagram function should return a map that includes all anagram words as a value of the sorted anagram key")
	{
		t.Log("\ttest 0: check if the output of ", originalList, " is: ", result)
		{
			x = AnagramFanout(originalList)
			for key, value := range x {
				sort.Strings(value) // sort the slice alphabetic before comparing
				sort.Strings(result[key])
				if reflect.DeepEqual(value, result[key]) == false {
					t.Error("Expected", result[key], "got", value)
				}
			}

		}
	}
}
