def apply_anagram(primary_list):
    try:
        final_list = []
        sorted_dict = {}
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
