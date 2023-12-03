def any_adjacent_symbol(lol, ii, jj):
    symbols = "!@#$%^&*()_+-=,/<>?;':[]{}\\|`~"
    # Define the range for checking adjacent cells.
    row_range = range(max(0, ii - 1), min(ii + 2, len(lol)))
    col_range = range(max(0, jj - 1), min(jj + 2, len(lol[0])))

    for i in row_range:
        for j in col_range:
            # Skip the cell itself.
            if i == ii and j == jj:
                continue
            
            # Check if the cell contains any of the specified symbols.
            if lol[i][j] in symbols:
                return True

    return False

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        adjacent_symbol = False
        for ii in range(len(lol)):
            curr = ""
            for jj in range(len(lol[ii])):
                if lol[ii][jj].isdigit():
                    curr += lol[ii][jj]
                    if any_adjacent_symbol(lol, ii, jj):
                        adjacent_symbol = True
                # We ran out of a symbol, now we want to review whether we should add it.
                elif len(curr) > 0:
                    if adjacent_symbol:
                        out += int(curr)
                    # Either way we reset the current trackers.
                    curr = ""
                    adjacent_symbol = False
                # At the end of each line we will need to stop and assess.
                if jj == len(lol[ii]) - 1 and adjacent_symbol and len(curr) > 0:
                    out += int(curr)
        print(out)
    
if __name__ == "__main__":
    main() 