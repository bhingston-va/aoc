import csv

input = []
reader = csv.reader(open('in2.txt'), delimiter='\t')
for r in reader:
    input.append([int(x) for x in r])

def part1(input):
    sum = 0
    for r in input:
        sum += max(r) - min(r)
    return sum

def div(a,b):
    return max(a,b) / min(a,b)

def mod(a,b):
    return max(a,b) % min(a,b)

def wrongpart2(input):
    sum = 0
    for r in input:
        sum += reduce(lambda x,y: div(x,y) if x%y==0 else (div(x,y) if y%x==0 else 0), r)
    return sum

def part2(input):
    sum = 0
    for r in input:
        sum += [x/y for x in r for y in r if x%y==0 and x!=y][0]
    return sum

# part one test
t1 = [[5,1,9,5],[7,5,3],[2,4,6,8]]
assert part1(t1)==18
assert part1(input)==36174

#part two test
t2 = [[5,9,2,8],[9,4,7,3],[3,8,6,5]]
assert part2(t2)==9
assert part2(input)==244
