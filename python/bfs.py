import unittest

"""
Breadth First Search

Searching a parents nodes children before going deeper in the branch

"""

class Node:
    def __init__(self, name):
        self.children = []
        self.name = name

    def addChild(self, name):
        self.children.append(Node(name))
        return self

    def breadthFirstSearch(self, array):
        queue = [self] # holds root node
        while len(queue) > 0:
            # Pop the first element to keep order
            curr = queue.pop(0)
            # Add popped element to result array
            array.append(curr.name)
            for child in curr.children:
                queue.append(child)
        return array



test1 = Node("A")
test1.addChild("B").addChild("C")
test1.children[0].addChild("D")

test2 = Node("A")
test2.addChild("B").addChild("C").addChild("D").addChild("E")
test2.children[1].addChild("F")


class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(test1.breadthFirstSearch([]), ["A", "B", "C", "D"])

    def test_case_2(self):
        self.assertEqual(test2.breadthFirstSearch([]), ["A", "B", "C", "D", "E", "F"])


if __name__ == "__main__":
    unittest.main()