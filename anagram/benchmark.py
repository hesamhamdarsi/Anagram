import timeit

from anagrampkg.anagramlib import apply_anagram
from main import main

print(timeit.timeit(main, number= 6))
