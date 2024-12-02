from sys import stdin
from itertools import pairwise


lines = [map(int, line.strip().split()) for line in stdin]

count = len(lines)
for line in lines:
    sign = 0
    for a, b in pairwise(line):
        diff = a - b
        if abs(diff) < 1:
            count -= 1
            break
        if abs(diff) > 3:
            count -= 1
            break
        if diff < 0 and sign > 0 or diff > 0 and sign < 0:
            count -= 1
            break
        sign = diff
print(count)

