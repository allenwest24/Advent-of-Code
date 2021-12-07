package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func testAllFuelScenarios(crabSubs []int, maxDist int) int {
	best := 0
	currTotal := 0
	for ii := 0; ii < maxDist; ii++ {
		currTotal = 0
		for jj := 0; jj < len(crabSubs); jj++ {
			currTotal += int(math.Abs(float64(ii - crabSubs[jj])))
		}
		if currTotal < best || best == 0 {
			best = currTotal
		}
	}

	return best
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day7puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	crabSubs := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, ",")
		for ii := 0; ii < len(input); ii++ {
			crab, _ := strconv.Atoi(input[ii])
			crabSubs = append(crabSubs, crab)
		}
	}

	best := testAllFuelScenarios(crabSubs, 2000)
	fmt.Println(best)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
