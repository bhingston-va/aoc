import csv

# read input
input = []
reader = csv.reader(open('in4.txt'), delimiter=' ')
for r in reader:
    input.append([x for x in r])


def is_valid_ps(passphrase):
    while passphrase:
        word = passphrase.pop()
        if word in passphrase:
            return False
    return True

def sort_words(words):
    return [''.join(sorted(w)) for w in words]


def part1(input):
    num = 0
    for p in input:
        if is_valid_ps(p):
            num += 1
    return num

def part2(input):
    num = 0
    for p in input:
        sort_p = sort_words(p)
        if is_valid_ps(sort_p):
            num += 1
    return num


def test_is_valid():
    t1 = ['aa','bb','cc','dd','ee']
    t2 = ['aa','bb','cc','dd','aa']
    t3 = ['aa','bb','cc','dd','aaa']
    assert is_valid_ps(t1)==True
    assert is_valid_ps(t2)==False
    assert is_valid_ps(t3)==True
    t4 = ['abcde', 'fghij']
    t5 = ['abcde', 'xyz', 'ecdab']
    t6 = ['a', 'ab', 'abc', 'abd', 'abf', 'abj']
    t7 = ['iiii', 'oiii', 'ooii', 'oooi', 'oooo']
    t8 = ['oiii', 'ioii', 'iioi', 'iiio']
    assert is_valid_ps(sort_words(t4))==True
    assert is_valid_ps(sort_words(t5))==False
    assert is_valid_ps(sort_words(t6))==True
    assert is_valid_ps(sort_words(t7))==True
    assert is_valid_ps(sort_words(t8))==False

test_is_valid()

#print part1(input)
print part2(input)
