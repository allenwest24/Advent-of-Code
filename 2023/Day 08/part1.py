# Main function definition.
def main():
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        
        # Build data structure
        dir_map = {}
        pattern = lol[0].strip()
        directions = lol[2:]
        for d in directions:
            key = d.strip().split(" ")[0]
            tup = (d.strip().split(" ")[2].replace("(", "").replace(",", ""), d.strip().split(" ")[3].replace(")", ""))
            dir_map[key] = tup
        
        total = 0
        curr = "AAA"
        while curr != "ZZZ":
            lor = pattern[total % len(pattern)]
            if lor == "L":
                curr = dir_map[curr][0]
            else:
                curr = dir_map[curr][1]
            total += 1
        print(total)


if __name__ == "__main__":
    main() 