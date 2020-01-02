import unittest

"""
Depth First Search

Searching through nodes by exploring the depth of a node first

"""

class Node:
    def __init__(self, name):
        self.children = []
        self.name = name

    def addChild(self, name):
        self.children.append(Node(name))
        return self

    def depthFirstSearch(self, array):
        # append the node to be searched
        array.append(self.name)
        for child in self.children:
            child.depthFirstSearch(array)
        return array


test1 = Node("A")
test1.addChild("B").addChild("C")
test1.children[0].addChild("D")


test5 = Node("A")
test5.addChild("B").addChild("C").addChild("D").addChild("L").addChild("M")
test5.children[0].addChild("E").addChild("F").addChild("O")
test5.children[1].addChild("P")
test5.children[2].addChild("G").addChild("H")
test5.children[0].children[0].addChild("Q").addChild("R")
test5.children[0].children[1].addChild("I").addChild("J")
test5.children[2].children[0].addChild("K")
test5.children[4].addChild("S").addChild("T").addChild("U").addChild("V")
test5.children[4].children[0].addChild("W").addChild("X")
test5.children[4].children[0].children[1].addChild("Y").addChild("Z")

class TestProgram(unittest.TestCase):
    def test_case_1(self):
        self.assertEqual(test1.depthFirstSearch([]), ["A", "B", "D", "C"])

    def test_case_5(self):
        self.assertEqual(
        test5.depthFirstSearch([]),
        [
                "A",
                "B",
                "E",
                "Q",
                "R",
                "F",
                "I",
                "J",
                "O",
                "C",
                "P",
                "D",
                "G",
                "K",
                "H",
                "L",
                "M",
                "S",
                "W",
                "X",
                "Y",
                "Z",
                "T",
                "U",
                "V",
        ],
    )



if __name__ == "__main__":
    unittest.main()
