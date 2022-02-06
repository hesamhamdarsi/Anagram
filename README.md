# TOC
## Objective
## solutions
### Python
### Golang
## Pros and Cons

### Python
there are multiple methods that enable us to write a bunch of codes for sorting a file in a way 
> - Nested for loop
> - Python read-to-go libraries
> - Dictionary and sort

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
the following lines show the main part of the packge:

,,,python
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
,,,

I have tried to chose variables that are kind of readble 

