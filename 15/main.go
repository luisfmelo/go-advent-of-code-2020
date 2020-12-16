package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func GetNumberInPlayingTurn(input []int, turn int) int {
	var numberSpoken, lastNumberSpoken int
	memory := map[int][]int{}

	for t, v := range input {
		if _, alreadySpoken := memory[v]; !alreadySpoken {
			memory[v] = []int{}
		}
		memory[v] = append(memory[v], t+1)
	}

	lastNumberSpoken = input[len(input)-1]
	for t := len(input)+1; t <= turn; t++ {
		switch len(memory[lastNumberSpoken]) {
		case 0:
			panic("this is not possible")
		case 1:
			// check if the last number was first time spoken: return 0
			numberSpoken = 0

		default:
			// check if the last number was spoken more time: return difference of the last 2 turns when the number was spoken
			nTurnsSpoken := len(memory[lastNumberSpoken])
			numberSpoken = memory[lastNumberSpoken][nTurnsSpoken-1] - memory[lastNumberSpoken][nTurnsSpoken-2]

		}
		if _, alreadySpoken := memory[numberSpoken]; !alreadySpoken {
			memory[numberSpoken] = []int{}
		}
		memory[numberSpoken] = append(memory[numberSpoken], t)

		lastNumberSpoken = numberSpoken
	}
	return numberSpoken
}

func main() {
	file, err := os.Open("15/input.txt")
	if err != nil {
		panic(err)
	}

	numbers, err := pkg.ReadIntsByDelimiter(bufio.NewReader(file), ",")
	if err != nil {
		panic(err)
	}

	start := time.Now()

	fmt.Println(GetNumberInPlayingTurn(numbers, 2020))
	log.Printf("First part took %s", time.Since(start))

	start = time.Now()
	fmt.Println(GetNumberInPlayingTurn(numbers, 30000000))
	log.Printf("Second part took %s", time.Since(start))
}
