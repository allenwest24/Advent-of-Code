from itertools import product

def char_frequencies(s):
    freq = {}
    for char in s:
        if char in freq:
            freq[char] += 1
        else:
            freq[char] = 1
    return freq

def hand_rank(hand):
    hierarchy = "AKQT98765432J"
    rank_map = {ch: i for i, ch in enumerate(hierarchy)}
    return tuple(rank_map[ch] for ch in hand)

def sort_by_hierarchy(loh):
    return sorted(loh, key=hand_rank)

def get_hand_analysis(hand):
    counts = char_frequencies(hand)
    if 5 in counts.values():
        return "5oak"
    elif 4 in counts.values():
        return "4oak"
    elif 3 in counts.values() and 2 in counts.values():
        return "full"
    elif 3 in counts.values():
        return "3oak"
    elif list(counts.values()).count(2) == 2:
        return "2pair"
    elif 2 in counts.values():
        return "pair"
    else:
        return "high"

def find_indices_of_char(string, char):
    return [index for index, c in enumerate(string) if c == char]

def substitute_characters(original_string, indices, substitution_chars):
    # Create a list of characters from the original string
    string_chars = list(original_string)

    # Replace characters at the specified indices
    for index, replacement in zip(indices, substitution_chars):
        string_chars[index] = replacement

    # Join the characters back into a string
    return ''.join(string_chars)

def generate_substitutions(string, indices, substitution_string):
    # Generate all combinations of substitutions
    all_combinations = product(substitution_string, repeat=len(indices))

    # Apply each combination to the original string
    substituted_strings = [substitute_characters(string, indices, combo) for combo in all_combinations]

    return substituted_strings

def get_best_outcome_for_jokers(hand):
    hierarchy = "AKQT98765432J"
    rankings_ordered = ["5oak", "4oak", "full", "3oak", "2pair", "pair", "high"]
    best = ""
    J_locations = find_indices_of_char(hand, "J")
    all_combos = generate_substitutions(hand, J_locations, hierarchy)
    for combo in all_combos:
        analysis = get_hand_analysis(combo)
        if best == "" or rankings_ordered.index(analysis) < rankings_ordered.index(best):
            best = analysis
    if best == "":
        return get_hand_analysis(hand)
    return best
    
# Main function definition.
def main():
    rankings = {
        "5oak": [],
        "4oak": [],
        "full": [],
        "3oak": [],
        "2pair": [],
        "pair": [],
        "high": []
    }
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

        for hand in hands:
            counts = char_frequencies(hand)
            best_rank = get_best_outcome_for_jokers(hand)
            rankings[best_rank].append(hand)

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