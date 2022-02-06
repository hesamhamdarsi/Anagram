package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
)

func main() {

	// Read the file, breake the lines and append each line as a value in the slice
	originalList := sliceFile("wordlist.txt")

	// calculate the anagram
	output := nagramFanout(originalList)

	// Create/truncate the output file
	outputFile := CreateFile("output.txt")
	defer outputFile.Close()

	// Write the anagram output to the output file
	for _, value := range output {
		line := strings.Join(value[:], " ")
		if len(value) > 1 {
			_, err2 := outputFile.WriteString(line + "\n")
			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}
}

func nagramFanout(originalList []string) map[string][]string {

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

func calculate(word string, originalList []string, mymap map[string][]string, mu *sync.Mutex) {
	wordAsci := []rune(word)
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

func sliceFile(fileName string) []string {

	var txtlines []string
	var list1 []string

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	for _, eachline := range txtlines {
		eachline = strings.ToLower(eachline)
		list1 = append(list1, eachline)
	}
	return list1
}

func closeFile(fileName *os.File) {
	err := fileName.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func CreateFile(fileName string) *os.File {
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return outputFile
}
