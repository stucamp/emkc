import sys


nums = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}


def parse_roman_nums(array):
    tot = 0

    if len(array) == 1:
        tot += nums[array[0]]
    else:
        i = len(array)-1
        while i >= 0:
            if i != 0:
                if nums[array[i]] > nums[array[i-1]]:
                    tot += nums[array[i]]
                    tot -= nums[array[i-1]]
                    i -= 1
                else:
                    tot += nums[array[i]]
            elif i == 0:
                if nums[array[0]] < nums[array[1]]:
                    tot -= nums[array[0]]
                else:
                    tot += nums[array[0]]
            i -= 1
    return tot


def split_string(word):
    return [char for char in word]


value1 = sys.argv[1]
arr = split_string(value1)

print(parse_roman_nums(arr))


