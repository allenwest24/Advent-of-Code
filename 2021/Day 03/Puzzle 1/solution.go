package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day3puzzle1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Solve the puzzle.
	// Data sturcture to hold the counts for each position.
	bitCounter := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	// Format the bitCounter.
	for scanner.Scan() {
		line := scanner.Text()
		// Need to add on to end in order to grab the last subsection of the string.
		line += "9"
		for ii := 0; ii < 12; ii++ {
			if line[ii:ii+1] == "0" {
				bitCounter[ii][0]++
			} else {
				bitCounter[ii][1]++
			}
		}
	}

	gammaBuilder := ""
	epsilonBuilder := ""
	// Now we need to make the gamma and epsilon rates.
	for jj := 0; jj < 12; jj++ {
		if bitCounter[jj][0] > bitCounter[jj][1] {
			gammaBuilder += "0"
			epsilonBuilder += "1"
		} else {
			gammaBuilder += "1"
			epsilonBuilder += "0"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaBuilder, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonBuilder, 2, 64)
	powerConsumption := gammaRate * epsilonRate

	// Print the results.
	fmt.Print(powerConsumption)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
