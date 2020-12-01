from math import ceil, sqrt

# spiral memory question
input = 277678

def layer(cell):
    for l in xrange(cell):
        mem = mem_range(l)
        if mem[0] <= cell and cell <= mem[1]:
            return l

def mem_range(layer):
    '''
    k = 2l + 1, l = [0,1,2,...]
    l = floor of k/2
    m = [ (k-2)^2 + 1, k^2 ]
    '''
    k = 2*layer + 1
    return ((k-2)**2+1, k**2) if layer > 0 else (1,1)

def dist(cell):
    '''
    d = (g mod h) + l
    g = m - (l-1)
    h = l + 1
    '''
    l = layer(cell)
    if l == 0:
        return 0
    return ((cell+1-l) % (l+1)) + l

def new_dist(cell):
    '''
    gen list starting at [2l-1, decr to min=l, incr to max=2l]
    l = layer; gen list only for memory range for layer, l
    '''
    if cell == 1:
        return 0

    l = layer(cell)
    n,m = mem_range(l)
    cells = xrange(n,m+1)
    min_d, max_d, start = l, 2*l, (2*l-1)
    return build_cell_dist(start, min_d, max_d, cells, cell)


def build_cell_dist(start, minn, maxx, cells, target):
    ls = []
    decr = True
    d = start
    for c in cells:
        ls += (c,d)

        # return just distance for now (could return whole list @ end)
        if c == target:
            return d

        if decr and d == minn:
            decr = False
        elif not decr and d == maxx:
            decr = True

        d = d-1 if decr else d+1


def test_layer():
    assert layer(1)==0
    assert layer(2)==1
    assert layer(7)==1
    assert layer(9)==1
    assert layer(10)==2
    assert layer(25)==2

def test_mem_range():
    assert mem_range(0)==(1,1)
    assert mem_range(1)==(2,9)
    assert mem_range(2)==(10,25)
    assert mem_range(3)==(26,49)

def test_dist():
    assert new_dist(1)==0
    assert new_dist(2)==1
    assert new_dist(3)==2
    assert new_dist(5)==2
    assert new_dist(16)==3
    assert new_dist(15)==2
    assert new_dist(13)==4

def test_all():
    test_layer()
    test_mem_range()
    test_dist()


test_all()

#part one
print 'layer:' + str(layer(input))
print 'distance to {}:{}'.format(input,new_dist(input))
"""
forgot to call new_dist instead of dist. manually got answer by:
For input the layer, l = 263.
Last number M = k^2, where k = 2l+1
Max distance (to M), D = 2*263.
Now M - input = 51, so D-51=475
"""


