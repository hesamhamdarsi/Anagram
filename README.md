# TOC
### Objective
### solutions
#### Python
#### Golang
### test results 


### Python
there are multiple methods that enable us to write a bunch of codes for sorting a file in a way 
> - Nested for loop
> - Python read-to-go libraries
> - Dictionary and sort

as the nested loops are not the proper way, I would have two options. either using available python libraries or wrtie a spesific code for this target. using some liberaries like [GroupBy](https://docs.python.org/3/library/itertools.html#itertools.groupby) I didn't get the fastest execution time that is possible. that's wht I choosed to write my own block of codes.

the process starts with the [main](/anagram-python/main.py) package where we I read the input file and convert it to a list of strings. remove the duplicated words from it, and I am converting all words to lower case to make sense during the comparison. then this list will be sent to the [Anagram package](/anagram-python/anagrampkg/anagramlib.py). 

the following lines show the main part of the Anagram packge:

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

here I just used to get the list of words, sort each one based on alphabet, re-assemble them as a word and use this sorted word as the key. then any word that is compatible with this sorted key, will be appended to the dictionary (the key will be created if not exist)

then the I am adding the output to a file in the main module and remove those that are not anagram.

### Golang
one benefit of using golang or any multi thread PL is using multiple services to get the job done simulteseusly 
this approach is ok if the scrpit includes more CPU process than IO process
there are different approachs for writing this spesific progeam with python using goroutins and channels in Golang
> - WaitForResult
> - WaitForTask
> - Fanout
> - Fanout-Bounded

I have choosen to write the code using **Fanout-Bounded** algorithm. this way, I have a buffered channel and the size of the buffer is limited to the number of threads my operating system can provides
using this method I don't have to create a lot of go routins and so wait to the cleaning cycle each time to free them up when I need to run this code multiple times 
I also included the heart of the program in a package so that I can test it well as well as make the program more readable and exposing the package anyway

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

I have tried to chose variables that are kind of clear what they are doing but breifly this block if code is getting each word in the list, send that to a queue and then multiple goroutins (depend on the OS threads) will pick the word from the buffer up and then send them to the calculation function. 

the calculation function then sorts the words based on their Ascci code and compare them with create/update a **Map** with that the same *sorted word* and will append the items that are anagram of eachother as the values of this spesific key.

then the unique function will check for any dupplicated value and remove them from the Map.

next will be writing this customized values into the output file in the main module.

### test results
I used multiple test files in python and golang for either testing the functionality of the core module(Anagram) using unit test and  benchmark to compare it with the other available solutions as well as compare it with the golang version of deployment.

regarding the benchmark the comparison between my application in golang and python is as bellow:

> Operating system: Mac

> total number of CPU cores: 8

> number of run: 6

> result:
> - Python : 2.202525291 s 
> - Golang: 2.2 s 

so both of them had almost the same execution time but golang is a little faster if we just increase the number of test loops 
