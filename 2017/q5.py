import csv

reader = csv.reader(open('in5.txt'), delimiter='\n')
input = [int(x[0]) for x in reader]

def jump(jumps):
    num = 0
    cur = 0
    while cur < len(jumps) and cur >= 0:
        offset = jumps[cur]
        jumps[cur] += 1
        cur += offset
        num += 1
    return jumps, num

def jump2(jumps, step=lambda x: 1):
    num = 0
    cur = 0
    while cur < len(jumps) and cur >= 0:
        offset = jumps[cur]
        jumps[cur] += step(offset)
        cur += offset
        num += 1
    return jumps, num

def tests():
    t1 = [0,3,0,1,-3]
    j,n = jump2(t1)
    assert n==5

    t1 = [0,3,0,1,-3]
    stepfn = lambda x: 1 if x < 3 else -1
    j,n = jump2(t1, stepfn)
    assert (j,n)==([2,3,2,3,-1], 10)

tests()
stepfn = lambda x: 1 if x < 3 else -1
j,n = jump2(input, stepfn)
print n

'''
[(0),3,0,1,-3]
[(1),3,0,1,-3]
[2,(3),0,1,-3]
[2,2,0,1,(-3)]
[2,(2),0,1,-2]
[2,3,0,(1),-2]
[2,3,0,2,(-2)]
[2,3,(0),2,-1]
[2,3,1,(2),-1]
[2,3,1,2,-1]()
'''
