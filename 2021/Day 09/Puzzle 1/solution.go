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
	f, err := os.Open("day9puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	floor := [][]int{}
	total := 0

	// Load into a reusable format.
	for scanner.Scan() {
		line := scanner.Text()
		horiz := []int{}
		for ii := 0; ii < len(line); ii++ {
			num, _ := strconv.Atoi(string(line[ii]))
			horiz = append(horiz, num)
		}
		floor = append(floor, horiz)
	}

	for ii := 0; ii < len(floor); ii++ {
		for jj := 0; jj < len(floor[ii]); jj++ {
			min := true
			if jj > 0 {
				if floor[ii][jj-1] <= floor[ii][jj] {
					min = false
				}
			}
			if jj < len(floor[ii])-1 {
				if floor[ii][jj+1] <= floor[ii][jj] {
					min = false
				}
			}
			if ii > 0 {
				if floor[ii-1][jj] <= floor[ii][jj] {
					min = false
				}
			}
			if ii < len(floor)-1 {
				if floor[ii+1][jj] <= floor[ii][jj] {
					min = false
				}
			}
			if min {
				total += (floor[ii][jj] + 1)
			}
		}
	}

	fmt.Println(total)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
