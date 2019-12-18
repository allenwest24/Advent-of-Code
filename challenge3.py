import sys
import os
import math

filepath = sys.argv[1]

total = 0
with open(filepath) as fp:
    for line in fp:
        currentline = line.split(",")
        myarray = []
        ii = 0
        while ii < 137:
            myarray.append(currentline[ii])
            ii += 1
        print('array created')
    myarray[1] = 12
    myarray[2] = 2
    jj = 0
    temp1 = 0
    temp2 = 0
    temp3 = 0
    temp4 = 0
    while jj < 138:
        temp1 = int(myarray[jj])
        temp2 = int(myarray[jj + 1])
        temp3 = int(myarray[jj + 2])
        temp4 = int(myarray[jj + 3])
        if temp1 == 1:
            temp1 = int(myarray[temp2]) + int(myarray[temp3])
            myarray[temp4] = temp1
            jj += 4
        elif temp1 == 2:
            temp1 = int(myarray[temp2]) * int(myarray[temp3])
            myarray[temp4] = temp1
            jj += 4
        elif temp1 == 99:
            print(myarray)
            break
        else:
            print('error')
