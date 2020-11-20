def createCoords(wire):
    coords = []
    curr = []
    posx = 0
    posy = 0
    num = 0
    for move in wire:
        num = move[1:]
        if move[0] == 'U':
            for x in range(int(num)):
                posy += 1
                curr = [posx, posy]
                coords.append(curr)
        if move[0] == 'D':
            for x in range(int(num)):
                posy -= 1
                curr = [posx, posy]
                coords.append(curr)
        if move[0] == 'L':
            for x in range(int(num)):
                posx -= 1
                curr = [posx, posy]
                coords.append(curr)
        if move[0] == 'R':
            for x in range(int(num)):
                posx += 1
                curr = [posx, posy]
                coords.append(curr)
    return coords

with open('challenge5input.txt', 'r') as fp:
    curr = 0
    wire1 = ''
    wire2 = ''
    for line in fp:
        if curr == 0:
            wire1 = line
            curr += 1
        else:
            wire2 = line
    wire1 = createCoords(wire1.split(','))
    wire2 = createCoords(wire2.split(','))
    aSteps = 0
    bSteps = 0
    for coordA in wire1:
        aSteps += 1
        for coordB in wire2:
            bSteps += 1
            if coordA == coordB:
                print(int(bSteps + aSteps))
            if bSteps == len(wire2):
                bSteps = 0
