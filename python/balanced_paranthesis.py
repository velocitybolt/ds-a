'''

Given an expression string exp , write a program to examine whether the pairs and the orders of “{“,”}”,”(“,”)”,”[“,”]” are correct in exp.

Example:

Input: exp = “[()]{}{[()()]()}”
Output: Balanced

Input: exp = “[(])”
Output: Not Balanced

'''

def balanced(input):
    stack = []

    for curr in range(len(input)):
        if input[curr] == '{' or input[curr] == '[' or input[curr] == '(':
            stack.append(input[curr])
        elif input[curr] == ']' and (len(stack) == 0 or stack.pop() != '['):
            return False
        elif input[curr] == '}' and (len(stack) == 0 or stack.pop() != '{'):
            return False
        elif input[curr] == ')' and (len(stack) == 0 or stack.pop() != '('):
            return False 

    return len(stack) == 0

print(balanced('[()]{}{[()()]()}'))
print(balanced('[(])'))




        