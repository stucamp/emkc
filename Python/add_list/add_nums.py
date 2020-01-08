import sys

value1 = sys.argv[1]

numlist = [int(x) for x in value1.split(',')]

numlist.sort(reverse=True)
ind = 0
for i in numlist:
    if ind != len(numlist)-1:
        print(f"{i}", end=',')
    else:
        print(f"{i}", end='')
    ind +=1
