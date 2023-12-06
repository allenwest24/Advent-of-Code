# Main function definition.
def main():
    out = 0
    # Open the input file as read only.
    with open('./input.txt', 'r') as f:
        # Reads each line into an array without the newline character 
        lol = f.readlines()
        times = lol[0].strip().split(" ")

        # Ugly code, but removes the non-digit characters.
        temp = []
        for ii in times:
            if ii.isdigit():
                temp.append(int(ii))
        times = temp
        distances = lol[1].strip().split(" ")
        temp = []
        for ii in distances:
            if ii.isdigit():
                temp.append(int(ii))
        distances = temp

        # Do the actual work now
        out = 1
        for ii in range(len(times)):
            d = distances[ii]
            t = times[ii]
            counter = 0
            for x in range(1, t):
                if x * (t - x) > d:
                    counter += 1
            if counter > 0:
                out *= counter
        print(out)

if __name__ == "__main__":
    main() 