package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day2puzzle2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Solve the puzzle.
	totalDistance := 0
	totalAltitude := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		direction := split[0]
		amount, _ := strconv.Atoi(split[1])
		if direction == "up" {
			aim -= amount
		} else if direction == "down" {
			aim += amount
		} else if direction == "forward" {
			totalDistance += amount
			totalAltitude += (aim * amount)
		}
	}
	location := totalDistance * totalAltitude
	// Print the results.
	fmt.Print(location)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
