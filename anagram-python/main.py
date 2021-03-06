#! /usr/bin/python3
from anagrampkg.anagramlib import apply_anagram, top_finder

def main():  
    try:         
        with open("wordlist.txt", 'r', encoding = "ISO-8859-1") as \
            input_file, open("output.txt", 'w', encoding = "ISO-8859-1") \
            as output_file:
                
            input_list = []
            for line in input_file:
                input_list.append(line.strip().lower())
            input_list = list(set(input_list))
            ordered_dict = apply_anagram(input_list)

            # wrting results to the file
            for key in ordered_dict:
                if len(ordered_dict[key]) > 1:             
                    output_file.write(" ".join(ordered_dict[key]))
                    output_file.write("\n")    
            
            # finding winners
            result = top_finder(ordered_dict)          
            print ("our winner is ", ordered_dict[result["most_words"]])
            print ("the biggest words that are anagram ", ordered_dict[result["biggest_words"]])
            
    except FileNotFoundError:
        print("File does not exist")
    except:
        print("unexpected error related to the file")
    

if __name__ == '__main__':
    main()
