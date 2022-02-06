package anagrampkg

import (
	"runtime"
	"sort"
	"sync"
)

func AnagramFanout(originalList []string) map[string][]string {

	// the number of simultaneous goroutins
	coreNum := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(coreNum)

	var mu sync.Mutex

	// creating a channel with a buffer size that is limited to the number of OS threads
	ch := make(chan string, coreNum)

	SortedMap := make(map[string][]string)

	for worker := 0; worker < coreNum; worker++ {
		go func(worker int) {
			defer wg.Done()
			for word := range ch {
				calculate(word, originalList, SortedMap, &mu)
			}
		}(worker)
	}

	for _, word := range originalList {
		ch <- word
	}

	close(ch)
	wg.Wait()

	for key, value := range SortedMap {
		SortedMap[key] = unique(value)
	}

	return SortedMap

}

func calculate(word string, originalList []string, mymap map[string][]string, mu *sync.Mutex) {
	wordAsci := []rune(word) //superset of asccis(all chars)
	var tempList []string
	sort.SliceStable(wordAsci, func(i, j int) bool { return wordAsci[i] < wordAsci[j] })
	mu.Lock()
	if myvalue, ok := mymap[string(wordAsci)]; ok {
		myvalue = append(myvalue, word)
		mymap[string(wordAsci)] = myvalue
	} else {
		tempList = append(tempList, word)
		mymap[string(wordAsci)] = tempList
	}
	mu.Unlock()
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	uniqList := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqList = append(uniqList, entry)
		}
	}
	return uniqList
}
