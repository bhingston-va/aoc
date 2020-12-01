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


def find_entries(ls):
    a = ls
    b = ls
    for i in a:
        for j in b:
            if not sum_is_2020(i, j):
                break
            return i,j


def test_find_entries():
    x,y = find_entries(test_input)
    assert x in [299, 1721]
    assert y in [299, 1721]


test_find_entries()
