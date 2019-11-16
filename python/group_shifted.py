'''
Given an array of strings (all lowercase letters), 
the task is to group them in such a way that all strings in a group are shifted versions of each other. 
Two string S and T are called shifted if,
    S.length = T.length 
    and
    S[i] = T[i] + K for 
    1 <= i <= S.length  for a constant integer K

Input  : ["acd", "dfg", "wyz", "yab", "mop",
                 "bdfh", "a", "x", "moqs"]

Output : a x
         acd dfg wyz yab mop
         bdfh moqs
All shifted strings are grouped together.
'''

def group_shifted_strs(input):

    shift_dict = {}

    for i in range(len(input)):
        # pull current string
        merged_strs = "".join(input[i])
        # check if in string of that length is in dict else create list for us to store vals
        merged_vals = shift_dict.get(len(merged_strs), [])
        # since the current value is new append it 
        merged_vals.append(input[i])
        # map it in the dict
        shift_dict[len(merged_strs)] = merged_vals
    return list(shift_dict.values())

print(group_shifted_strs(["acd", "dfg", "wyz", "yab", "mop",
                 "bdfh", "a", "x", "moqs"]))
print(group_shifted_strs(["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"]))