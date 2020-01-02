
a = [1,2,3,3,4,4,5,5,5,56,6,6]

y = a[::2]

x = a[::-1]

'''
print(y)
print(x)
print(list(enumerate(x)))

for i, num in enumerate(a):
    print('%d: %s' % (i + 1, num))

'''
president_names = ['Trump', 'Clinton', 'Obama', 'Washington', 'Bush']
letters = [len(n) for n in president_names]
print(f"The length of the list in terms of characters is {sum(letters)}")

for name in president_names:
    print(f"President {name} is of length {len(name)}")

max_letters = 0
longest_name = None

for name in range(len(president_names)):
    count = letters[name]
    if count > max_letters:
        longest_name = president_names[name]
        max_letters = count

print(f"The longest name in the list is {longest_name}")

# zip merges two iterators together..
print(list(zip(president_names, letters)))



