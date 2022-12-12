f = open("input.txt", "r")
lines = f.readlines()
f.close()

total = 0
grid = []
for row in lines:
    grid.append(row.strip())

best = 0
for ii in range(len(grid)):
    for jj in range(len(grid[ii])):
        curr = int(grid[ii][jj])
        if ii == 0:
            up = 0
        else:
            up = 0
            for kk in range(1, ii + 1):
                if int(grid[ii - kk][jj]) < curr:
                    up += 1
                else:
                    up += 1
                    break
        if ii == len(grid) - 1:
            down = 0
        else:
            down = 0
            for kk in range(ii + 1, len(grid)):
                if int(grid[kk][jj]) < curr:
                    down += 1
                else:
                    down += 1
                    break
        if jj == 0:
            left = 0
        else:
            left = 0
            for kk in range(1, jj + 1):
                if int(grid[ii][jj - kk]) < curr:
                    left += 1
                else:
                    left += 1
                    break
        if jj == len(grid[jj]) - 1:
            right = 0
        else:
            right = 0
            for kk in range(jj + 1, len(grid[ii])):
                if int(grid[ii][kk]) < curr:
                    right += 1
                else:
                    right += 1
                    break
        total = up * down * left * right
        if total > best:
            best = total

print(best, ii, jj)
