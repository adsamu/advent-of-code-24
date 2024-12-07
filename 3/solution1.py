from sys import stdin
import re
from functools import reduce

string = ""
for line in stdin:
    string += line

string = re.findall(r"(?:^|do\(\)).*?(?:don't\(\)|$)", string, re.DOTALL)
string = "".join(string)



muls = re.findall(r"mul\(\d{1,3},\d{1,3}\)" , string)

nums = list(map(lambda x: x[4:-1].split(','), muls))
print(reduce(lambda prev, x,: prev + (int(x[0]) * int(x[1])), nums, 0))

