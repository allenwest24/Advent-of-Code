# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        seeds = []
        mappings = {}
        curr_mapping = ""
        
        # Part 1: Create the mappings.
        for line in lol:
            if "seeds: " in line:
                seeds = line.strip().split("seeds: ")[1].split(" ")
                continue
            elif len(line.strip()) == 0:
                curr_mapping = ""
                continue
            elif "map:" in line:
                curr_mapping = line.strip().split(" ")[0]
                mappings[curr_mapping] = {}
                continue
            else:
                values = line.strip().split(" ")
                source = int(values[1])
                dest = int(values[0])
                curr_range = int(values[2])
                mappings[curr_mapping][source] = (dest, curr_range)
        
        # Part 2: Run the seeds.
        best = float("inf")
        for seed in seeds:
            curr_val_type = "seed"
            curr_val = int(seed)
            # Keep going until we have a location value.
            while curr_val_type != "location":
                for k in mappings:
                    if k.split("-")[0] == curr_val_type:
                        # If there is a mapping, take the new number, else remain the same
                        found_range = False
                        for kk in mappings[k]:
                            if curr_val >= kk and curr_val < kk + mappings[k][kk][1]:
                                curr_val = mappings[k][kk][0] + (curr_val - kk)
                                found_range = True
                                break
                        # Set the new type
                        curr_val_type = k.split("-")[2]
            
            # If we got this far we have the location.
            best = min(best, int(curr_val))
                        
        print(best)
        

if __name__ == "__main__":
    main() 