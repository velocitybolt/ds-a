'''
Write a function that takes in two non-empty arrays of ints. 

The function should find the pair of numbers whose absolute difference is closest to zero.

[-1,5,10,20,28,3], [26,134,135,15,17]

returns [28,26] => find pair smallest diff (2)

'''

def smallestDifference(arr_one, arr_two):
    arr_one.sort()
    arr_two.sort()

    """ arr_one [-1, 3, 5, 10, 20, 28]  arr_two [15, 17, 26, 134, 135] """
    # sort them because it allows you to make the sum closer to zero by interchanging values

    smallest_diff = float("inf")
    current_diff = float("inf")

    pair = []

    left = 0
    right = 0

    while left < len(arr_one) and right < len(arr_two):
        leftVal = arr_one[left]
        rightVal = arr_two[right]

        if leftVal < rightVal:
            left += 1
            current_diff = rightVal - leftVal

        elif rightVal < leftVal:
            right += 1
            current_diff = leftVal - rightVal
             
        else: # they are equal
            return [leftVal, rightVal]

        # check if current diff is smallest_diff

        if smallest_diff > current_diff:
            smallest_diff = current_diff
            pair = [leftVal, rightVal]
    
    return pair 


print(smallestDifference([-1, 5, 10, 20, 28, 3], [26, 134, 135, 15, 17]))
print(smallestDifference([-1, 5, 10, 20, 28, 3], [26, 134, 135, 15, 17]))







