package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func scoreFinishingLine(line string) int {
	total := 0
	for ii := len(line) - 1; ii >= 0; ii-- {
		total *= 5
		if string(line[ii]) == "(" {
			total += 1
		} else if string(line[ii]) == "[" {
			total += 2
		} else if string(line[ii]) == "{" {
			total += 3
		} else if string(line[ii]) == "<" {
			total += 4
		}
	}
	return total
}

func findFinishingLine(line string) string {
	tracker := ""
	for ii := 0; ii < len(line); ii++ {
		if string(line[ii]) == "(" || string(line[ii]) == "[" || string(line[ii]) == "{" || string(line[ii]) == "<" {
			tracker += string(line[ii])
		} else {
			tracker = tracker[:len(tracker)-1]
		}
	}
	return tracker
}

func findError(line string) string {
	tracker := []string{}
	for ii := 0; ii < len(line); ii++ {
		if string(line[ii]) == "(" || string(line[ii]) == "[" || string(line[ii]) == "{" || string(line[ii]) == "<" {
			tracker = append(tracker, string(line[ii]))
		} else if string(line[ii]) == ")" {
			if tracker[len(tracker)-1] != "(" {
				return ")"
			} else {
				tracker = tracker[:len(tracker)-1]
			}
		} else if string(line[ii]) == "]" {
			if tracker[len(tracker)-1] != "[" {
				return "]"
			} else {
				tracker = tracker[:len(tracker)-1]
			}
		} else if string(line[ii]) == "}" {
			if tracker[len(tracker)-1] != "{" {
				return "}"
			} else {
				tracker = tracker[:len(tracker)-1]
			}
		} else if string(line[ii]) == ">" {
			if tracker[len(tracker)-1] != "<" {
				return ">"
			} else {
				tracker = tracker[:len(tracker)-1]
			}
		}
	}
	return ""
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day10puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	finishingLineScores := []int{}
	// Load into a reusable format.
	for scanner.Scan() {
		line := scanner.Text()
		firstError := findError(line)

		if firstError == "" {
			finishingLine := findFinishingLine(line)
			score := scoreFinishingLine(finishingLine)
			finishingLineScores = append(finishingLineScores, score)
		}
	}

	sort.Ints(finishingLineScores)
	fmt.Println(finishingLineScores[(len(finishingLineScores) / 2)])

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
