# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for line in lol:
            # Example: 'Card   1: 77 45  9 81 96  5 91  3 66 76 |  7 56 66 49 96 58 54 34 37  5 14 85 45 91  9 22 81 50 88 77 76  3 83 93 18\n'
            line = line.strip()
            important_part = line.split(": ")[1]
            winning_numbers = important_part.split("|")[0].split(" ")
            print(winning_numbers)
            my_numbers = important_part.split("|")[1].split(" ")
            print(my_numbers)
            curr = 0
            for num in winning_numbers:
                if num.isdigit() and num in my_numbers:
                    if curr == 0:
                        curr += 1
                    else:
                        curr *= 2
            out += curr
        print(out)

if __name__ == "__main__":
    main() 