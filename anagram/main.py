from anagrampkg.anagramlib import apply_anagram

def main():  
    try:         
        with open("wordlist.txt", 'r', encoding = "ISO-8859-1") as this, open("output.txt", 'w', encoding = "ISO-8859-1") as that:
            mylist = []
            for line in this:
                mylist.append(line.strip().lower())
            mylist = list(set(mylist))
            mylist.sort(key=len)
            ordered_list = apply_anagram(mylist)
            for item in ordered_list:
                if len(item) > 1:
                    that.write(" ".join(item))
                    that.write("\n")
    except FileNotFoundError:
        print("File does not exist")
    except:
        print("unexpected error related ti the file")

if __name__ == '__main__':
    main()
