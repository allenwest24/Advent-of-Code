f = open("input.txt", "r")
lines = f.readlines()
f.close()

alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

total = 0
for ii in lines:
    h1 = ii[:(len(ii) - 1) // 2]
    h2 = ii[(len(ii) - 1) // 2:-1]
    print(len(h1), len(h2))
    for jj in h2:
        if jj in h1:
            total += alpha.index(jj) + 1
            break

print(total)
