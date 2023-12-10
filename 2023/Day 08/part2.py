def check_if_all_end_in_z(tracker):
    for t in tracker:
        if tracker[t][-1] != "Z":
            return False
    return True

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
        starters = []
        for d in directions:
            key = d.strip().split(" ")[0]
            tup = (d.strip().split(" ")[2].replace("(", "").replace(",", ""), d.strip().split(" ")[3].replace(")", ""))
            dir_map[key] = tup
            if key[-1] == "A":
                starters.append(key)
        
        tracker = {}
        total = 0
        for s in starters:
            tracker[s] = s
        while not check_if_all_end_in_z(tracker):
            lor = pattern[total % len(pattern)]
            for t in tracker:
                if lor == "L":
                    tracker[t] = dir_map[tracker[t]][0]
                else:
                    tracker[t] = dir_map[tracker[t]][1]
            total += 1
        print(total)

if __name__ == "__main__":
    main() 