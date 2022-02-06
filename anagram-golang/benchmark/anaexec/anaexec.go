package anaexec

import (
	"anagram/python/test/anagram-golang/anagrampkg"
	"bufio"
	"log"
	"os"
	"strings"
)

func Anaexec() {

	// Read the file, breake the lines and append each line as a value in the slice
	originalList := sliceFile("../wordlist.txt")

	// calculate the anagram
	output := anagrampkg.AnagramFanout(originalList)

	// Create/truncate the output file
	outputFile := CreateFile("output.txt")
	defer outputFile.Close()

	// Write the anagram output to the output file
	for _, value := range output {
		WriteFile(value, outputFile)
	}
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

func WriteFile(word []string, file *os.File) {
	line := strings.Join(word[:], " ")
	if len(word) > 1 {
		_, err2 := file.WriteString(line + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
