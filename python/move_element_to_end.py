
'''
Given an arrayay of integers and a toMove integer. 
Write a function that moves all instances of toMove integer to end of the list

ex: [3,4,5,6,4,4,5] T: 4 => [3,5,6,5,4,4,4]
'''

def move_to_end(array, toMove):
    # two pointers on left and right of arrayay to go through list O(n)
    left = 0
    right = len(array) - 1

    while left < right:
        # need second while to keep moving in case [12,4,4,65,7,7,7,7,7,7,7,7] (many of same number and that is target)
        while array[right] == toMove:
            # element already in correct place
            right -= 1

        if array[left] == toMove:
            # element found swap positions with right
            array[left], array[right] = array[right], array[left]
        left += 1

    return array

print(move_to_end([3,4,5,6,4,4,5],4))
print(move_to_end([3,7,7,7,7,7,4,5,6,4,4,5,7,7,7,7,7,7,7,7,7],7))
print(move_to_end([7,7,7,7,7,5,7,7,7,7,7,7,7,7,7],5))

    


