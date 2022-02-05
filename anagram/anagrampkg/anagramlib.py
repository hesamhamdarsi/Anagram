from itertools import groupby

def apply_anagram(mylist):
    try:
        # final_list = [list(group) for key,group in groupby(sorted(mylist,key=sorted),sorted)]
        
        
        final_list = []
        mylist2 = []
        mydict = {}
        for i in mylist:
            x = sorted(i)  # to sort every single word in the last based on the alphabets , e.g. act -> ['c','a,'t']
            join_factor = ""
            x = join_factor.join(x)
            mydict.setdefault(x,[]).append(i)
        
        
    except:
        print("cannot read the list")
    return mydict

# unit test
# integreration test -> test multiple component that works correctly together
# assertion -> checking a function that is working correctly. test a single component 

# collection library



# max size of the list is based on the cpu arch 32/54. to calculate you can use "print sys.maxsize"

# UnicodeDecodeError: 'utf-8' codec can't decode byte 0xe8 in position 7973 --> file reading error

# https://docs.python.org/3/library/itertools.html#itertools.groupby
