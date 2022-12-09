f = open("input.txt", "r")
lines = f.readlines()
f.close()

total = 0
grid = []
for row in lines:
    grid.append(row.strip())

for ii in range(len(grid)):
    for jj in range(len(grid[ii])):
        if ii == 0 or ii == len(grid) - 1 or jj == 0 or jj == len(grid[jj]) - 1:
            total += 1
            continue
        else:
            good = True
            curr = int(grid[ii][jj])
            up = True
            for kk in range(0, ii):
                if int(grid[kk][jj]) >= curr:
                    up = False
            down = True
            for kk in range(ii + 1, len(grid)):
                if int(grid[kk][jj]) >= curr:
                    down = False
            left = True
            for kk in range(0, jj):
                if int(grid[ii][kk]) >= curr:
                    left = False
            right = True
            for kk in range(jj + 1, len(grid[ii])):
                if int(grid[ii][kk]) >= curr:
                    right = False
            if up or down or left or right:
                total += 1

print(total)
