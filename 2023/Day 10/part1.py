def get_all_valid_positions(ground, curr):
    valid = []
    if ground[curr[0]][curr[1]] == "|":
        # Down
        if curr[0] + 1 < len(ground) and (ground[curr[0] + 1][curr[1]] == "L" or ground[curr[0] + 1][curr[1]] == "|" or ground[curr[0] + 1][curr[1]] == "J"):
            valid.append((curr[0] + 1, curr[1]))
        # Up
        if curr[0] - 1 >= 0 and (ground[curr[0] - 1][curr[1]] == "7" or ground[curr[0] - 1][curr[1]] == "|" or ground[curr[0] - 1][curr[1]] == "F"):
            valid.append((curr[0] - 1, curr[1]))
    elif ground[curr[0]][curr[1]] == "-":
        # Left
        if curr[1] - 1 >= 0 and (ground[curr[0]][curr[1] - 1] == "F" or ground[curr[0]][curr[1] - 1] == "-" or ground[curr[0]][curr[1] - 1] == "L"):
            valid.append((curr[0], curr[1] - 1))
        # Right
        if curr[1] + 1 < len(ground[0]) and (ground[curr[0]][curr[1] + 1] == "J" or ground[curr[0]][curr[1] + 1] == "-" or ground[curr[0]][curr[1] + 1] == "7"):
            valid.append((curr[0], curr[1] + 1))
    elif ground[curr[0]][curr[1]] == "7":
        # Left
        if curr[1] - 1 >= 0 and (ground[curr[0]][curr[1] - 1] == "F" or ground[curr[0]][curr[1] - 1] == "-" or ground[curr[0]][curr[1] - 1] == "L"):
            valid.append((curr[0], curr[1] - 1))
        # Down
        if curr[0] + 1 < len(ground) and (ground[curr[0] + 1][curr[1]] == "L" or ground[curr[0] + 1][curr[1]] == "|" or ground[curr[0] + 1][curr[1]] == "J"):
            valid.append((curr[0] + 1, curr[1]))
    elif ground[curr[0]][curr[1]] == "J":
        # Left 
        if curr[1] - 1 >= 0 and (ground[curr[0]][curr[1] - 1] == "F" or ground[curr[0]][curr[1] - 1] == "-" or ground[curr[0]][curr[1] - 1] == "L"):
            valid.append((curr[0], curr[1] - 1))
        # Up
        if curr[0] - 1 >= 0 and (ground[curr[0] - 1][curr[1]] == "7" or ground[curr[0] - 1][curr[1]] == "|" or ground[curr[0] - 1][curr[1]] == "F"):
            valid.append((curr[0] - 1, curr[1]))   
    elif ground[curr[0]][curr[1]] == "L":
        # Right
        if curr[1] + 1 < len(ground[0]) and (ground[curr[0]][curr[1] + 1] == "J" or ground[curr[0]][curr[1] + 1] == "-" or ground[curr[0]][curr[1] + 1] == "7"):
            valid.append((curr[0], curr[1] + 1))
        # Up
        if curr[0] - 1 >= 0 and (ground[curr[0] - 1][curr[1]] == "7" or ground[curr[0] - 1][curr[1]] == "|" or ground[curr[0] - 1][curr[1]] == "F"):
            valid.append((curr[0] - 1, curr[1]))
    elif ground[curr[0]][curr[1]] == "F":
        # Right
        if curr[1] + 1 < len(ground[0]) and (ground[curr[0]][curr[1] + 1] == "J" or ground[curr[0]][curr[1] + 1] == "-" or ground[curr[0]][curr[1] + 1] == "7"):
            valid.append((curr[0], curr[1] + 1))
        # Down
        if curr[0] + 1 < len(ground) and (ground[curr[0] + 1][curr[1]] == "L" or ground[curr[0] + 1][curr[1]] == "|" or ground[curr[0] + 1][curr[1]] == "J"):
            valid.append((curr[0] + 1, curr[1]))
    elif ground[curr[0]][curr[1]] == "S":
        # Down
        if curr[0] + 1 < len(ground) and (ground[curr[0] + 1][curr[1]] == "L" or ground[curr[0] + 1][curr[1]] == "|" or ground[curr[0] + 1][curr[1]] == "J"):
            valid.append((curr[0] + 1, curr[1]))   
        # Up
        if curr[0] - 1 >= 0 and (ground[curr[0] - 1][curr[1]] == "7" or ground[curr[0] - 1][curr[1]] == "|" or ground[curr[0] - 1][curr[1]] == "F"):
            valid.append((curr[0] - 1, curr[1]))
        # Left
        if curr[1] - 1 >= 0 and (ground[curr[0]][curr[1] - 1] == "F" or ground[curr[0]][curr[1] - 1] == "-" or ground[curr[0]][curr[1] - 1] == "L"):
            valid.append((curr[0], curr[1] - 1))
        # Right
        if curr[1] + 1 < len(ground[0]) and (ground[curr[0]][curr[1] + 1] == "J" or ground[curr[0]][curr[1] + 1] == "-" or ground[curr[0]][curr[1] + 1] == "7"):
            valid.append((curr[0], curr[1] + 1))
    return valid     

def get_next_position(ground, curr, prev):
    valid_pos = get_all_valid_positions(ground, curr)
    for p in valid_pos:
        if p != prev and p is not None:
            return p

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        ground = []
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        start = None
        for l in range(len(lol)):
            better_line = lol[l].strip()
            ground.append(better_line)
            for c in range(len(better_line)):
                if better_line[c] == "S":
                    start = (l, c)
        
        prev = None
        curr = start
        while curr != start or out == 0:
            temp = curr
            curr = get_next_position(ground, curr, prev)
            prev = temp
            out += 1
            if curr is None:
                break

        print(out // 2)

if __name__ == "__main__":
    main() 