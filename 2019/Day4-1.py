total = 0
tempdouble = 0
tempnotdecreasing = 1
prev = 0
winners = []
for x in range(256310, 732736):
    y = 0
    for y in range(6):
        if y == 0:
            prev = int(str(x)[y])
        else:
            if prev == int(str(x)[y]):
                tempdouble = 1
            if prev > int(str(x)[y]):
                tempnotdecreasing = 0
            prev = int(str(x)[y])
    if tempdouble and tempnotdecreasing:
        total += 1
        winners.append(x)
    tempdouble = 0
    tempnotdecreasing = 1
    prev = 0
print(total)
