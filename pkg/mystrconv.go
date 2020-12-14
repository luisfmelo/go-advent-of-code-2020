package pkg

import (
	"fmt"
	"strconv"
)

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func RuneToInt(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		panic(err)
	}
	return i
}

func ParseBinaryString(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func IntToBinaryString(i int, n int) string {
	bin := strconv.FormatInt(int64(i), 2)
	if n == -1 {
		return fmt.Sprintf("%v", bin)
	}
	return fmt.Sprintf("%0*v", n, bin)
}

func ReplaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}
