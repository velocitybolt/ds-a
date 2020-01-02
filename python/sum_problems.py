'''
Two number sum and three number sum problems
'''

def two_number_sum(input_array, target): 
    """ calculate the complement and search for it in array  O(N) runtime """
    pairs = {}
    for current_number in input_array:
        comp = target - current_number
        # check if complement is in pairs? duh lol
        if comp in pairs:
            return sorted([comp, current_number])

        else:
            pairs[current_number] = True
    
    return []

print(two_number_sum([3, 5, -4, 8, 11, 1, -1, 6], 10))


def three_number_sum(arr, target):
    """ [12, 3, 1, 2, -6, 5, 0, -8, -1] sum = 0 
        result = [-8, 3, 5], [-6, 1, 5], [-1, 0, 1] """

    arr.sort()
    result = []

    for num in range(len(arr) - 2):
        left = num + 1 # bc there is no number left of the first element
        right = len(arr) - 1
        
        while left < right:
            currsum = arr[num] + arr[left] + arr[right]
            
            # all ifs means every condition checked elifs checks the first one that passes
            if currsum == target:
                result.append( [ arr[num], arr[left], arr[right] ] )
                # pair found reset... 
                left += 1
                right -= 1
            
            elif currsum >= target:
                # you need to reduce the sum bc its greater than target
                # reduce right because that makes the sum smaller (after sort larger numbers on right)
                right -= 1

            elif currsum <= target:
                left += 1
    
    return result
        

print(three_number_sum([12, 3, 1, 2, -6, 5, 0, -8, -1, 5], 0))