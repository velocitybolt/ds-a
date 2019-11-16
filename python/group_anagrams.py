
'''
Match anagrams to their counterparts given a list

input => ["eat", "tea", "ate", "nat", "bat"]

output => [
    ["ate","eat","tea"]
    ["nat","tan"]
    ["bat"]
]
'''
# O(n * n lg n)
def group_anagrams(test):
   # Loop through anagrams input
   # Sort the anagrams from input and store in variable
   # Check if the current anagram is in the dictionary else return empty list
   # since the first combination wont exist in the dictionary it automatically returns a list which can be appended to?
   # append all new combinations
   # set the respective sorted anagram to its place in the dictionary 
   # return the values of the dictionary 
   gram_dict = {}

   for i in range(len(test)):
       sorted_grams = "".join(sorted(test[i]))
       sorted_grams_result = gram_dict.get(sorted_grams, [])
       sorted_grams_result.append(test[i])
       #print(sorted_grams_result)
       gram_dict[sorted_grams] = sorted_grams_result
       #print(gram_dict)
   return list(gram_dict.values()) 
       


print(group_anagrams(["eat", "tea", "ate", "nat", "bat"]))