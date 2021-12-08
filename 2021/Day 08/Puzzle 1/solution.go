package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day8puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	total1478 := 0
	for scanner.Scan() {
		line := scanner.Text()
		outputIndex := strings.Index(line, "|")
		outputIndex += 2
		output := line[outputIndex:]
		outputSegments := strings.Split(output, " ")
		for seg := 0; seg < len(outputSegments); seg++ {
			if len(outputSegments[seg]) == 2 || len(outputSegments[seg]) == 3 || len(outputSegments[seg]) == 4 || len(outputSegments[seg]) == 7 {
				total1478++
			}
		}
	}
	fmt.Println(total1478)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
