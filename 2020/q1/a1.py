import csv

'''
Before you leave, the Elves in accounting just need you to fix your expense
report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020
and then multiply those two numbers together.

For example, suppose your expense report contained the following:
1721, 979, 366, 299, 675, 1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying
them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to
2020; what do you get if you multiply them together?
'''

test_input = [
    1721, 979, 366, 299, 675, 1456
]

def sum_is_2020(a, b):
    return a + b == 2020


def find_entries_product(ls):
    a = ls
    b = ls
    for i in a:
        for j in b:
            if not sum_is_2020(i, j):
                break
            return i*j
    print('nothing found')

def find_product(ls):
    first = ls[0]
    rest = ls[1:]
    products = [first * x for x in rest if sum_is_2020(first, x)]
    (p, found) = (products[0], True) if len(products) > 0 else (None, False)
    return p, found


def part1(i):
    return _defer_part1(i)

def _defer_part1(ls):
    p, found = find_product(ls)
    if found:
        return p

    l = ls
    done = found
    while not done:
        l = l[1:]
        p, done = find_product(l)
        if len(l) < 2:
            done = True

    return p


input = []
reader = csv.reader(open('input-a.txt'), delimiter='\t')
for r in reader:
    input.append(int(r[0]))

# part one test
assert part1(test_input) == 514579
a = part1(input)
print('part 1 answer:')
print(a)
