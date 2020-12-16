package pkg

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

// ReadLines reads \n separated strings from r. If there's an error, it
// returns the lines successfully read so far as well as the error value.
func ReadLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

// ReadLines reads \n separated strings from r. If there's an error, it
// returns the lines successfully read so far as well as the error value.
func ReadByDelimiter(r io.Reader, delimiter string) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitAt(delimiter))
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
func ReadIntsByDelimiter(r io.Reader, delimiter string) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitAt(delimiter))
	var result []int
	for scanner.Scan() {
		result = append(result, StrToInt(scanner.Text()))
	}
	return result, scanner.Err()
}

func SplitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(substring)
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

// ReadMatrix reads \n separated strings from r. And then each character.
// If there's an error, it returns the lines successfully read so far as well as the error value.
func ReadMatrix(r io.Reader) ([][]rune, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result [][]rune
	for scanner.Scan() {
		var row []rune
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		result = append(result, row)
	}
	return result, scanner.Err()
}
