f = open("input.txt", "r")
lines = f.readlines()
f.close()

alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

total = 0
ii = 0
while ii < len(lines) - 2:
    common = []
    elf1 = lines[ii]
    elf2 = lines[ii + 1]
    elf3 = lines[ii + 2]
    for item in elf1:
        if item in elf2 and item in elf3:
            total += alpha.index(item) + 1
            break
    ii += 3

print(total)
