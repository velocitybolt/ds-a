#Implement an algorithim to determine if a string has all unique characters

# "abc" => true
# "" => true
# "aabc" => false

# Optimal -> O(n)
def optimal_unique_chars(input):
    visited_chars = set()

    for i in range(len(input)):
        letter = input[i]

        if letter in visited_chars:
            return False

        visited_chars.add(letter)

    return True


print(optimal_unique_chars("abc"))
print(optimal_unique_chars(""))
print(optimal_unique_chars("aabc"))