import csv
from math import ceil

reader = csv.reader(open('in6.txt'), delimiter='\t')
input = [int(x) for r in reader for x in r]

def memShareSize(a,b):
    return int(ceil(float(a)/b))

def shuffle(input):
    size = len(input)
    valu = max(input)
    indx = input.index(valu)

    input[indx] = 0
    i = indx
    memdiv = memShareSize(size, valu)
    while valu > 0:
        i += 1
        i = 0 if i == size else i
        valu - memdiv
        input[i] += memdiv


    return [2,4,1,2,], 5

def tests():
    t1 = [0,2,7,0]
    assert shuffle(t1)==([2,4,1,2], 5)

tests()

