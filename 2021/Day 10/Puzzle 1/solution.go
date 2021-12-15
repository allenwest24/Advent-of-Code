package main

// List of needed imports.
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	total := 0
	// Load into a reusable format.
	for scanner.Scan() {
		line := scanner.Text()
		firstError := findError(line)

		if firstError == ")" {
			total += 3
		} else if firstError == "]" {
			total += 57
		} else if firstError == "}" {
			total += 1197
		} else if firstError == ">" {
			total += 25137
		}
	}

	fmt.Println(total)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
