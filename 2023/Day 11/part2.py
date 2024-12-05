

def hashtag_in_vertical(lol, ii):   
    for jj in range(len(lol)):
        if lol[jj][ii] == "#":
            return True
    return False

def expand_universe(lol):
    double_is = []
    double_js = []

    # Check horizontal
    ii = 0
    while ii < len(lol):
        if not "#" in lol[ii]:
            double_is.append(ii)
        ii += 1
    
    # Check vertical
    ii = 0
    while ii < len(lol[0].strip()):
        if not hashtag_in_vertical(lol, ii):
            double_js.append(ii)
        ii += 1
    return lol, double_is, double_js

def map_stars(lol):
    starmap = []
    for ii in range(len(lol)):
        for jj in range(len(lol[ii])):
            if lol[ii][jj] == "#":
                starmap.append((ii, jj))
    return starmap

def count_all_shortest_paths_between_stars(starmap, double_is, double_js):
    out = 0
    mult = 999999
    for ii in range(len(starmap) - 1):
        for jj in range(ii + 1, len(starmap)):
            out += abs(starmap[ii][0] - starmap[jj][0]) + abs(starmap[ii][1] - starmap[jj][1])
            mini = min(starmap[ii][0], starmap[jj][0])
            maxi = max(starmap[ii][0], starmap[jj][0])
            for x in double_is:
                if mini < x < maxi:
                    out += mult
            minj = min(starmap[ii][1], starmap[jj][1])
            maxj = max(starmap[ii][1], starmap[jj][1])
            for y in double_js:
                if minj < y < maxj:
                    out += mult
    return out

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        ground = []
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        expanded, double_is, double_js = expand_universe(lol)
        starmap = map_stars(expanded)
        out = count_all_shortest_paths_between_stars(starmap, double_is, double_js)
        print(out)

if __name__ == "__main__":
    main() 
