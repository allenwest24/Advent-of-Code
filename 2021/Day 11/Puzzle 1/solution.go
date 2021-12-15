package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func flash(octos [][]int, ii int, jj int) [][]int {
	tempOctos := octos
	octos[ii][jj] = 0
	if ii > 0 && jj > 0 && tempOctos[ii-1][jj-1] != 0 {
		tempOctos[ii-1][jj-1]++
	}
	if ii > 0 && tempOctos[ii-1][jj] != 0 {
		tempOctos[ii-1][jj]++
	}
	if ii > 0 && jj < len(tempOctos[ii])-1 && tempOctos[ii-1][jj+1] != 0 {
		tempOctos[ii-1][jj+1]++
	}
	if jj > 0 && tempOctos[ii][jj-1] != 0 {
		tempOctos[ii][jj-1]++
	}
	if jj < len(tempOctos[ii])-1 && tempOctos[ii][jj+1] != 0 {
		tempOctos[ii][jj+1]++
	}
	if ii < len(tempOctos)-1 && jj > 0 && tempOctos[ii+1][jj-1] != 0 {
		tempOctos[ii+1][jj-1]++
	}
	if ii < len(tempOctos)-1 && tempOctos[ii+1][jj] != 0 {
		tempOctos[ii+1][jj]++
	}
	if ii < len(tempOctos)-1 && jj < len(tempOctos[ii])-1 && tempOctos[ii+1][jj+1] != 0 {
		tempOctos[ii+1][jj+1]++
	}
	if ii > 0 && jj > 0 && tempOctos[ii-1][jj-1] > 9 {
		tempOctos = flash(tempOctos, ii-1, jj-1)
	}
	if ii > 0 && tempOctos[ii-1][jj] > 9 {
		tempOctos = flash(tempOctos, ii-1, jj)
	}
	if ii > 0 && jj < len(tempOctos[ii])-1 && tempOctos[ii-1][jj+1] > 9 {
		tempOctos = flash(tempOctos, ii-1, jj+1)
	}
	if jj > 0 && tempOctos[ii][jj-1] > 9 {
		tempOctos = flash(tempOctos, ii, jj-1)
	}
	if jj < len(tempOctos[ii])-1 && tempOctos[ii][jj+1] > 9 {
		tempOctos = flash(tempOctos, ii, jj+1)
	}
	if ii < len(tempOctos)-1 && jj > 0 && tempOctos[ii+1][jj-1] > 9 {
		tempOctos = flash(tempOctos, ii+1, jj-1)
	}
	if ii < len(tempOctos)-1 && tempOctos[ii+1][jj] > 9 {
		tempOctos = flash(tempOctos, ii+1, jj)
	}
	if ii < len(tempOctos)-1 && jj < len(tempOctos[ii])-1 && tempOctos[ii+1][jj+1] > 9 {
		tempOctos = flash(tempOctos, ii+1, jj+1)
	}
	return tempOctos
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day11puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	octos := [][]int{}
	totalFlashes := 0
	// Load into a reusable format.
	for scanner.Scan() {
		line := scanner.Text()
		newOctoLine := []int{}
		for ii := 0; ii < len(line); ii++ {
			octoAsNum, _ := strconv.Atoi(string(line[ii]))
			newOctoLine = append(newOctoLine, octoAsNum)
		}
		octos = append(octos, newOctoLine)
	}

	for x := 0; x < 100; x++ {
		for ii := 0; ii < len(octos); ii++ {
			for jj := 0; jj < len(octos[ii]); jj++ {
				octos[ii][jj]++
			}
		}

		for ii := 0; ii < len(octos); ii++ {
			for jj := 0; jj < len(octos[ii]); jj++ {
				if octos[ii][jj] > 9 {
					octos = flash(octos, ii, jj)
				}
			}
		}

		for ii := 0; ii < len(octos); ii++ {
			for jj := 0; jj < len(octos[ii]); jj++ {
				if octos[ii][jj] == 0 {
					totalFlashes++
				}
			}
		}
	}

	fmt.Println(totalFlashes)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
