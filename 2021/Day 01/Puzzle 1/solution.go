package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1puzzle1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	prev := -1
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		if prev == -1 {
			prev = num
			continue
		} else if num > prev {
			counter++
		}
		prev = num
	}
	fmt.Print(counter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
