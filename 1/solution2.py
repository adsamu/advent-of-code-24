import sys
from collections import Counter

lines = [line.strip().split() for line in sys.stdin]

a, b = zip(*[(int(x), int(y)) for x, y in lines])

c = Counter(b)

print(sum([x*c[x] for x in a]))

