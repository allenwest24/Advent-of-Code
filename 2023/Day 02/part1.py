# Main function definition.
def main():
    # Parameters for the question.
    cubes = {
        "red": 12,
        "green": 13,
        "blue": 14
    }
    out = 0

    # Open the input file as read only.
    with open('./input.txt', 'r') as f:

        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for ii in range(len(lol)):
            good_game = True
            game = lol[ii].split(": ")[1]
            rounds = game.split("; ")
            for r in rounds:
                cube_counts = r.split(", ")
                for c in cube_counts:
                    cube_count = int(c.split(" ")[0])
                    cube_color = c.split(" ")[1].strip()
                    if cubes[cube_color] < cube_count:
                        good_game = False
            if good_game:
                out += ii + 1
    print(out)



    
if __name__ == "__main__":
    main() 