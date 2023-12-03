# Main function definition.
def main():
    out = 0

    # Open the input file as read only.
    with open('./input.txt', 'r') as f:

        # Reads each line into an array without the newline character 
        lol = f.readlines()
        for ii in range(len(lol)):
            good_game = True
            game = lol[ii].split(": ")[1]
            rounds = game.split("; ")
            max_cubes = {
                "red": 0,
                "green": 0,
                "blue": 0
            }
            for r in rounds:
                cube_counts = r.split(", ")
                for c in cube_counts:
                    cube_count = int(c.split(" ")[0])
                    cube_color = c.split(" ")[1].strip()
                    max_cubes[cube_color] = max(max_cubes[cube_color], cube_count)
            out += max_cubes["red"] * max_cubes["green"] * max_cubes["blue"]
    print(out)
    
if __name__ == "__main__":
    main() 