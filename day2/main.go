package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontal := 0
	depth := 0
	d, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic("input file not found.")
	}

	s := strings.Split(string(d), "\n")

	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "forward") {
			move := s[i]
			stringNumber := move[len(move)-1:]
			intNumber, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Panic("input is not an int!")
			}
			horizontal += intNumber

		}

		if strings.HasPrefix(s[i], "down") {
			move := s[i]
			stringNumber := move[len(move)-1:]
			intNumber, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Panic("input is not an int!")
			}
			depth += intNumber

		}

		if strings.HasPrefix(s[i], "up") {
			move := s[i]
			stringNumber := move[len(move)-1:]
			intNumber, err := strconv.Atoi(stringNumber)
			if err != nil {
				log.Panic("input is not an int!")
			}
			depth -= intNumber

		}
	}

	fmt.Println(horizontal * depth)
}
