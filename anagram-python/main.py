#! /usr/bin/python3
from anagrampkg.anagramlib import apply_anagram

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
            
            # vars to find the anagram containing the most words 
            most_words = ""
            max_len = 0
            # vars to find the biggest words that are anagram
            biggest_words = ""
            max_chars = 0
            
            for key in ordered_dict:
                if len(ordered_dict[key]) > 1:
                    
                    # find the anagram with the most words 
                    if len(ordered_dict[key]) > max_len:
                        max_len = len(ordered_dict[key])
                        most_words = key  
                    if len(key) > max_chars:
                        max_chars = len(key)  
                        biggest_words = key               
                    output_file.write(" ".join(ordered_dict[key]))
                    output_file.write("\n")                  
            print ("our winner is ", ordered_dict[most_words])
            print ("the biggest words that are anagram ", ordered_dict[biggest_words])
            
    except FileNotFoundError:
        print("File does not exist")
    except:
        print("unexpected error related to the file")

# def top_finder():
    

if __name__ == '__main__':
    main()
