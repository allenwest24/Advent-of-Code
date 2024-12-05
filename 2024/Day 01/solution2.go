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

	total := 0

	// Count occurrences of item in list 1 in list 2
	for i := 0; i < len(numbers1); i++ {
		curr_count := 0
		for j := 0; j < len(numbers2); j++ {
			if numbers1[i] == numbers2[j] {
				curr_count++
			}
		}
		total += numbers1[i] * curr_count
	}

	// Print the total
	fmt.Println(total)
}
