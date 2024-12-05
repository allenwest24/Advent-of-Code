package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Track two lists of numbers
	var numbers1 []int
	var numbers2 []int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := strings.Split(line, "   ")
		// Convert the strings to integers
		p1, _ := strconv.Atoi(lineNumbers[0])
		p2, _ := strconv.Atoi(lineNumbers[1])
		numbers1 = append(numbers1, p1)
		numbers2 = append(numbers2, p2)
	}

	// Sort list1
	for i := 0; i < len(numbers1); i++ {
		for j := i + 1; j < len(numbers1); j++ {
			if numbers1[i] > numbers1[j] {
				numbers1[i], numbers1[j] = numbers1[j], numbers1[i]
			}
		}
	}

	// Sort list2
	for i := 0; i < len(numbers2); i++ {
		for j := i + 1; j < len(numbers2); j++ {
			if numbers2[i] > numbers2[j] {
				numbers2[i], numbers2[j] = numbers2[j], numbers2[i]
			}
		}
	}
	fmt.Println(numbers2)
	// Go throught the two lists and get all of the differences
	total := 0
	for i := 0; i < len(numbers1); i++ {
		curr := numbers1[i] - numbers2[i]
		// get absolute value
		if curr < 0 {
			curr = -curr
		}
		total += curr
	}

	// Print the total
	fmt.Println(total)
}
