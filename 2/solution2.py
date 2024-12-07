from sys import stdin
from itertools import pairwise


def isSafe(line):
    sign = 0
    for a, b in pairwise(line):
        diff = a - b
        if abs(diff) < 1:
            return False
        if abs(diff) > 3:
            return False
        if diff < 0 and sign > 0 or diff > 0 and sign < 0:
            return False
        sign = diff
    return True


lines = [list(map(int, line.strip().split())) for line in stdin]

count = 0
for line in lines:
    sign = 0
    if isSafe(line):
        count += 1
    else:
        for i in range(len(line)):
            if isSafe(line[:i] + line[i+1:]):
                count += 1
                break

print(count)
