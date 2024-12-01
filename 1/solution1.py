import sys

lines = [line.strip().split() for line in sys.stdin]

a, b = zip(*[(int(x), int(y)) for x, y in lines])

print(sum(abs(x - y) for x, y in zip(sorted(a), sorted(b))))



