package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"math"
	"os"
)

type CountSeatsRule func(rowIndex, colIndex int, matrix [][]rune) int

func OccupiedSeatsOnSight(rowIndex, colIndex int, matrix [][]rune) int {
	occupiedSeatsOnSight := 0

	// up
	for ri := rowIndex - 1; ri >= 0; ri-- {
		if matrix[ri][colIndex] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][colIndex] == 'L' {
			break
		}
	}

	// down
	for ri := rowIndex + 1; ri < len(matrix); ri++ {
		if matrix[ri][colIndex] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][colIndex] == 'L' {
			break
		}
	}

	// left
	for ci := colIndex - 1; ci >= 0; ci-- {
		if matrix[rowIndex][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[rowIndex][ci] == 'L' {
			break
		}
	}

	// right
	for ci := colIndex + 1; ci < len(matrix[rowIndex]); ci++ {
		if matrix[rowIndex][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[rowIndex][ci] == 'L' {
			break
		}
	}

	// diagonal up/right
	for ri, ci := rowIndex-1, colIndex+1; ri >= 0 && ci < len(matrix[rowIndex]); ri, ci = ri-1, ci+1 {
		if matrix[ri][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][ci] == 'L' {
			break
		}
	}

	// diagonal up/left
	for ri, ci := rowIndex-1, colIndex-1; ri >= 0 && ci >= 0; ri, ci = ri-1, ci-1 {
		if matrix[ri][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][ci] == 'L' {
			break
		}
	}

	// diagonal down/right
	for ri, ci := rowIndex+1, colIndex+1; ri < len(matrix) && ci < len(matrix[rowIndex]); ri, ci = ri+1, ci+1 {
		if matrix[ri][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][ci] == 'L' {
			break
		}
	}

	// diagonal down/left
	for ri, ci := rowIndex+1, colIndex-1; ri < len(matrix) && ci >= 0; ri, ci = ri+1, ci-1 {
		if matrix[ri][ci] == '#' {
			occupiedSeatsOnSight++
			break
		}
		if matrix[ri][ci] == 'L' {
			break
		}
	}

	return occupiedSeatsOnSight
}

func OccupiedAdjacentSeats(rowIndex, colIndex int, matrix [][]rune) int {
	occupiedAdjacentSeats := 0

	if rowIndex < len(matrix)-1 && matrix[rowIndex+1][colIndex] == '#' {
		occupiedAdjacentSeats++
	}
	if rowIndex >= 1 && matrix[rowIndex-1][colIndex] == '#' {
		occupiedAdjacentSeats++
	}
	if colIndex >= 1 && matrix[rowIndex][colIndex-1] == '#' {
		occupiedAdjacentSeats++
	}
	if colIndex < len(matrix[rowIndex])-1 && matrix[rowIndex][colIndex+1] == '#' {
		occupiedAdjacentSeats++
	}
	if rowIndex < len(matrix)-1 && colIndex < len(matrix[rowIndex])-1 && matrix[rowIndex+1][colIndex+1] == '#' {
		occupiedAdjacentSeats++
	}
	if rowIndex < len(matrix)-1 && colIndex >= 1 && matrix[rowIndex+1][colIndex-1] == '#' {
		occupiedAdjacentSeats++
	}
	if rowIndex >= 1 && colIndex < len(matrix[rowIndex])-1 && matrix[rowIndex-1][colIndex+1] == '#' {
		occupiedAdjacentSeats++
	}
	if rowIndex >= 1 && colIndex >= 1 && matrix[rowIndex-1][colIndex-1] == '#' {
		occupiedAdjacentSeats++
	}
	return occupiedAdjacentSeats
}

type NewSeatStateRule func(rowIndex, colIndex int, matrix [][]rune) rune

func GetNewSeatStatePart1(rowIndex, colIndex int, matrix [][]rune) rune {
	switch matrix[rowIndex][colIndex] {
	case 'L':
		if OccupiedAdjacentSeats(rowIndex, colIndex, matrix) == 0 {
			return '#'
		}

	case '#':
		if OccupiedAdjacentSeats(rowIndex, colIndex, matrix) >= 4 {
			return 'L'
		}

	}
	return matrix[rowIndex][colIndex]
}

func GetNewSeatStatePart2(rowIndex, colIndex int, matrix [][]rune) rune {
	switch matrix[rowIndex][colIndex] {
	case 'L':
		if OccupiedSeatsOnSight(rowIndex, colIndex, matrix) == 0 {
			return '#'
		}

	case '#':
		if OccupiedSeatsOnSight(rowIndex, colIndex, matrix) >= 5 {
			return 'L'
		}

	}
	return matrix[rowIndex][colIndex]
}

func CountOccupiedSeats(matrix [][]rune) int {
	occupiedSeats := 0
	for _, row := range matrix {
		for _, seat := range row {
			if seat == '#' {
				occupiedSeats++
			}
		}
	}
	return occupiedSeats
}

func Print(matrix [][]rune) {
	for _, row := range matrix {
		for _, seat := range row {
			fmt.Printf("%c", seat)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func NumberOfOccupiedSeatsAfterStabilization(matrix [][]rune, getNewSeatState NewSeatStateRule) int {
	iterations := 0
	seatsChanged := float64(len(matrix))
	var matrixSeatsCount int

	for seatsChanged > 0 {
		Print(matrix)
		iterations++

		// Create new matrix of changes
		newMatrix := make([][]rune, len(matrix))
		for rowIndex, row := range matrix {
			newMatrix[rowIndex] = make([]rune, len(matrix[rowIndex]))
			for colIndex, _ := range row {
				newMatrix[rowIndex][colIndex] = getNewSeatState(rowIndex, colIndex, matrix)
			}
		}

		// count changes
		newMatrixSeatsCount := CountOccupiedSeats(newMatrix)

		// update matrix
		seatsChanged = math.Abs(float64(newMatrixSeatsCount - matrixSeatsCount))
		matrixSeatsCount = newMatrixSeatsCount
		matrix = newMatrix
	}

	return CountOccupiedSeats(matrix)
}

func main() {
	file, err := os.Open("11/input.txt")
	if err != nil {
		panic(err)
	}

	matrix, err := pkg.ReadMatrix(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	// fmt.Println(NumberOfOccupiedSeatsAfterStabilization(matrix, GetNewSeatStatePart1))
	fmt.Println(NumberOfOccupiedSeatsAfterStabilization(matrix, GetNewSeatStatePart2))
}
