import sys
import os
import math

filepath = sys.argv[1]

total = 0
besttotal = 0
with open(filepath) as fp:
    line = fp.readline()
    currentline = line.split(",")
    myarray = []
    ii = 0
    while ii < 137:
        myarray.append(currentline[ii])
        ii += 1
    print('array created')
    jj = 0
    vert = 0
    horz = 0
    coordTracker1 = []
    while jj < 137:
        temp = myarray[jj]
        way = temp[0]
        if way == 'U':
            vert += int(temp[1:])
        if way == 'D':
            vert -= int(temp[1:])
        if way == 'L':
            horz -= int(temp[1:])
        if way == 'R':
            horz += int(temp[1:])
        coordTracker1.append(horz)
        coordTracker1.append(vert)
        jj += 1
    #print(coordTracker1)

    line = fp.readline()
    currentline = line.split(",")
    myarray = []
    ii = 0
    while ii < 137:
        myarray.append(currentline[ii])
        ii += 1
    print('array created')
    jj = 0
    vert = 0
    horz = 0
    coordTracker2 = []
    while jj < 137:
        temp = myarray[jj]
        way = temp[0]
        if way == 'U':
            vert += int(temp[1:])
        if way == 'D':
            vert -= int(temp[1:])
        if way == 'L':
            horz -= int(temp[1:])
        if way == 'R':
            horz += int(temp[1:])
        coordTracker2.append(horz)
        coordTracker2.append(vert)
        jj += 1
    #print(coordTracker2)

    for ii in range(68):
        for jj in range(68):
            if coordTracker1[2 * ii] == coordTracker2[2 * jj] and coordTracker1[2 * ii + 1] == coordTracker2[2 * jj + 1]:
                intsersect = abs(coordTracker1[2 * ii]) + abs(coordTracker1[2 * ii + 1])
                print(intersect)
                jj += 1
            ii += 1
       

