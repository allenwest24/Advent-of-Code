# Refactoring the solution with optimizations

# Function to apply the mapping to a number with memoization
def process_seed_memoized(seed, mappings, memo):
    if seed in memo:
        return memo[seed]

    curr_val_type = "seed"
    curr_val = seed
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
                break
    memo[seed] = curr_val
    return curr_val

# Main function with optimizations
def main_optimized():
    with open('./input.txt', 'r') as f:
        lines = f.readlines()

    seeds = []
    mappings = {}
    curr_mapping = ""

    # Creating the mappings
    for line in lines:
        if "seeds: " in line:
            seed_ranges = line.strip().split("seeds: ")[1].split(" ")
            s = 0
            while s < len(seed_ranges) // 2:
                curr_seed = int(seed_ranges[s])
                curr_seed_range = int(seed_ranges[s + 1])
                seeds.append((curr_seed, curr_seed + curr_seed_range))
                s += 2
        elif len(line.strip()) == 0:
            curr_mapping = ""
        elif "map:" in line:
            curr_mapping = line.strip().split(" ")[0]
            mappings[curr_mapping] = {}
        else:
            values = line.strip().split(" ")
            source = int(values[1])
            dest = int(values[0])
            curr_range = int(values[2])
            mappings[curr_mapping][source] = (dest, curr_range)

    # Memoization dictionary
    memo = {}

    # Process seeds using range optimization
    best = float("inf")
    for start, end in seeds:
        for seed in range(start, end):
            location = process_seed_memoized(seed, mappings, memo)
            best = min(best, location)

    return best

# Execute the optimized main function
main_optimized()
