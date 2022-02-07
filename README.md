### Objective
### solutions
- ### Python
- ### Golang
### Test results

#

### Objective:
Wrting a progeram to do the following tasks: 
> - pick a file including any number of words, look for the words that are anagram and print them in the output file.
> - print the largest set of words that are anagram
> - print the biggest words that are anagram
```diff
- Notice: the result file should not included duplicated items as well as words that are not anagram.
``` 

### Python
There are multiple ways to write a such program leveraging python language
> - Nested for loop
> - Python ready-to-go libraries
> - Dictionary and sort

Using the nested loops is bad practice because it most of the time overkill the resources, especially when we're working on huge numbers. So we have two more choices. either using available python libraries or wrtie a spesific code for this target. 

beside, using some liberaries like [GroupBy](https://docs.python.org/3/library/itertools.html#itertools.groupby) were not the most efficient way when I have test them (the result was almost two times more than my own solution).

##### the final approach:
the process starts with the [main](/anagram-python/main.py) package where we inject a file as input (including more than 300,000 words in each line) and then we convert it to a list of strings. 

for the next step, we remove the duplicated words from the list, and we convert all of the words to lower case to prevent false positive for the next of the rounds. then the final list will be sent to the [Anagram package](/anagram-python/anagrampkg/anagramlib.py). 

the following lines represent the main module of the [Anagram package](/anagram-python/anagrampkg/anagramlib.py) that is responsible to find anagrams, sort the words based on a *sorted_key* and then return a sorted dictionary:

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

here I just grab a list of words, sort them alphabetically, re-assemble them as a word and use this sorted word as a key in a dictionary. then any word that is matched with this sorted key, will be appended to the dictionary (the key will be created if not exist)

at the end, the output will be written to a file in the main package and we remove the words that are not anagram.

the other part of our Anagram package is a module to find the largest words that are anagram as well as the biggest one.
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
what I do in this block of code is simple, I just go through the sorted dictionary and pick two keys:
> - the key that it's value has the longest length (a list that has the maximum members)
> - the key that is the biggest key between the others (as the values are anagram, the key length and values length are equal)

### Golang
one benefit of using golang or any **real** multi threading PL is leveraging all threads of the OS(on top of the hardware) to get the job done simultaneously and fast
this approach is ok if our scrpit includes more CPU process than IO process

there are different approachs for writing this spesific program using golang and channels in Golang:
> - WaitForResult
> - WaitForTask
> - Fanout
> - Fanout-Bounded
> - ...

I have decided to write the code using **Fanout-Bounded** algorithm. this way, I have a buffered channel and the size of the buffer is limited to the number of threads that my operating system can provides

using this method I don't have to create a lot of go routins and wait for the cleaning cycle to terminate them especially if we need to run this code multiple times without deply 

The heart of the program is included in a package so that I can test it. this way we also make the program more readable and we're able to expose our package

the process starts with the [main](/anagram-golang/main.go) package where we I read the input file and convert it to a slice of strings. then this slice will be sent to the [Anagram package](/anagram-golang/anagrampkg/anagrampkg.go#L9). 

the following lines show the main part of the Anagram packge:

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

I have tried to use convenient naming for variables to declare what they are doing but let's see what it is doing:
> - this block of code is taking the words one by one from the slice, send them to a queue(channel's buffer).
> - then multiple goroutins (depend on the OS threads) will pick the word from the buffer up and then send them to the calculation function. 
> - the calculation function then sorts the words based on their Ascci code and compare them with the available values of the key. if there is no key available, then a new key will be created.
> - we are actually creating/updating a **Map** with using *sorted word* while *sorted words* prevents having unnesessary key. the the values will be append to the spesific key and we will have a list of anagram as a value for that key.
> - then the unique function will check for any dupplicated value and remove them from the Map.
> - next will be writing this customized values into the output file in the main module.

the other part of our Anagram package is a function to find the largest words that are anagram as well as the biggest one.
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
what I do in this block of code is simple, I just go through a sorted Map and pick two keys:
> - the key that it's value has the longest length (a list that has the maximum members)
> - the key that is the biggest key between the others (as the values are anagram, the key length and values length are equal)


### test results
there are a couble of test scripts in anagram-python and anargram-golang directories for either testing the unit test of the Anagram package as well as benchmark to compare the excution time between **Golang** and **Python**

regarding the benchmark the output is:
```s
Operating system: Mac
total number of CPU cores: 8
number of run: 6
result:
> - Python : 2.202525291 s 
> - Golang: 2.2 s 
```

so both of them had almost the same execution time but golang is a little faster if we just increase the number of test loops 
