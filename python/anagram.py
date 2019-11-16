#Write a method to check if two strings are anagrams

# "abc" "cba" => true
# "cinema" "iceman" => true
# "bumblebee" "icecream" => false
# must have same length and same chars


def anagram(wordOne, wordTwo):
    wordOne.lower()
    wordTwo.lower()
    if len(wordOne) != len(wordTwo):
        return False 
    
    wordOneDict = {}
    wordTwoDict = {}

    for i in range(len(wordOne)):
        if wordOne[i] not in wordOneDict.keys():
            wordOneDict[wordOne[i]] = 1
        else:
            wordOneDict[wordOne[i]] += 1
    
    for i in range(len(wordTwo)):
        if wordTwo[i] not in wordTwoDict.keys():
            wordTwoDict[wordTwo[i]] = 1
        else:
            wordTwoDict[wordTwo[i]] += 1
    
    for key, value in wordOneDict.items():
        if wordTwoDict[key] != value:
            return False

    return True

print(anagram("abc", "cba"))
print(anagram("ccac", "ccacc"))
print(anagram("icemmanioo", "cinemmaioo"))
print(anagram("bumblebee", "icecream"))
print(anagram("", ""))

    