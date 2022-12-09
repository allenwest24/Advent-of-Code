f = open("input.txt", "r")
lines = f.readlines()
f.close()

board = [[],[],[],[],[],[],[],[],[],[]]
out = ""

ii = 0
while 'move' not in lines[ii]:
    for jj in range(len(lines[ii])):
        if lines[ii][jj].isalpha():
            board[(jj // 4) + 1].insert(0, lines[ii][jj])
    ii += 1

print(board)

for kk in range(ii, len(lines)):
    command = lines[kk].split(' ')
    num = int(command[1])
    frm = int(command[3])
    to = int(command[5])
    for _x in range(num):
        board[to].append(board[frm].pop())

for column in range(1,len(board)):
    out += board[column][-1]

print(out)
