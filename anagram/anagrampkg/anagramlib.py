from itertools import groupby

def apply_anagram(mylist):
    try:
        final_list = [list(group) for key,group in groupby(sorted(mylist,key=sorted),sorted)]
    except:
        print("cannot read the list")
    return final_list

# unit test
# integreration test -> test multiple component that works correctly together
# assertion -> checking a function that is working correctly. test a single component 

# collection library



# max size of the list is based on the cpu arch 32/54. to calculate you can use "print sys.maxsize"

# UnicodeDecodeError: 'utf-8' codec can't decode byte 0xe8 in position 7973 --> file reading error
