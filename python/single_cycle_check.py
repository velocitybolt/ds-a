"""
Single Cycle Check

Given an array of integers each integer reporesnts a jump of its value in the array

2 means two forward. -3 means three backward. -1 at index 0 jumps to the last index in the array (wrapping) 

A cycle occurs when starting at any index in the array doing the given jumps every element is visited once

before landing back on the starting index

SI : [2, 3, 1, -4, -4, 2]
SO: True
"""


def single_cycle(array):
    # The starting index has to be recorded to see if there has been more than one cycle
    startingIndex = 0
    # Elements visited needs to be recorded to keep bounds check and another way to see if there 
    # is more than one cycle
    elementsVisited = 0
    while elementsVisited < len(array):
        if elementsVisited > 0 and startingIndex == 0:
            # double cycle because we are not at the end of array yet
            # being at index 0 at this point in the loop means we visited
            # the start more than once
            return false
        # get next starting index
        startingIndex = getNextIndex(array, startingIndex)
        # exit out of the loop after one run if we end up at the start index (0 in this case)
    return startingIndex == 0


def getNextIndex(array, startingIndex):
    # The purpose of this method is get the next index depending on the jump
    # All at the sametime watching for edgecases to stay inbound
    jump = array[startingIndex]
    nextIndex = (startingIndex + jump) % len(array)
    return nextIndex if startingIndex >= 0 else nextIndex + len(array)







 