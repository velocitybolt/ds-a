"""
Largest Range

Write a function that takes an array of integers and returns an array of length 2 
representing the largest range of numbers contained in that array. Range meaning
consecutive numbers like 1,2,3....

Input: [2,3,4,5,6]
Output: [2,6]
"""

def largest_range(arr):
    nums = {}
    largest_pair = []
    longest_length = 0

    # mark all as unvisited for now
    for num in arr:
        nums[num] = False
    
    # if the number hasnt been visited traverse
    for num in arr:
        if nums[num]:
            continue
        # mark the number as visited
        nums[num] = True
        # calculate 1 less and 1 greater than current number 
        # to check if those are in the the map 
        left = num - 1
        right = num + 1
        current_length = 1

        while left in nums:
            # Mark the left as visited, increment length
            # Decrement left to look for next number in dict
            nums[left] = True
            current_length += 1
            left -= 1

        while right in nums:
            # Mark the right number as visited, increment length,
            # Increment right to look for next number in dict
            nums[right] = True
            current_length += 1
            right += 1

        if current_length > longest_length:
            longest_length = current_length
            largest_pair = [left + 1, right - 1]

    return largest_pair

print(largest_range([2,3,4,5,6]))
print(largest_range([19, -1, 18, 17, 2, 10, 3, 12, 5, 16, 4, 11, 8, 7, 6, 15, 12, 12, 2, 1, 6, 13, 14]))