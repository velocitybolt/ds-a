
'''
Code for testing and learning purposes only
'''
import numpy as np 

array_two = np.arange(1,4) ** 2
array_three = np.arange(1,4) ** 3
exponent_array = np.array([1, 4, 5])

print(array_two)
print(array_three)

# summation of arrays rather than just appending w regular arrayus
print(array_three + array_two)

# power func
print(np.power(np.array([1,2,3]), 4))

# negative
print(np.negative(np.array([1,2,3])))

#exponent etc.......... check documentation
print(np.exp(exponent_array))

#multi-dimensional array
x = np.arange(3)
y = np.arange(3, 6)
z = np.arange(6, 9)

multi_arr = np.array([x,y,z])
print(multi_arr)

#can specify shape
print(multi_arr.shape)

#generate numbers
w = np.linspace(1,10,100)
print(w)
#every thrid number 
b = np.arange(1,30,3)
print(b)

print(multi_arr)
print(multi_arr[0])
print(multi_arr[2,2])

comp = multi_arr >= 5
print(comp)






