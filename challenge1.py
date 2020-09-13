import sys
import os
import math

filepath = 'C:\\Users\\Allen\\Desktop\\input1.txt'

total = 0
with open(filepath) as fp:
    cnt = 0
    for line in fp:
        total += math.floor(int(line)/3) - 2

print(total)
