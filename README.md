### Objective
### solutions
- #### Python
- #### Golang
### Test results

#

### Objective:
Wrting a program to do the following tasks: 
> - pick a file including any number of words, look for the words that are anagram and print them in the output file.
> - print the largest set of words that are anagram
> - print the biggest words that are anagram
```diff
- Notice: the output should not be included with duplicated items as well as words that are not anagram.
``` 

### Python
There are multiple ways to write such a program leveraging python language
> - Nested loops
> - Python ready-to-go libraries
> - Dictionary and sort


Using nested loops is a bad practice as most of the time it overkills the resources, especially when we're working on huge numbers. So we have two more choices. either using available python libraries or wrtie a specific module for this target. 

Using the available python liberaries (e.g. [GroupBy](https://docs.python.org/3/library/itertools.html#itertools.groupby)) were not the most efficient way when I have tested them (the result took almost two times more time than my own solution).

#### the final approach:
The process starts with the [main](/anagram-python/main.py) package where we inject a file as an input (including more than 300,000 words in each line) and then we convert it to a list of strings. 

For the next step, we remove the duplicated words from the list, and we convert all of the words to lower case to prevent false positive for the next rounds. then the final list will be sent to the [Anagram package](/anagram-python/anagrampkg/anagramlib.py). 

The following lines represent the main module of the [Anagram package](/anagram-python/anagrampkg/anagramlib.py) that is responsible to find anagrams, sort the words based on a *sorted_key* and then return a sorted dictionary:

```python
try:
    final_list = []
    sorted_dict = {}
    for word in primary_list:
        # to sort every word of the list alphabetic , e.g. act -> ['c','a,'t']
        temp = sorted(word) 
        join_factor = ""
        temp = join_factor.join(temp) 
        # create/update dictionary with keys that are "sorted word"
        sorted_dict.setdefault(temp,[]).append(word)
except:
    print("cannot read the list")
return sorted_dict
```

> - here I just grab a list of words, sort them alphabetically, re-assemble them as a word and use this sorted word as a key in a dictionary. then any word that is matched with this sorted key, will be appended to the dictionary (the key will be created if not exist)
> - at the end, the output will be written to a file in the main package and we remove the words that are not anagram.

The other part of our Anagram package is a module to find the largest words that are anagram as well as the biggest one.
let's take a look at this part:

```python
def top_finder(ordered_dict):
    most_words = ""
    max_len = 0
    biggest_words = ""
    max_chars = 0
    for key in ordered_dict:      
        if len(ordered_dict[key]) > 1:
            # find the anagram with the most words 
            if len(ordered_dict[key]) > max_len:
                max_len = len(ordered_dict[key])
                most_words = key            
            # find the anagram with the biggest words 
            if len(key) > max_chars:
                max_chars = len(key)  
                biggest_words = key
    
    return {"most_words": most_words , "biggest_words": biggest_words}
```
What I do in this block of code is simple, I just go through the sorted dictionary and pick two keys:
> - the key that it's value has the longest length (a list that has the maximum members)
> - the key that is the biggest key between the others (as the values are anagram, the key length and values length are equal)

#

### Golang
One benefit of using golang or any **real** multi threading PL is leveraging all threads of the OS(on top of the hardware) to get the job done simultaneously and fast

This approach is ok if our scrpit includes more CPU process than IO process

There are different approachs for writing this spesific program using golang and channels in Golang:
> - WaitForResult
> - WaitForTask
> - Fanout
> - Fanout-Bounded
> - ...

I have decided to write the code using **Fanout-Bounded** algorithm. this way, I have a buffered channel and the size of the buffer is limited to the number of threads that my operating system can provides

Using this method I don't have to create a lot of go routins and wait for the cleaning cycle to terminate them especially if we need to run this code multiple times without delay 

The heart of the program is included in a package so that I can test it. this way we also make the program more readable and we're able to expose our package

The process starts with the [main](/anagram-golang/main.go) package where we read the input file and convert it to a slice of strings. then this slice will be sent to the [Anagram package](/anagram-golang/anagrampkg/anagrampkg.go#L9). 

The following lines show the main part of the Anagram packge:

```go
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
```

```diff
- I have tried to use convenient naming for variables to declare what they are doing
``` 

Let's see what it is doing:

> - this block of code is taking the words one by one from the slice, send them to a queue(channel's buffer).
> - then multiple goroutins (depend on the OS threads) will pick the word from the buffer up and then send them to the calculation function. 
> - the calculation function then sorts the words based on their ascci code and compare them with the available values of the key. if there is no key available, then a new key will be created.
> - we are actually creating/updating a **Map** using *sorted words* while *sorted words* prevent MAP to have extra keys. the the values will be append to the spesific key and we will have a list of anagram as a value for that key.
> - then the unique function will check for any dupplicated value and remove them from the Map.
> - next will be writing this customized values into the output file in the main module.

The other part of our Anagram package is a function to find the largest words that are anagram as well as the biggest one.
let's take a look at this part:

```go
func TopFinder(sortedMap map[string][]string) map[string]string {
	mostWords := ""
	maxLen := 0
	var biggestWords string
	result := make(map[string]string)
	maxChars := 0
	for key, value := range sortedMap {
		// find the anagram with the most words
		if len(value) > maxLen {
			maxLen = len(value)
			mostWords = key
		}
		// find the anagram with the biggest words
		if len(key) > maxChars {
			maxChars = len(key)
			biggestWords = key
		}
	}
	result["mostWords"] = mostWords
	result["biggestWords"] = biggestWords
	return result
}
```
What I do in this block of code is simple, I just go through a sorted Map and pick two keys:
> - the key that it's value has the longest length (a list that has the maximum members)
> - the key that is the biggest key between the others (as the values are anagram, the key length and values length are equal)


### test results
There are a couble of test scripts in anagram-python and anargram-golang directories for either doing tunit test of the Anagram package or benchmark to compare the excution time between **Golang** and **Python**

Regarding the benchmark the output is:
```s
Operating system: Mac
total number of CPU cores: 8
number of run: 6
result:
> - Python : 2.202525291 s 
> - Golang: 2.2 s 
```

As a result, both Golang and Python has almost the same execution time but golang is a little bit faster if we just increase the number of loops in the test. 
