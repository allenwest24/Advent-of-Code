f = open("input.txt", "r")
data = f.read()
f.close()

track = '    '
for c in range(len(data)):
    track = track[1:] + data[c]
    if len(set(track)) == 4 and len(track.strip()) == 4:
        print(c + 1)
        break
