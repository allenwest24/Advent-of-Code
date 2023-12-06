# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()

        # Ugly code, but removes the non-digit characters.
        times = lol[0].strip().split(" ")
        time = ""
        for ii in times:
            if ii.isdigit():
                time += ii
        distances = lol[1].strip().split(" ")
        distance = ""
        for ii in distances:
            if ii.isdigit():
                distance += ii

        t = int(time)
        d = int(distance)
        # Do the actual work now
        c = 0
        for x in range(1, t):
            if x * (t - x) > d:
                c += 1
        print(c)

if __name__ == "__main__":
    main() 