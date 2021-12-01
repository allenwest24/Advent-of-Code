package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1puzzle2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	prevSum := 0
	a := 0
	b := 0
	c := 0
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		a = b
		b = c
		c = num
		currSum := a + b + c
		if currSum > prevSum && prevSum != 0 {
			counter++
		}
		if a != 0 && b != 0 && c != 0 {
			prevSum = currSum
		}
	}
	fmt.Print(counter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
