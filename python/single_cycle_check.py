"""
Single Cycle Check

Given an array of integers each integer reporesnts a jump of its value in the array

2 means two forward. -3 means three backward. -1 at index 0 jumps to the last index in the array (wrapping) 

A cycle occurs when starting at any index in the array doing the given jumps every element is visited once

before landing back on the starting index. Return True if One cycle False if > 1.

SI : [2, 3, 1, -4, -4, 2]
SO: True
"""


def single_cycle(array):
    startingIndex = 0 
    elementsVisited = 0 
    while elementsVisited < len(array):
        # this means more than one cycle because we are not at the start but we are at starting index
        if elementsVisited > 0 and startingIndex == 0:
            return False
        
        elementsVisited += 1
        startingIndex = getNextIndex(array, startingIndex)

    return startingIndex == 0


def getNextIndex(array, startingIndex):
    jump = array[startingIndex]
    nextIndex = (jump + startingIndex) % len(array)
    return nextIndex if nextIndex > 0 else nextIndex + len(array)







 