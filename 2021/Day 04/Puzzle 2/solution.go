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

func wonByColumn(board [][]int) bool {
	for column := 0; column < len(board); column++ {
		count := 0
		for row := 0; row < len(board); row++ {
			if board[row][column] == -1 {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func wonByRow(board [][]int) bool {
	for row := 0; row < len(board); row++ {
		count := 0
		for column := 0; column < len(board); column++ {
			if board[row][column] == -1 {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func aNotWithinB(a int, b []int) bool {
	isNotWithin := true
	for ii := 0; ii < len(b); ii++ {
		if b[ii] == a {
			isNotWithin = false
		}
	}
	return isNotWithin
}

func checkForLastWinner(listOfBoard [][][]int, winnersSoFar []int) []int {
	temp := winnersSoFar
	for ii := 0; ii < len(listOfBoard); ii++ {
		if (wonByColumn(listOfBoard[ii]) || wonByRow(listOfBoard[ii])) && (aNotWithinB(ii, winnersSoFar)) {
			temp = append(temp, ii)
		}
	}
	return temp
}

func sumUncalled(winningBoard [][]int) int {
	total := 0
	for ii := 0; ii < 5; ii++ {
		for jj := 0; jj < 5; jj++ {
			if winningBoard[ii][jj] != -1 {
				total += winningBoard[ii][jj]
			}
		}
	}
	return total
}

// Plays all of the games and returns the score designated by the instructions.
func playAllGames(numberQueue []int, listOfBoard [][][]int) int {
	winnersSoFar := []int{}
	boardHasWon := -1
	currVal := -1
	for num := 0; num < len(numberQueue); num++ {
		currVal = numberQueue[num]
		for ii := 0; ii < len(listOfBoard); ii++ {
			for jj := 0; jj < 5; jj++ {
				for kk := 0; kk < 5; kk++ {
					if listOfBoard[ii][jj][kk] == currVal {
						listOfBoard[ii][jj][kk] = -1
					}
				}
			}
		}
		// Check if any board has won.
		winnersSoFar = checkForLastWinner(listOfBoard, winnersSoFar)
		if len(winnersSoFar) == len(listOfBoard) {
			boardHasWon = winnersSoFar[len(winnersSoFar)-1]
			break
		}
	}
	sumOfUncalled := sumUncalled(listOfBoard[boardHasWon])
	out := (sumOfUncalled * currVal)
	return out
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day4puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	// Format the input into data structures to use easier.
	lineTracker := 0
	numberQueueAsStrings := []string{}
	numberQueue := []int{}
	listOfBoard := [][][]int{{{}}}

	for scanner.Scan() {
		line := scanner.Text()
		// Handles numberQueue to creat a list of numbers for us to use as the bingos.
		if lineTracker == 0 {
			numberQueueAsStrings = strings.Split(line, ",")
			for num := 0; num < len(numberQueueAsStrings); num++ {
				temp, _ := strconv.Atoi(numberQueueAsStrings[num])
				numberQueue = append(numberQueue, temp)
			}
		} else if !(((lineTracker - 1) % 6) == 0) {
			lineOfBoard := (((lineTracker - 1) % 6) - 1)
			if lineOfBoard == 0 {
				newBoard := [][]int{{}}
				listOfBoard = append(listOfBoard, newBoard)
			}
			boardLineAsStrings := strings.Split(line, " ")
			boardLine := []int{}
			for ii := 0; ii < len(boardLineAsStrings); ii++ {
				if len(boardLineAsStrings[ii]) != 0 {
					temp, _ := strconv.Atoi(boardLineAsStrings[ii])
					boardLine = append(boardLine, temp)
				}
			}
			listOfBoard[len(listOfBoard)-1] = append(listOfBoard[len(listOfBoard)-1], boardLine)
		}
		lineTracker++
	}
	// Formatting needed because I am bad.
	listOfBoard = listOfBoard[1:]
	for board := 0; board < len(listOfBoard); board++ {
		listOfBoard[board] = listOfBoard[board][1:]
	}
	for ii := 0; ii < len(listOfBoard); ii++ {
		fmt.Println(listOfBoard[ii])
	}

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	outcome := playAllGames(numberQueue, listOfBoard)
	fmt.Println(outcome)
}
