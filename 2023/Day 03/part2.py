def total_adjacent_nums(lol, ii, jj):
    num_locations = {}
    out = 0
    # Define the range for checking adjacent cells.
    row_range = range(max(0, ii - 1), min(ii + 2, len(lol)))
    col_range = range(max(0, jj - 1), min(jj + 2, len(lol[0])))

    for i in row_range:
        running_num = False
        for j in col_range:
            # Skip the cell itself.
            if i == ii and j == jj:
                running_num = False
                continue
            
            if lol[i][j].isdigit():
                if not running_num:
                    num_locations[out] = (i, j)
                    out += 1
                running_num = True
            else:
                running_num = False

    return out, num_locations

def grow_left_while_nums(line, jj, total_number):
    while jj >= 0 and line[jj].isdigit():
        total_number = line[jj] + total_number
        jj -= 1
    return total_number

def grow_right_while_nums(line, jj, total_number):
    while jj < len(line) and line[jj].isdigit():
        total_number += line[jj]
        jj += 1
    return total_number

def sum_adjacent_nums(lol, locations):
    out = 1
    for l in locations:
        total_number = ""
        ii = locations[l][0]
        jj = locations[l][1]
        total_number = grow_left_while_nums(lol[ii], jj, total_number)
        total_number = grow_right_while_nums(lol[ii], jj + 1, total_number)  
        out *= int(total_number)
    return out

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for ii in range(len(lol)):
            for jj in range(len(lol[ii])):
                if lol[ii][jj] == "*":
                    num_touching, locations = total_adjacent_nums(lol, ii, jj)
                    if num_touching == 2:
                        out += sum_adjacent_nums(lol, locations)
        print(out)
    
if __name__ == "__main__":
    main() 