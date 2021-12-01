package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	counter := -1
	previous := 0
	input, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		next, _ := strconv.Atoi(scanner.Text())
		if next > previous {
			counter++
		}
		fmt.Printf("old previous: %s\n", previous)
		previous = next
		fmt.Printf("new previous: %s\n", previous)
	}
	fmt.Println("it's a final counter to to ru to")
	fmt.Println(counter)
}
