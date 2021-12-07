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

func simulateFishReproductionCycles(lanternFish []int, totalDays int) int {
	tempFishArray := lanternFish
	for day := 0; day < totalDays; day++ {
		newFish := 0
		for fish := 0; fish < len(tempFishArray); fish++ {
			if tempFishArray[fish] == 0 {
				tempFishArray[fish] = 6
				newFish++
			} else {
				tempFishArray[fish]--
			}
		}
		for ii := 0; ii < newFish; ii++ {
			tempFishArray = append(tempFishArray, 8)
		}
	}
	return len(tempFishArray)
}

// Main method for this solution.
func main() {
	// Reading in the input file.
	f, err := os.Open("day6puzzles.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	lanternFish := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, ",")
		for ii := 0; ii < len(input); ii++ {
			fish, _ := strconv.Atoi(input[ii])
			lanternFish = append(lanternFish, fish)
		}
	}

	fmt.Print(lanternFish)

	totalFish := simulateFishReproductionCycles(lanternFish, 80)
	fmt.Println(totalFish)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
