

# read input
input = []
with open('in1.txt') as f:
    for l in f:
        for c in l:
            if c == '\n':
                break
            else:
                input.append(int(c))


def part1(input):
    sum = 0
    for i in range(0, len(input)-1):
        if input[i] == input[i+1]:
            sum += input[i]

    if input[0] == input [-1:][0]:
        sum += input[0]

    return sum


def part2(input):
    sum = 0
    full = len(input)
    half = full/2
    for i in range(0, half):
        if input[i] == input[half+i]:
            sum += input[i]

    return sum*2

# part 1 test data
t1 = [1,1,2,2]
t2 = [1,1,1,1]
t3 = [1,2,3,4]
t4 = [9,1,2,1,2,1,2,9]
assert part1(t1)==3
assert part1(t2)==4
assert part1(t3)==0
assert part1(t4)==9
assert part1(input)==997


# part 2 test data
t1 = [1,2,1,2]
t2 = [1,2,2,1]
t3 = [1,2,3,4,2,5]
t4 = [1,2,3,1,2,3]
t5 = [1,2,1,3,1,4,1,5]
assert part2(t1)==6
assert part2(t2)==0
assert part2(t3)==4
assert part2(t4)==12
assert part2(t5)==4
assert part2(input)==1358
