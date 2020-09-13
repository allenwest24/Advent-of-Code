import sys
import os
import math

filepath = 'C:\\Users\\Allen\\Desktop\\input1.txt'

total = 0
with open(filepath) as fp:
    cnt = 0
    for line in fp:
        subtotal = (math.floor(int(line)/3) - 2)
        total += subtotal
        while subtotal > 0:
            subtotal = (math.floor(subtotal/3) - 2)
            if subtotal > 0:
                total += subtotal
            else:
                break
print(total)
