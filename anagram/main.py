mylist = []
mydict = {}
mylist4 = []
with open("myfile.txt", 'r', encoding = "ISO-8859-1") as this:
    for line in this:
        mylist.append(line.strip().lower())
    mylist.sort(key=len)
    mylist2 = mylist.copy()
    i = 0
    for m1 in mylist:
        mylist3 = []
        print(i)
        for m2 in mylist2:
            if len(m2) != len(m1):
                break
            if sorted(m1) == sorted(m2):
                mylist3.append(m2)
        #print("2222222")
        if len(mylist3) != 0:
            # mydict.update({m1: mylist3})
            mylist4.append(mylist3)
        #print("3333333")
        for item in mylist3:
            mylist2.remove(item)
        i += 1
    print("44444")
    print(mylist4)
        #print(mylist2)
    #print("the largest = ", mylist[-1], "and it's dolingos are ", mydict[mylist[-1]])
    #print(mydict)
    
with open("myfile2.txt", 'w', encoding = "ISO-8859-1") as that:
    for item in mylist4:
        for i in item:
            that.write(i)
            that.write(" ")
        that.write("\n")
        #print(i)





# max size of the list is based on the cpu arch 32/54. to calculate you can use "print sys.maxsize"

# UnicodeDecodeError: 'utf-8' codec can't decode byte 0xe8 in position 7973 --> file reading error
