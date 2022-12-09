f = open("input.txt", "r")
lines = f.readlines()
f.close()

dir_contents = {}
sizes = {}

# When given a file, returns size, when given a directory, recursively adds up to total size.
def get_size(parent, list_of_paths):
    total = 0
    for ii in list_of_paths:
        parts = ii.split(" ")
        if parts[0] != 'dir':
            total += int(parts[0])
        else:
            dir_size = get_size(parent + parts[1] + '/', dir_contents[parent + parts[1] + '/'])
            total += dir_size
    sizes[parent] = total
    return total


curr = "/"
for ii in range(len(lines)):
    if "$ cd" in lines[ii]:
        if ".." in lines[ii]:
            curr = '/'.join(curr.split('/')[:-2]) + '/'
            print("back to " + curr)
        else:
            if lines[ii].split(" ")[2].strip() != curr:
                curr += lines[ii].split(" ")[2].strip()+ '/'
                print(curr)
    elif "$ ls" in lines[ii]:
        dir_contents[curr] = []
        jj = ii + 1
        while (jj < len(lines)) and (not '$' in lines[jj]):
            dir_contents[curr].append(lines[jj].strip())
            jj += 1

for k, v in dir_contents.items():
    if not k in sizes:
        _g = get_size(k, v)

total_disk = 70000000
root_size = sizes['/']
avail = total_disk - root_size
needed = 30000000 - avail
best = root_size
for k, v in sizes.items():
    curr = v - needed
    if curr >= 0 and curr < best - needed:
        best = v

print(best)
