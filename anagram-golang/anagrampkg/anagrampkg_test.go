package anagrampkg

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnagramFanout(t *testing.T) {
	x := make(map[string][]string)
	result := map[string][]string{
		"act":  {"tac", "act", "cat"},
		"actx": {"catx", "xcat", "atxc"},
		"aimn": {"aimn", "iamn", "main"},
	}
	originalList := []string{"tacio", "tac", "act", "cat", "catx",
		"xcat", "atxc", "aimn", "iamn", "main", "xcot"}
	// given section
	t.Log(`Anagram function should return a map that includes 
	all anagram words as a value of the sorted anagram key`)
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

func TestTopFinder(t *testing.T) {
	x := make(map[string]string)
	result := map[string]string{
		"biggestWords": "deglnoorstw",
		"mostWords":    "actx",
	}
	originalMap := map[string][]string{ // map[key_type]value_type
		"aciot":       {"tacio", "tacoi"},
		"act":         {"cat", "tac"},
		"actx":        {"atxc", "xcat", "catx"},
		"aimn":        {"main", "aimn"},
		"deglnoorstw": {"dongestworl", "longestword"},
	}
	// given section
	t.Log(`TopFinder function should return a map that includes
	the keys of the longets and biggest anagrams`)
	{
		t.Log("\ttest 0: check if the output of ", originalMap, " is: ", result)
		{
			x = TopFinder(originalMap)
			for key, value := range x {
				if reflect.DeepEqual(value, result[key]) == false {
					t.Error("Expected", result[key], "got", value)
				}
			}

		}
	}
}
