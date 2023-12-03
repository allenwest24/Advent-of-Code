# Main function definition.
def main():
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        nums = "1234567890"
        nums_to_sums = 0

        # Check each line
        for l in lol:
            first = None
            last = None

            # Parse each line for numbers
            for c in l:
                # If we find a num
                if c in nums:
                    # If this is the first found in the line, set as first
                    if first is None:
                        first = c
                    # No matter what, set as the most recent seen (last)
                    last = c
            
            # Squish the strings together then convert to an integer
            # Store within the list of integers to add at the end. 
            nums_to_sums += int(first + last)
        
        print(nums_to_sums)
    
if __name__ == "__main__":
    main() 