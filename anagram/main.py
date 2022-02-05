from anagrampkg.anagramlib import apply_anagram

def main():  
    try:         
        with open("wordlist.txt", 'r', encoding = "ISO-8859-1") as this, open("output.txt", 'w', encoding = "ISO-8859-1") as that:
            mylist = []
            for line in this:
                mylist.append(line.strip().lower())
            mylist = list(set(mylist))
            ordered_dict = apply_anagram(mylist)
            for key in ordered_dict:
                if len(ordered_dict[key]) > 1:
                    that.write(" ".join(ordered_dict[key]))
                    that.write("\n")
    except FileNotFoundError:
        print("File does not exist")
    except:
        print("unexpected error related to the file")

if __name__ == '__main__':
    main()
