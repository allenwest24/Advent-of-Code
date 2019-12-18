# Uses input file 1
import sys
import os
import math

filepath = sys.argv[1]

total = 0
with open(filepath) as fp:
    for line in fp:
        num = float(line)
        num = math.floor(num / 3)
        num = num - 2
        total += num
        while num > -1:
            num = math.floor(num / 3)
            num -= 2
            if num <= 0:
                break
            else:
                total += num
print(total)
