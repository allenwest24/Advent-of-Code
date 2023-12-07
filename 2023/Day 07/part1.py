def char_frequencies(s):
    freq = {}
    for char in s:
        if char in freq:
            freq[char] += 1
        else:
            freq[char] = 1
    return freq

def hand_rank(hand):
    hierarchy = "AKQJT98765432"
    rank_map = {ch: i for i, ch in enumerate(hierarchy)}
    return tuple(rank_map[ch] for ch in hand)

def sort_by_hierarchy(loh):
    return sorted(loh, key=hand_rank)

# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        hands = {}
        for line in lol:
            hand = line.strip()
            cards = hand.split(" ")[0]
            bet = hand.split(" ")[1]
            hands[cards] = bet
        
        rankings = {
            "5oak": [],
            "4oak": [],
            "full": [],
            "3oak": [],
            "2pair": [],
            "pair": [],
            "high": []
        }

        for hand in hands:
            counts = char_frequencies(hand)
            if 5 in counts.values():
                rankings["5oak"].append(hand)
            elif 4 in counts.values():
                rankings["4oak"].append(hand)
            elif 3 in counts.values() and 2 in counts.values():
                rankings["full"].append(hand)
            elif 3 in counts.values():
                rankings["3oak"].append(hand)
            elif list(counts.values()).count(2) == 2:
                rankings["2pair"].append(hand)
            elif 2 in counts.values():
                rankings["pair"].append(hand)
            else:
                rankings["high"].append(hand)

        # Sort each had by hierarchy.
        for rank in rankings:
            rankings[rank] = sort_by_hierarchy(rankings[rank])

        multiplier = len(lol)
        for rank in rankings:
            for hand in rankings[rank]:
                print(f"Multiplier {multiplier} for {hand} with bet {hands[hand]}")
                out += int(hands[hand]) * multiplier
                multiplier -= 1
        print(out)

if __name__ == "__main__":
    main() 