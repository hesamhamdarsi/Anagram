mylist = []
mydict = {}
mylist4 = []
with open("wordlist.txt", 'r', encoding = "ISO-8859-1") as this, open("myfile2.txt", 'w', encoding = "ISO-8859-1") as that:
    for line in this:
        mylist.append(line.strip().lower())
    mylist = list(set(mylist))
    mylist.sort(key=len)
    mylist2 = mylist.copy()
    i = 0
    for m1 in mylist:
        mylist3 = []
        #print(i)
        for m2 in mylist2:
            if len(m2) != len(m1):
                #print("break")
                break
            if sorted(m1) == sorted(m2):
                mylist3.append(m2)
                #print(m2)
                that.write(m2)
                that.write(" ")
        that.write("\n")
        for item in mylist3:
            mylist2.remove(item)
        #i += 1




# max size of the list is based on the cpu arch 32/54. to calculate you can use "print sys.maxsize"

# UnicodeDecodeError: 'utf-8' codec can't decode byte 0xe8 in position 7973 --> file reading error
