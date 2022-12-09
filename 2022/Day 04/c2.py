f = open("input.txt", "r")
lines = f.readlines()
f.close()

total = 0
for ii in lines:
    elves = ii.split(',')
    e1 = elves[0].split('-')
    e2 = elves[1].split('-')
    l1 = int(e1[0])
    h1 = int(e1[1])
    l2 = int(e2[0])
    h2 = int(e2[1])
    if (l1 >= l2 and l1 <= h2) or (h1 <= h2 and h1 >= l2) or (l2 >= l1 and l2 <= h1) or (h2 <= h1 and h2 >= l1):
        total += 1

print(total)
