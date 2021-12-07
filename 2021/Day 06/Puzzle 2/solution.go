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

func simulateFishReproduction(fishTypeCounter []int, days int) int {
	for day := 0; day < days; day++ {
		newFish := fishTypeCounter[0]
		fishTypeCounter[0] = fishTypeCounter[1]
		fishTypeCounter[1] = fishTypeCounter[2]
		fishTypeCounter[2] = fishTypeCounter[3]
		fishTypeCounter[3] = fishTypeCounter[4]
		fishTypeCounter[4] = fishTypeCounter[5]
		fishTypeCounter[5] = fishTypeCounter[6]
		fishTypeCounter[6] = fishTypeCounter[7] + newFish
		fishTypeCounter[7] = fishTypeCounter[8]
		fishTypeCounter[8] = newFish
	}
	total := 0
	for ii := 0; ii < len(fishTypeCounter); ii++ {
		total += fishTypeCounter[ii]
	}
	return total
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

	fishTypeCounter := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for ii := 0; ii < len(lanternFish); ii++ {
		fishTypeCounter[lanternFish[ii]]++
	}

	fmt.Print(fishTypeCounter)
	total := simulateFishReproduction(fishTypeCounter, 256)
	fmt.Println(total)

	// Checking for any errors that occurred.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
