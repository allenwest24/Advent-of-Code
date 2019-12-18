import sys
import os
import math

print('got here')
filepath = sys.argv[1]

total = 0
with open(filepath) as fp:
    cnt = 0
    for line in fp:
        num = float(line)
        num = math.floor(num / 3)
        num = num - 2
        total =  total + num

print(total)
