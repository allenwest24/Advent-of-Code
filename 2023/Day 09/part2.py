def get_all_differences(list_of_ints):
    return [list_of_ints[i+1] - list_of_ints[i] for i in range(len(list_of_ints)-1)]

def is_all_zeros(list_of_ints):
    for i in list_of_ints:
        if i != 0:
            return False
    return True

def make_all_ints(list_of_strings):
    return [int(i) for i in list_of_strings]

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for l in lol:
            sequences = []
            curr = make_all_ints(l.strip().split(" "))
            sequences.append(curr)
            while not is_all_zeros(curr):
                curr = get_all_differences(curr)
                sequences.append(curr)
                
            x = len(sequences) - 1
            while x > 0:
                last_next = sequences[x-1][0] - sequences[x][0]
                sequences[x-1] = [last_next] + sequences[x-1]
                x -= 1

            out += sequences[0][0]
        print(out)

if __name__ == "__main__":
    main() 