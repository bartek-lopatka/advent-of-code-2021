package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	counter := 0
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

	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			counter++
		}
	}
	fmt.Println("it's a final counter to to ru to")
	fmt.Println(counter)
}
