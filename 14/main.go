package main

import (
	"advent-of.code.2020/pkg"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var inputLineRegex = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func DecoderVersion1(input []string) int {
	var forcedBits map[int]int
	memory := map[int]int{}
	for _, line := range input {
		// process bit mask
		if strings.HasPrefix(line, "mask") {
			forcedBits = map[int]int{}
			bitMask := strings.Replace(line, "mask = ", "", 1)
			for index, bit := range bitMask {
				if bit == 'X' {
					continue
				}
				forcedBits[len(bitMask)-1-index] = pkg.RuneToInt(bit)
			}
			continue
		}

		// process memory addresses changes
		gs := inputLineRegex.FindStringSubmatch(line)
		memoryAddressPosition := pkg.StrToInt(gs[1])
		memoryAddressValue := pkg.StrToInt(gs[2])

		// replace bits represented in mask
		for index, bit := range forcedBits {
			memoryAddressValue = pkg.ForceBit(memoryAddressValue, uint(index), bit)
		}

		memory[memoryAddressPosition] = memoryAddressValue
	}

	// sum memory addresses
	sum := 0
	for _, memoryAddress := range memory {
		sum += memoryAddress
	}

	return sum
}

func DecoderVersion2(input []string) int {
	var bitMask map[int]string
	memory := map[int]int{}
	for _, line := range input {
		// process bit mask
		if strings.HasPrefix(line, "mask") {
			bitMask = map[int]string{}
			bitMaskStr := strings.Replace(line, "mask = ", "", 1)
			for index, bit := range bitMaskStr {
				bitMask[index] = string(bit)
			}
			continue
		}

		// process memory addresses changes
		gs := inputLineRegex.FindStringSubmatch(line)
		memoryAddressPositionStr := gs[1]
		memoryAddressPositionBinaryStr := pkg.IntToBinaryString(pkg.StrToInt(memoryAddressPositionStr), 36)
		var floatingBitsIndexes []int
		for index, bit := range bitMask {
			if bit == "0" {
				continue
			} else if bit == "1" {
				memoryAddressPositionBinaryStr = pkg.ReplaceAtIndex(memoryAddressPositionBinaryStr, "1", index)
			} else if bit == "X" {
				memoryAddressPositionBinaryStr = pkg.ReplaceAtIndex(memoryAddressPositionBinaryStr, "X", index)
				floatingBitsIndexes = append(floatingBitsIndexes, index)
			}
		}

		// generate combinations for memory addresses
		nCombinations := len(floatingBitsIndexes)
		var memoryAddressesCombinations []string
		for i := float64(0); i < math.Pow(2, float64(nCombinations)); i++ {
			bin := pkg.IntToBinaryString(int(i), len(floatingBitsIndexes))
			memoryAddressesCombination := memoryAddressPositionBinaryStr
			for index, floatingBitIndex := range floatingBitsIndexes {
				memoryAddressesCombination = pkg.ReplaceAtIndex(memoryAddressesCombination, string(bin[index]), floatingBitIndex)
			}
			memoryAddressesCombinations = append(memoryAddressesCombinations, memoryAddressesCombination)
		}

		// update memory addresses
		memoryAddressValue := pkg.StrToInt(gs[2])
		for _, memoryAddressesCombination := range memoryAddressesCombinations {
			memory[pkg.ParseBinaryString(memoryAddressesCombination)] = memoryAddressValue
		}

	}

	// sum memory addresses
	sum := 0
	for _, memoryAddress := range memory {
		sum += memoryAddress
	}

	return sum
}

func main() {
	file, err := os.Open("14/input.txt")
	if err != nil {
		panic(err)
	}

	lines, err := pkg.ReadLines(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}

	// fmt.Println(DecoderVersion1(lines))
	fmt.Println(DecoderVersion2(lines))
}
