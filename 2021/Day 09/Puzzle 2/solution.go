package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func multiplyTop3(basinTracker [][]int) int {
	lengthTracker := []int{}
	for ii := 0; ii < len(basinTracker); ii++ {
		lengthTracker = append(lengthTracker, len(basinTracker[ii]))
	}
	fmt.Println(lengthTracker)

	best1 := 0
	best2 := 0
	best3 := 0

	for ii := 0; ii < len(lengthTracker); ii++ {
		if lengthTracker[ii] >= best1 {
			best3 = best2
			best2 = best1
			best1 = lengthTracker[ii]
		} else if lengthTracker[ii] >= best2 {
			best3 = best2
			best2 = lengthTracker[ii]
		} else if lengthTracker[ii] >= best3 {
			best3 = lengthTracker[ii]
		}
	}
	fmt.Println(best1)
	fmt.Println(best2)
	fmt.Println(best3)

	return (best1 * best2 * best3)
}

func isWithin(x int, basin []int) bool {
	for ii := 0; ii < len(basin); ii++ {
		if basin[ii] == x {
			return true
		}
	}
	return false
}

func withinSeen(x int, lsOfBasin [][]int) bool {
	for ii := 0; ii < len(lsOfBasin); ii++ {
		if isWithin(x, lsOfBasin[ii]) {
			return true
		}
	}
	return false
}

func placeAppropriately(basinTracker [][]int, floor [][]int, ii int, jj int) [][]int {
	update := basinTracker
	for x := 0; x < len(basinTracker); x++ {
		if jj > 0 {
			if isWithin((ii*1000)+(jj-1), update[x]) {
				update[x] = append(update[x], (ii*1000)+(jj))
				return update
			}
		}
		if jj < len(floor[ii])-1 {
			if isWithin((ii*1000)+(jj+1), update[x]) {
				update[x] = append(update[x], (ii*1000)+(jj))
				return update
			}
		}
		if ii > 0 {
			if isWithin(((ii-1)*1000)+jj, update[x]) {
				update[x] = append(update[x], (ii*1000)+jj)
				return update
			}
		}
		if ii < len(floor)-1 {
			if isWithin(((ii+1)*1000)+jj, update[x]) {
				update[x] = append(update[x], (ii*1000)+jj)
				return update
			}
		}
	}
	newBasin := []int{((ii * 1000) + jj)}
	update = append(update, newBasin)
	return update
}

func fillBasin(basinTracker [][]int, floor [][]int, ii int, jj int) [][]int {
	update := placeAppropriately(basinTracker, floor, ii, jj)
	if jj > 0 {
		if floor[ii][jj-1] != 9 && !withinSeen((ii*1000)+(jj-1), update) {
			update = fillBasin(update, floor, ii, jj-1)
		}
	}
	if jj < len(floor[ii])-1 {
		if floor[ii][jj+1] != 9 && !withinSeen((ii*1000)+(jj+1), update) {
			update = fillBasin(update, floor, ii, jj+1)
		}
	}
	if ii > 0 {
		if floor[ii-1][jj] != 9 && !withinSeen(((ii-1)*1000)+jj, update) {
			update = fillBasin(update, floor, ii-1, jj)
		}
	}
	if ii < len(floor)-1 {
		if floor[ii+1][jj] != 9 && !withinSeen(((ii+1)*1000)+jj, update) {
			update = fillBasin(update, floor, ii+1, jj)
		}
	}
	return update
}

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
	basinTracker := [][]int{}

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
			fmt.Println(jj)
			if floor[ii][jj] != 9 && !withinSeen((ii*1000)+jj, basinTracker) {
				basinTracker = fillBasin(basinTracker, floor, ii, jj)
			}
		}
	}

	out := multiplyTop3(basinTracker)
	fmt.Println(out)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
