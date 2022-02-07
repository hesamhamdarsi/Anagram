def apply_anagram(primary_list):
    final_list = []
    sorted_dict = {}
    try:
        for word in primary_list:
            # to sort every word of the list alphabetic , e.g. act -> ['c','a,'t']
            temp = sorted(word) 
            join_factor = ""
            temp = join_factor.join(temp) 
            # create/update dictionary with keys that are "sorted word"
            sorted_dict.setdefault(temp,[]).append(word)
    except:
        print("cannot read the list")
    return sorted_dict


def top_finder(ordered_dict):
    most_words = ""
    max_len = 0
    biggest_words = ""
    max_chars = 0
    for key in ordered_dict:      
        if len(ordered_dict[key]) > 1:
            # find the anagram with the most words 
            if len(ordered_dict[key]) > max_len:
                max_len = len(ordered_dict[key])
                most_words = key            
            # find the anagram with the biggest words 
            if len(key) > max_chars:
                max_chars = len(key)  
                biggest_words = key
    
    return {"most_words": most_words , "biggest_words": biggest_words}
