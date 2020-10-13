total = 0
tempdouble = 0
tempnotdecreasing = 1
prev = 0
winners = []
case1 = 0
case2 = 0
for x in range(256310, 732736):
    y = 0
    for y in range(6):
        if y == 0:
            prev = int(str(x)[y])
        else:
            if prev == int(str(x)[y]) and not tempdouble:
                if y < 5:
                    if prev == int(str(x)[y+1]):
                        case1 = 1
                if y > 1:
                    if prev == int(str(x)[y-2]):
                        case2 = 1
                if not case1 and not case2:
                    tempdouble = 1
                case1 = 0
                case2 = 0
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
