import sys

value1 = sys.argv[1]

num_dict = {
    '0': 0,
    '1': 0,
    '2': 0,
    '3': 0,
    '4': 0,
    '5': 0,
    '6': 0,
    '7': 0,
    '8': 0,
    '9': 0,
}

highest = '0'

# write your solution here
for i in value1:
    num_dict[i] += 1
    if num_dict[i] >= num_dict[highest]:
        highest = i

print(f"{highest}\n")