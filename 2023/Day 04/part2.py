# Main function definition.
def main():
    out = 0
    # 0-indexed card num : count of card
    card_holder = {}
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for ii in range(len(lol)):
            # Example: 'Card   1: 77 45  9 81 96  5 91  3 66 76 |  7 56 66 49 96 58 54 34 37  5 14 85 45 91  9 22 81 50 88 77 76  3 83 93 18\n'
            card_holder[ii] = 1

        for jj in card_holder:
            line = lol[jj].strip()
            important_part = line.split(": ")[1]
            winning_numbers = important_part.split("|")[0].split(" ")
            my_numbers = important_part.split("|")[1].split(" ")
            curr_wins = 0
            for num in winning_numbers:
                if num.isdigit() and num in my_numbers:
                    curr_wins += 1
            if curr_wins > 0:
                for kk in range(1, curr_wins + 1):
                    if jj + kk < len(card_holder):
                        card_holder[jj + kk] += 1 * card_holder[jj]
            
        for ll in card_holder:
            out += card_holder[ll]
        print(card_holder)
        print(out)

if __name__ == "__main__":
    main() 