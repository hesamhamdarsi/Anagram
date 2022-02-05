from itertools import groupby

def apply_anagram(primary_list):
    try:
        final_list = []
        sorted_dict = {}
        for word in primary_list:
            temp = sorted(word)  # to sort every single word in the last based on the alphabets , e.g. act -> ['c','a,'t']
            join_factor = ""
            temp = join_factor.join(temp)
            sorted_dict.setdefault(temp,[]).append(word)
    except:
        print("cannot read the list")
    return sorted_dict
