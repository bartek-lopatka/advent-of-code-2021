package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	gamma := ""
	epsilon := ""
	positions := make([]int, 12, 12)
	d, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic("input file not found.")
	}

	s := strings.Split(string(d), "\n")

	for i := 0; i < len(s); i++ {
		for j := 0; j < 12; j++ {
			switch {
			case s[i][j] == '0':
				positions[j] = positions[j] - 1
			case s[i][j] == '1':
				positions[j] = positions[j] + 1
			}
		}
	}

	for i := 0; i < len(positions); i++ {
		switch {
		case positions[i] > 0:
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		case positions[i] < 0:
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	gammaDecimal, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Panic("problem with the string.")
	}
	epsilonDecimal, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Panic("problem with the string.")
	}
	fmt.Printf("Solution of the part1: %d", epsilonDecimal*gammaDecimal)

}
