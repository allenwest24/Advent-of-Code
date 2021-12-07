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

func markYs(coords []int, ventMap [][]string) [][]string {
	temp := ventMap
	min := 0
	max := 0
	if coords[1] < coords[3] {
		min = coords[1]
		max = coords[3]
	} else {
		min = coords[3]
		max = coords[1]
	}
	for ii := min; ii <= max; ii++ {
		if temp[coords[0]][ii] == "." {
			temp[coords[0]][ii] = "1"
		} else {
			curr, _ := strconv.Atoi(temp[coords[0]][ii])
			curr++
			new := strconv.Itoa(curr)
			temp[coords[0]][ii] = new
		}
	}
	return temp
}

func markXs(coords []int, ventMap [][]string) [][]string {
	temp := ventMap
	min := 0
	max := 0
	if coords[0] < coords[2] {
		min = coords[0]
		max = coords[2]
	} else {
		min = coords[2]
		max = coords[0]
	}
	for ii := min; ii <= max; ii++ {
		if temp[ii][coords[1]] == "." {
			temp[ii][coords[1]] = "1"
		} else {
			curr, _ := strconv.Atoi(temp[ii][coords[1]])
			curr++
			new := strconv.Itoa(curr)
			temp[ii][coords[1]] = new
		}
	}
	return temp
}

func markDiagonal(coords []int, ventMap [][]string) [][]string {
	temp := ventMap
	caseNum := 0
	if coords[0] < coords[2] && coords[1] < coords[3] {
		caseNum = 1
	} else if coords[0] > coords[2] && coords[1] < coords[3] {
		caseNum = 2
	} else if coords[0] < coords[2] && coords[1] > coords[3] {
		caseNum = 3
	} else if coords[0] > coords[2] && coords[1] > coords[3] {
		caseNum = 4
	}

	xdiff := int(math.Abs(float64(coords[0] - coords[2])))
	ydiff := int(math.Abs(float64(coords[1] - coords[3])))
	for ii := 0; ii <= xdiff; ii++ {
		for jj := 0; jj <= ydiff; jj++ {
			if ii == jj {
				if caseNum == 1 {
					if temp[coords[0]+ii][coords[1]+jj] == "." {
						temp[coords[0]+ii][coords[1]+jj] = "1"
					} else {
						curr, _ := strconv.Atoi(temp[coords[0]+ii][coords[1]+jj])
						curr++
						new := strconv.Itoa(curr)
						temp[coords[0]+ii][coords[1]+jj] = new
					}
				} else if caseNum == 2 {
					if temp[coords[0]-ii][coords[1]+jj] == "." {
						temp[coords[0]-ii][coords[1]+jj] = "1"
					} else {
						curr, _ := strconv.Atoi(temp[coords[0]-ii][coords[1]+jj])
						curr++
						new := strconv.Itoa(curr)
						temp[coords[0]-ii][coords[1]+jj] = new
					}
				} else if caseNum == 3 {
					if temp[coords[0]+ii][coords[1]-jj] == "." {
						temp[coords[0]+ii][coords[1]-jj] = "1"
					} else {
						curr, _ := strconv.Atoi(temp[coords[0]+ii][coords[1]-jj])
						curr++
						new := strconv.Itoa(curr)
						temp[coords[0]+ii][coords[1]-jj] = new
					}
				} else if caseNum == 4 {
					if temp[coords[0]-ii][coords[1]-jj] == "." {
						temp[coords[0]-ii][coords[1]-jj] = "1"
					} else {
						curr, _ := strconv.Atoi(temp[coords[0]-ii][coords[1]-jj])
						curr++
						new := strconv.Itoa(curr)
						temp[coords[0]-ii][coords[1]-jj] = new
					}
				}
			}
		}
	}
	return temp
}

func isDiagonal(coords []int) bool {
	if math.Abs(float64(coords[0]-coords[2])) == math.Abs(float64(coords[1]-coords[3])) {
		return true
	}
	return false
}

func surveySeaFloor(queue [][]int, ventMap [][]string) int {
	total := 0
	tempMap := ventMap
	for ii := 0; ii < len(queue); ii++ {
		if queue[ii][0] == queue[ii][2] {
			tempMap = markYs(queue[ii], tempMap)
		} else if queue[ii][1] == queue[ii][3] {
			tempMap = markXs(queue[ii], tempMap)
		} else if isDiagonal(queue[ii]) {
			tempMap = markDiagonal(queue[ii], tempMap)
		}
	}

	for ii := 0; ii < 1000; ii++ {
		for jj := 0; jj < 1000; jj++ {
			if tempMap[ii][jj] == "." {
				continue
			}
			tempNum, _ := strconv.Atoi(tempMap[ii][jj])
			if tempNum >= 2 {
				total++
			}
		}
	}

	return total
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day5puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	ventMap := [][]string{}
	for ii := 0; ii < 1000; ii++ {
		ventRow := []string{}
		for jj := 0; jj < 1000; jj++ {
			ventRow = append(ventRow, ".")
		}
		ventMap = append(ventMap, ventRow)
	}

	formattedQueue := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		inParts := strings.Split(line, " ")
		pointA := strings.Split(inParts[0], ",")
		pointB := strings.Split(inParts[2], ",")
		addToQueue := []int{}
		x1, _ := strconv.Atoi(pointA[0])
		y1, _ := strconv.Atoi(pointA[1])
		x2, _ := strconv.Atoi(pointB[0])
		y2, _ := strconv.Atoi(pointB[1])
		addToQueue = append(addToQueue, x1)
		addToQueue = append(addToQueue, y1)
		addToQueue = append(addToQueue, x2)
		addToQueue = append(addToQueue, y2)
		formattedQueue = append(formattedQueue, addToQueue)
	}
	fmt.Print(formattedQueue)

	results := surveySeaFloor(formattedQueue, ventMap)
	fmt.Println(results)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
