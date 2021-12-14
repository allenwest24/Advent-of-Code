package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func allWithin(str1 string, str2 string) bool {
	isWithin := true
	str1 += "Z"
	str2 += "Z"
	for ii := 0; ii < len(str1)-1; ii++ {
		if !isWithin {
			return false
		}
		isWithin = false
		for jj := 0; jj < len(str2)-1; jj++ {
			if str1[ii:ii+1] == str2[jj:jj+1] {
				isWithin = true
			}
		}
	}
	return isWithin
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day8puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0

	// Load into a reusable format.
	for scanner.Scan() {
		line := scanner.Text()
		curr1 := ""
		curr4 := ""
		curr7 := ""
		curr6 := ""
		outputIndex := strings.Index(line, "|") + 2
		inputIndex := strings.Index(line, "|") - 1
		output := line[outputIndex:]
		input := line[:inputIndex]

		inputSegments := strings.Split(input, " ")
		for ii := 0; ii < len(inputSegments); ii++ {
			if len(inputSegments[ii]) == 2 {
				curr1 = inputSegments[ii]
			} else if len(inputSegments[ii]) == 3 {
				curr7 = inputSegments[ii]
			} else if len(inputSegments[ii]) == 4 {
				curr4 = inputSegments[ii]
			}
		}

		for ii := 0; ii < len(inputSegments); ii++ {
			if !allWithin(curr7, inputSegments[ii]) && len(inputSegments[ii]) == 6 {
				curr6 = inputSegments[ii]
			}
		}

		outputSegments := strings.Split(output, " ")
		tempTotal := 0
		for seg := 0; seg < len(outputSegments); seg++ {
			if len(outputSegments[seg]) == 2 {
				tempTotal = (tempTotal * 10) + 1
			} else if len(outputSegments[seg]) == 3 {
				tempTotal = (tempTotal * 10) + 7
			} else if len(outputSegments[seg]) == 4 {
				tempTotal = (tempTotal * 10) + 4
			} else if len(outputSegments[seg]) == 5 {
				if allWithin(curr1, outputSegments[seg]) {
					tempTotal = (tempTotal * 10) + 3
				} else if allWithin(outputSegments[seg], curr6) {
					tempTotal = (tempTotal * 10) + 5
				} else {
					tempTotal = (tempTotal * 10) + 2
				}
			} else if len(outputSegments[seg]) == 6 {
				if !allWithin(curr7, outputSegments[seg]) {
					tempTotal = (tempTotal * 10) + 6
				} else if !allWithin(curr4, outputSegments[seg]) {
					tempTotal = (tempTotal * 10) + 0
				} else {
					tempTotal = (tempTotal * 10) + 9
				}
			} else if len(outputSegments[seg]) == 7 {
				tempTotal = (tempTotal * 10) + 8
			}
		}
		total += tempTotal
	}
	fmt.Println(total)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
