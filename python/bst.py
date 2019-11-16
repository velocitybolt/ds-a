# Tree problems

class Node:
    def __init__(self, val, left = None, right = None):
        self.val = val
        self.left = left
        self.right = right

# O(n) o(n2) worse case
def inorder(root):
    # Traversal of all the left nodes first then the ones on the right
    if root == None:
        return 
    inorder(root.left)
    print(root.val)
    inorder(root.right)

# O(lgn) O(n) worse case
def search(root, target):
    if root == None:
        return None
    # the search target can equal target
    elif target == root.val:
        return root
    # the search value can be greater than root which means on the rightside
    elif target > root.val:
        return search(root.right, target)
    # then its smaller because we know that all bsts have smaller nodes on left
    else:
        return search(root.left, target)

def same_tree(root1, root2):
    if root1 == None and root2 == None:
        return True
    elif root1 == None or root2 == None:
        return False
    elif root1.val != root2.val:
        return False
    
    return same_tree(root1.left, root2.left) and same_tree(root2.right, root2.left)

def sym_tree(root):
    if root == None:
        return True
    return same_tree(root.left, root.right)


def is_valid_help(root, low, high):
    if root == None:
        return True
    elif high != None and root.val >= high or low != None and root.val <= low:
        return False
    
    return is_valid_help(root.left, low, root.val) and is_valid_help(root.right, root.val, high)

def is_valid_bst(root):
   return is_valid_help(root, None, None)


def convert_bst_to_ll(root):
    if root == None:
        return None
    # check if leaf node
    elif root.left == None and root.right == None:
        # returns head and tail of list
        return [root, root]

    # store elements from list (array of head and tail from left)
    left = convert_bst_to_ll(root.left)
    # store elements from right (array of head and tail from right)
    right = convert_bst_to_ll(root.right)

    # begin to build ll --> head_tail[0] = head ---> head_tail[1] = tail 
    head_tail = [None, None]
    
    # make sure left tail is connected to root. 
    if left != None:
        # left[1] represents tail from left side
        left[1].right = root
        root.left = left[1]
        
        # append head from the left side to the head of this list
        head_tail[0] = left[0]
    else:
        # nothing left then set root to head of list
        head_tail[0] = root

    if right != None:
        # take head from right 
        right[0].left = root
        root.right = right[0]

        # append tail from right side to the tail of this list
        head_tail[1] = right[1]
    else:
        # otherwise it probably doesnt exist so set it to root
        head_tail[1] = root

    return head_tail


leaf1 = Node(4)
leaf2 = Node(7)
leaf3 = Node(11)
leaf4 = Node(14)
leaf5 = Node(37)
leaf6 = Node(17)

node1 = Node(5, leaf1, leaf2)
node2 = Node(12, leaf3, leaf4)
node3 = Node(39, leaf5, None)
node4 = Node(32, leaf6, node3)
node5 = Node(1, None, node1)
node6 = Node(10, node5, node2)

root = Node(15, node6, node4)

'''
Function Testing
'''
output = convert_bst_to_ll(root)

head = output[0]

while head:
    print(head.val)
    head = head.right

print("----------------------")

tail = output[1]

while tail:
    print(tail.val)
    tail = tail.left
'''
print(search(root,37).val)
print(search(root,4444))
print(search(root, 15).val)
inorder(root)
'''



