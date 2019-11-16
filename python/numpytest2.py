'''
This 
file
contains
more
testing
of
numpy
functions
'''

import numpy as np 

# declare 3x3 array with numbers from 0 - 8
x = np.arange(9).reshape(3,3)
print(x)

# flatten array to 1D w ravel (returns view changes both new and old array)
ravelled_array = x.ravel()
print(ravelled_array)

# flatten function makes new array (copy)
flattend_array = x.flatten()
print(flattend_array)

#another way to resize
y = np.arange(9)
y.shape = [3,3]
print(y.shape[0])
print(y.T)
print(np.resize(y, [6,6]))

#array with hella zeros
dw = np.zeros([6], dtype = int)
print(dw)

print(np.eye(3))
print(np.random.rand(8,8))

#matrices
mat_a = np.matrix([0,3,5,5,5,2]).reshape(2,3)
mat_b = np.matrix([3,4,3,-2,4,-2]).reshape(3,2)
print(mat_a * mat_b)
print(np.matmul(mat_a,mat_b))
print(mat_a @ mat_b)
print(y.shape[0])

#hstack horizontally combines matrices
xy = np.arange(4).reshape(2,2)
yx = np.arange(4, 8).reshape(2,2)
wxy = np.hstack((xy,yx))
print(xy)
print(yx)
print(wxy)
print("----part 2-------")
#depth stack
d_stack = np.dstack((xy,yx))
print(d_stack)
print(d_stack.shape)
