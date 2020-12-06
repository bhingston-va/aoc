import csv

'''
--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation at
a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold
coins used there have a little picture of a starfish; the locals just call them
stars. None of the currency exchanges seem to have heard of them, but somehow,
you'll need to find fifty of these coins by the time you arrive so you can pay
the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day
in the Advent calendar; the second puzzle is unlocked when you complete the
first. Each puzzle grants one star. Good luck!

--- Part One ---

Before you leave, the Elves in accounting just need you to fix your expense
report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then
multiply those two numbers together.

--- Part Two ---

The Elves in accounting are thankful for your help; one of them even offers you
a starfish coin they had left over from a past vacation. They offer you a second
one if you can find three numbers in your expense report that meet the same
criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366,
and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to
2020?
'''

test_input = [
    1721, 979, 366, 299, 675, 1456
]


def sum_is_s(to_sum, s):
    return sum(to_sum) == s


def _find_product_of_2(ls, summ):
    first = ls[0]
    rest = ls[1:]
    products = [first * x for x in rest if sum_is_s([first, x], summ)]
    (p, found) = (products[0], True) if len(products) > 0 else (None, False)
    return p, found


def find_product_of_2(ls, summ):
    p, found = _find_product_of_2(ls, summ)
    if found:
        return p, found

    l = ls
    while not found:
        if len(l) <= 2:
            break
        l = l[1:]
        p, found = _find_product_of_2(l, summ)

    return p, found


def find_product_of_3(ls, summ):
    first = ls[0]
    rest = ls[1:]
    new_sum = summ - first
    p, found = find_product_of_2(rest, new_sum)
    if found:
        return p * first, True

    l = ls
    while not found:
        l = l[1:]
        if len(l) < 3:
            break
        p, found = find_product_of_3(l, summ)

    return p, found


def part1(ls):
    p, _ = find_product_of_2(ls, 2020)
    return p


def part2(ls):
    p, _ = find_product_of_3(ls, 2020)
    return p


if __name__ == '__main__':
    input = []
    reader = csv.reader(open('input-a.txt'), delimiter='\t')
    for r in reader:
        input.append(int(r[0]))

    # part one test
    assert part1(test_input) == 514579
    a = part1(input)
    print('part 1 answer:')
    print(a)
    assert a == 440979

    # part two test
    assert part2(test_input) == 241861950

