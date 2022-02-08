package main

import (
	"anagram/python/test/anagram-golang/anagrampkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Read the file, breake the lines and append each line as a value in the slice
	originalList := sliceFile("wordlist.txt")

	// calculate the anagram
	output := anagrampkg.AnagramFanout(originalList)

	// Create/truncate the output file
	outputFile := CreateFile("output.txt")
	defer outputFile.Close()

	// Write the anagram output to the output file
	for _, value := range output {
		WriteFile(value, outputFile)
	}

	result := anagrampkg.TopFinder(output)
	fmt.Println("our winner is ", output[result["mostWords"]])
	fmt.Println("the biggest words that are anagram ", output[result["biggestWords"]])
}

func sliceFile(fileName string) []string {

	var word []string
	var temp_list []string

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		word = append(word, scanner.Text())
	}

	for _, eachline := range word {
		eachline = strings.ToLower(eachline)
		temp_list = append(temp_list, eachline)
	}
	return temp_list
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
	_, err2 := file.WriteString(line + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
}
