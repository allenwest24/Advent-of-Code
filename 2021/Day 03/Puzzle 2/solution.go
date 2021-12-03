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
	ogrCandidates := []string{}
	csrCandidates := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		line += "9"
		ogrCandidates = append(ogrCandidates, line)
		csrCandidates = append(csrCandidates, line)
	}

  
	// Solve the puzzle.
  // First get the oxygen generator rating.
	curr := 0
	ogrBuilder := ""
	for len(ogrCandidates) > 1 {
		bitCounter := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		for line := 0; line < len(ogrCandidates); line++ {
			for ii := 0; ii < 12; ii++ {
				if ogrCandidates[line][ii:ii+1] == "0" {
					bitCounter[ii][0]++
				} else {
					bitCounter[ii][1]++
				}
			}
		}
		ogrBuilder = ""
		for jj := 0; jj < 12; jj++ {
			if bitCounter[jj][0] > bitCounter[jj][1] {
				ogrBuilder += "0"
			} else {
				ogrBuilder += "1"
			}
		}
		ogrBuilder += "9"
		temp := []string{}
		for kk := 0; kk < len(ogrCandidates); kk++ {
			if ogrBuilder[curr:curr+1] == ogrCandidates[kk][curr:curr+1] {
				temp = append(temp, ogrCandidates[kk])
			}
		}
		curr++
		ogrCandidates = temp
	}
  // Should contain a single remainingvalue.
	fmt.Print(ogrCandidates)

  // Now onto the CO2 Scrubber Rating.
	curr = 0
	csrBuilder := ""
	for len(csrCandidates) > 1 {
		bitCounter := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		for line := 0; line < len(csrCandidates); line++ {
			for ii := 0; ii < 12; ii++ {
				if csrCandidates[line][ii:ii+1] == "0" {
					bitCounter[ii][0]++
				} else {
					bitCounter[ii][1]++
				}
			}
		}
		csrBuilder = ""
		for jj := 0; jj < 12; jj++ {
			if bitCounter[jj][0] > bitCounter[jj][1] {
				csrBuilder += "1"
			} else {
				csrBuilder += "0"
			}
		}
		csrBuilder += "9"
		temp := []string{}
		for ll := 0; ll < len(csrCandidates); ll++ {
			if csrBuilder[curr:curr+1] == csrCandidates[ll][curr:curr+1] {
				temp = append(temp, csrCandidates[ll])
			}
		}
		curr++
		csrCandidates = temp
	}
  // Should contain a single remainingvalue.
	fmt.Print(csrCandidates)

  // Reformat from binary.
	ogrRate, _ := strconv.ParseInt(ogrCandidates[0][0:len(ogrBuilder)-1], 2, 64)
	csrRate, _ := strconv.ParseInt(csrCandidates[0][0:len(csrBuilder)-1], 2, 64)
  // Multiply together to get the Life Support Rating.
	lsr := ogrRate * csrRate

	// Print the results.
	fmt.Print(lsr)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
