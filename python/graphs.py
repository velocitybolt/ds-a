class Node:

    def __init__(self, val, neighbors):
        self.val = val
        self.neighbors = neighbors

# Keep track of visited nodes
# dfs - > Depth First Search
# apply recursion

def dfs_helper(node, visited):
    if node == None:
        return None

    elif node in visited.keys():
        return visited[node]
    
    neighbors = []
    new_node = Node(node.val, neighbors)

    visited[node] = new_node

    for i in range(len(node.neighbors)):
        neighbor_node = dfs_helper(node.neighbors[i], visited)
        neighbors.append(neighbor_node)

    return new_node

def dfs(node):
    return dfs_helper(node, dict())


def bfs(node):
    queue = []
    visited = {}

    queue.append(node)

    while len(queue) > 0:
    




node = Node(1, [])
node2 = Node(2, [])
node3 = Node(3, [])
node4 = Node(4, [])

node.neighbors.append(node2)
node.neighbors.append(node4)

node2.neighbors.append(node)
node2.neighbors.append(node3)

node3.neighbors.append(node2)
node3.neighbors.append(node4)

node4.neighbors.append(node)
node4.neighbors.append(node3)

