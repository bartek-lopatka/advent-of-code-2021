package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	counterPart1 := 0
	counterPart2 := 0
	//previous := 0
	var lines []int
	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, item)

	}

	// part one
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			counterPart1++
		}
	}

	//part two
	for i := 0; i < len(lines)-3; i++ {
		if lines[i]+lines[i+1]+lines[i+2] < lines[i+1]+lines[i+2]+lines[i+3] {
			counterPart2++
		}
	}
	fmt.Println("it's a final counter for part1 to to ru to")
	fmt.Println(counterPart1)
	fmt.Println("it's a final counter for part2 to to ru to")
	fmt.Println(counterPart2)
}
