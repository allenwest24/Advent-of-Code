def hashtag_in_vertical(lol, ii):   
    for jj in range(len(lol)):
        if lol[jj][ii] == "#":
            return True
    return False

def expand_universe(lol):
    # Check horizontal
    ii = 0
    while ii < len(lol):
        if not "#" in lol[ii]:
            add = ""
            for jj in range(len(lol[ii])):
                add += "."
            lol = lol[:ii] + [add] + lol[ii:]
            ii += 1
        ii += 1
    
    # Check vertical
    ii = 0
    while ii < len(lol[0].strip()):
        if not hashtag_in_vertical(lol, ii):
            for jj in range(len(lol)):
                lol[jj] = lol[jj][:ii] + "." + lol[jj][ii:]
            ii += 1
        ii += 1
    return lol

def map_stars(lol):
    starmap = []
    for ii in range(len(lol)):
        for jj in range(len(lol[ii])):
            if lol[ii][jj] == "#":
                starmap.append((ii, jj))
    return starmap

def count_all_shortest_paths_between_stars(starmap):
    out = 0
    for ii in range(len(starmap) - 1):
        for jj in range(ii + 1, len(starmap)):
            out += abs(starmap[ii][0] - starmap[jj][0]) + abs(starmap[ii][1] - starmap[jj][1])
    return out

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        ground = []
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        expanded = expand_universe(lol)
        starmap = map_stars(expanded)
        out = count_all_shortest_paths_between_stars(starmap)
        print(out)

if __name__ == "__main__":
    main() 