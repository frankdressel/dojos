#! /usr/bin/env python

import random

with open('testfile.txt', 'r') as f:
    data = f.readlines()[0:3]

for i in range(1000):
    randomstart = random.sample(population=[0, 3, 6, 9, 12, 15, 18, 21, 24], k=9)
    newlines = []
    for i in range(3):
        newlines.append(''.join([data[i][r:r+3] for r in randomstart]))
    print('{}\n{}\n{}'.format(*newlines))
    print()
